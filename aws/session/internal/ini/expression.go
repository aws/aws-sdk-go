package ini

import (
	"fmt"
)

type Expr struct {
	Root iniToken
}

// newExpression will return an expression AST.
func newExpression(token iniToken) Expr {
	return Expr{
		Root: token,
	}
}

func (expr Expr) Kind() ASTKind {
	return ASTKindExpr
}

func (expr Expr) String() string {
	return fmt.Sprintf("%s", expr.Root.Raw())
}

type EqualExpr struct {
	Root  iniToken
	Left  AST
	Right AST
}

func newEqualExpr(left AST, token iniToken) EqualExpr {
	expr := EqualExpr{
		Left: left,
	}
	expr.Root = token
	return expr
}

func (node EqualExpr) Kind() ASTKind {
	return ASTKindExpr
}

func (expr EqualExpr) String() string {
	return fmt.Sprintf("{%s %s %s}", expr.Left, expr.Root, expr.Right)
}

func (expr EqualExpr) Key() string {
	return expr.Left.(Expr).Root.Raw()
}
