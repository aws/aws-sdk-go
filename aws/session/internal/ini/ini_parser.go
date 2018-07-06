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
// stmt -> nop | op id
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

const (
	InvalidState = iota
	StatementState
	StatementPrimeState
	ValueState
	OpenScopeState
	SectionState
	CloseScopeState
	SkipState
	SkipTokenState
	CommentState
	EpsilonState
	TerminalState
)

var parseTable = map[ASTKind]map[tokenType]int{
	ASTKindStart: map[tokenType]int{
		tokenLit:     StatementState, // stmt -> expr stmt'
		tokenSep:     OpenScopeState, // table -> [ table' | [ array_table
		tokenWS:      SkipTokenState, // skip token
		tokenNL:      SkipTokenState, // skip token
		tokenComment: CommentState,   // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    TerminalState,
	},
	ASTKindCommentStatement: map[tokenType]int{
		tokenLit:     StatementState, // stmt -> expr stmt'
		tokenSep:     OpenScopeState, // table -> [ table' | [ array_table
		tokenWS:      SkipTokenState, // skip token
		tokenNL:      SkipTokenState, // skip token
		tokenComment: CommentState,   // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    EpsilonState,
	},
	ASTKindExpr: map[tokenType]int{
		tokenOp:      StatementPrimeState, // stmt' -> nop | op stmt
		tokenLit:     ValueState,          // value -> number | string | boolean
		tokenSep:     OpenScopeState,
		tokenWS:      SkipTokenState, // skip token
		tokenNL:      SkipState,      // skip section
		tokenComment: CommentState,   // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    EpsilonState,
	},
	ASTKindStatement: map[tokenType]int{
		tokenLit:     SectionState,    // table -> [ table' | [ array_table
		tokenSep:     CloseScopeState, // array_close -> ] epsilon
		tokenWS:      SkipTokenState,
		tokenNL:      SkipTokenState,
		tokenComment: CommentState, // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    EpsilonState,
	},
	ASTKindExprStatement: map[tokenType]int{
		tokenLit:     ValueState, // stmt -> expr stmt'
		tokenSep:     OpenScopeState,
		tokenOp:      ValueState,
		tokenWS:      ValueState,
		tokenNL:      EpsilonState,
		tokenComment: CommentState, // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    TerminalState,
	},
	ASTKindCompletedSectionStatement: map[tokenType]int{
		tokenWS:      SkipTokenState,
		tokenNL:      SkipTokenState,
		tokenLit:     StatementState, // stmt -> expr stmt'
		tokenSep:     OpenScopeState, // table -> [ table' | [ array_table
		tokenComment: CommentState,   // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    EpsilonState,
	},
	ASTKindSkipStatement: map[tokenType]int{
		tokenLit:     StatementState, // stmt -> expr stmt'
		tokenSep:     OpenScopeState, // table -> [ table' | [ array_table
		tokenWS:      SkipTokenState, // skip token
		tokenNL:      SkipTokenState, // skip token
		tokenComment: CommentState,   // comment -> #comment' | ;comment' | /comment_slash
		tokenNone:    TerminalState,
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
			step = SkipTokenState
		}

		switch step {
		case TerminalState:
			if k.Kind() != ASTKindStart {
				stack.Epsilon(k)
			}
			break loop
		case SkipTokenState:
			stack.Push(k)
		case StatementState:
			if k.Kind() != ASTKindStart {
				stack.Epsilon(k)
			}
			expr := newExpression(tok)
			stack.Push(expr)
		case StatementPrimeState:
			if tok.Type() != tokenOp {
				stack.Epsilon(k)
				continue
			}

			expr := newEqualExpr(k, tok)
			stack.Push(expr)
		case ValueState:
			switch v := k.(type) {
			case EqualExpr:
				v.Right = newExpression(tok)
				stack.Push(newExprStatement(v))
			case ExprStatement:
				expr, ok := v.V.(EqualExpr)
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
					v.V = expr
					stack.Push(v)
				} else {
					stack.Push(k)
				}
			default:
				return stack.list, awserr.New(ErrCodeParseError, "invalid expression", nil)
			}
		case OpenScopeState:
			if tok.Raw() != "[" {
				return stack.list, awserr.New(ErrCodeParseError, "expected '['", nil)
			}

			stmt := newStatement()
			stack.Push(stmt)
		case CloseScopeState:
			if tok.Raw() == "]" {
				stack.Push(newCompletedSectionStatement(k))
			} else {
				return stack.list, awserr.New(ErrCodeParseError, "expected ']'", nil)
			}
		case SectionState:
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
		case EpsilonState:
			if k.Kind() != ASTKindStart {
				stack.Epsilon(k)
			}

			if stack.Len() == 0 {
				stack.Push(start)
			}
		case SkipState:
			stack.Push(newSkipStatement(k))
			s.Skip()
		case CommentState:
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
