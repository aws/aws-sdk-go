package ini

import (
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

const (
	ErrCodeParseError = "ParseError"
)

// id -> value stmt
// stmt -> expr stmt'
// stmt' -> nop | op stmt
// value -> number | string | boolean
//
// table -> [ table' | [ array_table
// table' -> label array_close
// array_close -> ] epsilon
//
// array_table -> [ table_nested
// table_nested -> label nested_array_close
// nested_array_close -> ] array_close
//
// comment -> #comment' | ;comment' | /comment_slash
// comment_slash -> / comment'
// comment' -> string | epsilon
//
// epsilon -> nop
var parseTable = map[ASTKind]map[tokenType]int{
	ASTKindStart: map[tokenType]int{
		tokenLit:     1,  // stmt -> expr stmt'
		tokenSep:     4,  // table -> [ table' | [ array_table
		tokenWS:      -1, // skip token
		tokenNL:      -1, // skip token
		tokenComment: 11, // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    -2,
	},
	ASTKindCommentStatement: map[tokenType]int{
		tokenLit:     1,  // stmt -> expr stmt'
		tokenSep:     4,  // table -> [ table' | [ array_table
		tokenWS:      -1, // skip token
		tokenNL:      -1, // skip token
		tokenComment: 11, // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    13,
	},
	ASTKindExpr: map[tokenType]int{
		tokenOp:      2, // stmt' -> nop | op stmt
		tokenLit:     3, // value -> number | string | boolean
		tokenSep:     2,
		tokenWS:      -1, // skip token
		tokenNL:      10, // skip section
		tokenComment: 11, // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    13,
	},
	ASTKindStatement: map[tokenType]int{
		// TODO: fix 5 and 6 state transitions. Should have TableStatement return
		// ASTKindTableStatement instead of ASTKindStatement.
		tokenLit:     5, // table -> [ table' | [ array_table
		tokenSep:     6, // array_close -> ] epsilon
		tokenWS:      -1,
		tokenNL:      -1,
		tokenComment: 11, // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    13,
	},
	ASTKindExprStatement: map[tokenType]int{
		tokenLit:     12, // stmt -> expr stmt'
		tokenSep:     2,
		tokenOp:      12,
		tokenWS:      12,
		tokenNL:      13,
		tokenComment: 11, // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    -2,
	},
	ASTKindNestedSectionStatement: map[tokenType]int{
		tokenLit:     7, // table_nested -> label nested_array_close
		tokenSep:     8, // nested_array_close -> ] array_close
		tokenWS:      -1,
		tokenNL:      -1,
		tokenComment: 11, // comment -> #comment' | ;comment' | /comment_slash
	},
	ASTKindCompletedNestedSectionStatement: map[tokenType]int{
		tokenSep:     9, // nested_array_close -> ] epsilon
		tokenWS:      -1,
		tokenNL:      -1,
		tokenComment: 11, // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    13,
	},
	ASTKindCompletedSectionStatement: map[tokenType]int{
		tokenWS:      -1,
		tokenNL:      -1,
		tokenLit:     1,  // stmt -> expr stmt'
		tokenSep:     4,  // table -> [ table' | [ array_table
		tokenComment: 11, // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    13,
	},
	ASTKindSkipStatement: map[tokenType]int{
		tokenLit:     1,  // stmt -> expr stmt'
		tokenSep:     4,  // table -> [ table' | [ array_table
		tokenWS:      -1, // skip token
		tokenNL:      -1, // skip token
		tokenComment: 11, // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    -2,
	},
}

// skipper is used to skip certain blocks of an ini file.
// Currently skipper is used to skip nested blocks of ini
// files. See example below
//
//	[ foo ]
//	nested = // this section will be skipped
//		a=b
//		c=d
//	bar=baz // this will be included
type skipper struct {
	shouldSkip bool
	prevTok    iniToken
}

func (s *skipper) ShouldSkip(tok iniToken) bool {
	if s.shouldSkip && s.prevTok != nil && s.prevTok.Type() == tokenNL && tok.Type() != tokenWS {
		s.Continue()
		return false
	}
	s.prevTok = tok

	return s.shouldSkip
}

func (s *skipper) Skip() {
	s.shouldSkip = true
	s.prevTok = nil
}

func (s *skipper) Continue() {
	s.shouldSkip = false
	s.prevTok = nil
}

// ParseAST will parse input from an io.Reader using
// an LL(1) parser.
func ParseAST(r io.Reader) ([]AST, error) {
	lexer := iniLexer{}
	tokens, err := lexer.Tokenize(r)
	if err != nil {
		return []AST{}, err
	}

	start := Start{}
	stack := ParseStack{}
	stack.Push(start)
	s := skipper{}

loop:
	for stack.Len() > 0 {
		k := stack.Pop()

		var tok iniToken
		if len(tokens) == 0 {
			tok = emptyToken{}
		} else {
			tok = tokens[0]
		}

		step := parseTable[k.Kind()][tok.Type()]
		if s.ShouldSkip(tok) {
			step = -1
		}

		switch step {
		case -2:
			if k.Kind() != ASTKindStart {
				stack.Epsilon(k)
			}
			break loop
		case -1:
			stack.Push(k)
		case 1:
			if k.Kind() != ASTKindStart {
				stack.Epsilon(k)
			}
			expr := newExpression(tok)
			stack.Push(expr)
		case 2:
			if tok.Type() != tokenOp {
				stack.Epsilon(k)
				continue
			}

			expr := newEqualExpr(k, tok)
			stack.Push(expr)
		case 3:
			v, ok := k.(EqualExpr)
			if !ok {
				return stack.list, awserr.New(ErrCodeParseError, "invalid expression", nil)
			}

			v.Right = newExpression(tok)
			stack.Push(newExprStatement(v))
		case 4:
			if tok.Raw() != "[" {
				return stack.list, awserr.New(ErrCodeParseError, "expected '['", nil)
			}

			stmt := newStatement()
			stack.Push(stmt)
		case 5:
			if k.Kind() != ASTKindStatement {
				return stack.list, awserr.New(ErrCodeParseError, "invalid statement", nil)
			}

			var stmt AST

			if t, ok := k.(SectionStatement); ok {
				t.Name = strings.Join([]string{t.Name, tok.Raw()}, " ")
				stmt = t
			} else {
				stmt = newSectionStatement(tok)
			}
			stack.Push(stmt)
		case 12:
			stmt, ok := k.(ExprStatement)
			if !ok {
				return stack.list, awserr.New(ErrCodeParseError, "invalid expression", nil)
			}

			expr, ok := stmt.V.(EqualExpr)
			if !ok {
				return stack.list, awserr.New(ErrCodeParseError, "invalid expression", nil)
			}

			rhs, ok := expr.Right.(Expr)
			if !ok {
				return stack.list, awserr.New(ErrCodeParseError, "invalid expression", nil)
			}

			if rhs.Root.Type() != tokenLit {
				return stack.list, awserr.New(ErrCodeParseError, "invalid expression", nil)
			}

			t := rhs.Root.(literalToken)
			if t.Value.Type != QuotedStringType {
				t.Value.Append(tok)

				rhs.Root = t
				expr.Right = rhs
				stmt.V = expr
				stack.Push(stmt)
			} else {
				stack.Push(k)
			}
		case 13:
			if k.Kind() != ASTKindStart {
				stack.Epsilon(k)
			}

			if stack.Len() == 0 {
				stack.Push(start)
			}
		case 6:
			if tok.Raw() == "]" {
				stack.Push(newCompletedSectionStatement(k))
			} else if tok.Raw() == "[" {
				stmt := newNestedSectionStatement()
				stack.Push(stmt)
			} else {
				return stack.list, awserr.New(ErrCodeParseError, "expected ']'", nil)
			}
		case 7:
			switch tok.Type() {
			case tokenLit:
				stmt, ok := k.(NestedSectionStatement)
				if !ok {
					return stack.list, awserr.New(ErrCodeParseError, "expected nested section statement", nil)
				}

				stmt.Labels = append(stmt.Labels, tok.Raw())
				stack.Push(stmt)
			default:
				return stack.list, awserr.New(ErrCodeParseError, "expected literal", nil)
			}
		case 8:
			if tok.Raw() != "]" {
				return stack.list, awserr.New(ErrCodeParseError, "expected closing bracket", nil)
			}

			stmt := newCompletedNestedSectionStatement(k)
			stack.Push(stmt)
		case 9:
			if tok.Raw() != "]" {
				return stack.list, awserr.New(ErrCodeParseError, "expected closing bracket", nil)
			}

			stack.Epsilon(k)
		case 10:
			stack.Push(newSkipStatement(k))
			s.Skip()
		case 11:
			if _, ok := k.(Start); ok {
				stack.Push(k)
			} else {
				stack.Epsilon(k)
			}

			stmt := newCommentStatement(tok)
			stack.Push(stmt)
		default:
			return stack.list, awserr.New(ErrCodeParseError, "parse error: invalid state", nil)
		}

		if len(tokens) > 0 {
			tokens = tokens[1:]
		}
	}

	// this occurs when a statement has not been completed
	if len(stack.container) > 1 {
		return stack.list, awserr.New(ErrCodeParseError, "parse error", nil)
	}

	// returns a sublist which exludes the start symbol
	return stack.list, nil
}
