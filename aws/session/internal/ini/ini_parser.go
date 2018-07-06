package ini

import (
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

const (
	// ErrCodeParseError is returned when a parsing error
	// has occurred.
	ErrCodeParseError = "ParseError"
)

// State enums for the parse table
const (
	InvalidState = iota
	// stmt -> value stmt'
	StatementState
	// stmt' -> epsilon | op stmt
	StatementPrimeState
	// value -> number | string | boolean | quoted_string
	ValueState
	// section -> [ section'
	OpenScopeState
	// section' -> value section_close
	SectionState
	// section_close -> ]
	CloseScopeState
	// SkipState will skip (NL WS)+
	SkipState
	// SkipTokenState will skip any token and push the previous
	// state onto the stack.
	SkipTokenState
	// comment -> # comment' | ; comment' | / comment_slash
	// comment_slash -> / comment'
	// comment' -> epsilon | value
	CommentState
	// Epsilon state will complete statements and move that
	// to the completed AST list
	EpsilonState
	// TerminalState signifies that the tokens have been fully parsed
	TerminalState
)

// parseTable is a state machine to dictate the grammar above.
var parseTable = map[ASTKind]map[TokenType]int{
	ASTKindStart: map[TokenType]int{
		TokenLit:     StatementState,
		TokenSep:     OpenScopeState,
		TokenWS:      SkipTokenState,
		TokenNL:      SkipTokenState,
		TokenComment: CommentState,
		TokenNone:    TerminalState,
	},
	ASTKindCommentStatement: map[TokenType]int{
		TokenLit:     StatementState,
		TokenSep:     OpenScopeState,
		TokenWS:      SkipTokenState,
		TokenNL:      SkipTokenState,
		TokenComment: CommentState,
		TokenNone:    EpsilonState,
	},
	ASTKindExpr: map[TokenType]int{
		TokenOp:      StatementPrimeState,
		TokenLit:     ValueState,
		TokenSep:     OpenScopeState,
		TokenWS:      SkipTokenState,
		TokenNL:      SkipState,
		TokenComment: CommentState,
		TokenNone:    EpsilonState,
	},
	ASTKindStatement: map[TokenType]int{
		TokenLit:     SectionState,
		TokenSep:     CloseScopeState,
		TokenWS:      SkipTokenState,
		TokenNL:      SkipTokenState,
		TokenComment: CommentState,
		TokenNone:    EpsilonState,
	},
	ASTKindExprStatement: map[TokenType]int{
		TokenLit:     ValueState,
		TokenSep:     OpenScopeState,
		TokenOp:      ValueState,
		TokenWS:      ValueState,
		TokenNL:      EpsilonState,
		TokenComment: CommentState,
		TokenNone:    TerminalState,
		TokenComma:   SkipState,
	},
	ASTKindCompletedSectionStatement: map[TokenType]int{
		TokenWS:      SkipTokenState,
		TokenNL:      SkipTokenState,
		TokenLit:     StatementState,
		TokenSep:     OpenScopeState,
		TokenComment: CommentState,
		TokenNone:    EpsilonState,
	},
	ASTKindSkipStatement: map[TokenType]int{
		TokenLit:     StatementState,
		TokenSep:     OpenScopeState,
		TokenWS:      SkipTokenState,
		TokenNL:      SkipTokenState,
		TokenComment: CommentState,
		TokenNone:    TerminalState,
	},
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

		var tok Token
		if len(tokens) == 0 {
			// this occurs when all the tokens have been processed
			// but reduction of what's left on the stack needs to
			// occur.
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
			// Finished parsing. Push what should be the last
			// statement to the stack. If there is anything left
			// on the stack, an error in parsing has occurred.
			if k.Kind() != ASTKindStart {
				stack.Epsilon(k)
			}
			break loop
		case SkipTokenState:
			// When skipping a token, the previous state was popped off the stack.
			// To maintain the correct state, the previous state will be pushed
			// onto the stack.
			stack.Push(k)
		case StatementState:
			if k.Kind() != ASTKindStart {
				stack.Epsilon(k)
			}
			expr := newExpression(tok)
			stack.Push(expr)
		case StatementPrimeState:
			if tok.Type() != TokenOp {
				stack.Epsilon(k)
				continue
			}

			expr := newEqualExpr(k, tok)
			stack.Push(expr)
		case ValueState:
			// ValueState requires the previous state to either be an equal expression
			// or an expression statement.
			//
			// This grammar occurs when the RHS is a number, word, or quoted string.
			// equal_expr -> lit op equal_expr'
			// equal_expr' -> number | string | quoted_string
			// quoted_string -> " quoted_string'
			// quoted_string' -> string quoted_string_end
			// quoted_string_end -> "
			//
			// otherwise
			// expr_stmt -> equal_expr (expr_stmt')*
			// expr_stmt' -> ws S | op S | epsilon
			// S -> equal_expr' expr_stmt'
			switch v := k.(type) {
			case EqualExpr:
				v.Right = newExpression(tok)
				stack.Push(newExprStatement(v))
			case ExprStatement:
				expr, ok := v.V.(EqualExpr)
				if !ok {
					return stack.list, awserr.New(
						ErrCodeParseError,
						fmt.Sprintf("invalid expression: expected equal expression, but found %T",
							v.V,
						),
						nil,
					)
				}

				rhs, ok := expr.Right.(Expr)
				if !ok {
					return stack.list, awserr.New(
						ErrCodeParseError,
						fmt.Sprintf("invalid expression: RHS is not an expression:  %T",
							expr.Right,
						),
						nil,
					)
				}

				if rhs.Root.Type() != TokenLit {
					return stack.list, awserr.New(
						ErrCodeParseError,
						fmt.Sprintf("invalid expression: RHS is not a literal:  %v",
							rhs.Root,
						),
						nil,
					)
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
				return stack.list, awserr.New(
					ErrCodeParseError,
					fmt.Sprintf("invalid statement: expected statement: %T", k.Kind()),
					nil,
				)
			}

			var stmt AST

			if t, ok := k.(SectionStatement); ok {
				// If there are multiple literals inside of a scope declaration,
				// then the current token's raw value will be appended to the Name.
				//
				// This handles cases like [ profile default ]
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
