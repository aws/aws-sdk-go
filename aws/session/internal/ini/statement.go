package ini

import "fmt"

type Statement struct{}

func newStatement() Statement {
	return Statement{}
}

func (stmt Statement) Kind() ASTKind {
	return ASTKindStatement
}

type TableStatement struct {
	Name string
}

func newTableStatement(tok iniToken) TableStatement {
	return TableStatement{
		Name: tok.Raw(),
	}
}

func (stmt TableStatement) Kind() ASTKind {
	return ASTKindStatement
}

func (stmt TableStatement) String() string {
	return fmt.Sprintf("{%s}", stmt.Name)
}

type ExprStatement struct {
	V AST
}

func newExprStatement(v AST) ExprStatement {
	return ExprStatement{
		V: v,
	}
}

func (stmt ExprStatement) Kind() ASTKind {
	return ASTKindExprStatement
}

func (stmt ExprStatement) String() string {
	return fmt.Sprintf("{%v}", stmt.V)
}

type NestedTableStatement struct {
	Labels []string
	kind   ASTKind
}

func newNestedTableStatement() NestedTableStatement {
	return NestedTableStatement{
		kind: ASTKindNestedTableStatement,
	}
}

func (stmt NestedTableStatement) Kind() ASTKind {
	return ASTKindNestedTableStatement
}

func (stmt NestedTableStatement) String() string {
	return fmt.Sprintf("{[[ %v ]]}", stmt.Labels)
}

type CompletedNestedTableStatement struct {
	Root AST
}

func newCompletedNestedTableStatement(k AST) CompletedNestedTableStatement {
	return CompletedNestedTableStatement{
		Root: k,
	}
}

func (stmt CompletedNestedTableStatement) Kind() ASTKind {
	return ASTKindCompletedNestedTableStatement
}

func (stmt CompletedNestedTableStatement) String() string {
	return fmt.Sprintf("{[[ %v ]]}", stmt.Root)
}

// CommentStatement represents a comment in the ini defintion.
//
//	grammar:
//	comment -> #comment' | ;comment' | /comment_slash
//	comment_slash -> /comment'
//	comment' -> value
type CommentStatement struct {
	Comment iniToken
}

func newCommentStatement(tok iniToken) CommentStatement {
	return CommentStatement{
		Comment: tok,
	}
}

func (stmt CommentStatement) Kind() ASTKind {
	return ASTKindCommentStatement
}

func (stmt CommentStatement) String() string {
	return stmt.Comment.Raw()
}
