package ini

import (
	"fmt"
)

// Expr represents an expression
//
//	grammar:
//	expr -> string | number
type Expr struct {
	Root Token
}

// newExpression will return an expression AST.
func newExpression(token Token) Expr {
	return Expr{
		Root: token,
	}
}

// Kind will return the type of AST
func (expr Expr) Kind() ASTKind {
	return ASTKindExpr
}

func (expr Expr) String() string {
	return string(expr.Root.Raw())
}

// EqualExpr AST
type EqualExpr struct {
	Root  Token
	Left  AST
	Right AST
}

func newEqualExpr(left AST, token Token) EqualExpr {
	expr := EqualExpr{
		Left: left,
	}
	expr.Root = token
	return expr
}

// Kind will return the type of AST
func (expr EqualExpr) Kind() ASTKind {
	return ASTKindExpr
}

func (expr EqualExpr) String() string {
	return fmt.Sprintf("{%s %s %s}", expr.Left, expr.Root, expr.Right)
}

// Key will return a LHS value in the equal expr
func (expr EqualExpr) Key() string {
	return string(expr.Left.(Expr).Root.Raw())
}
