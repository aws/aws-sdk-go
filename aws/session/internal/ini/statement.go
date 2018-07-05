package ini

import "fmt"

type Statement struct{}

func newStatement() Statement {
	return Statement{}
}

func (stmt Statement) Kind() ASTKind {
	return ASTKindStatement
}

type SectionStatement struct {
	Name string
}

func newSectionStatement(tok iniToken) SectionStatement {
	return SectionStatement{
		Name: tok.Raw(),
	}
}

func (stmt SectionStatement) Kind() ASTKind {
	return ASTKindStatement
}

func (stmt SectionStatement) String() string {
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

type NestedSectionStatement struct {
	Labels []string
	kind   ASTKind
}

func newNestedSectionStatement() NestedSectionStatement {
	return NestedSectionStatement{
		kind: ASTKindNestedSectionStatement,
	}
}

func (stmt NestedSectionStatement) Kind() ASTKind {
	return ASTKindNestedSectionStatement
}

func (stmt NestedSectionStatement) String() string {
	return fmt.Sprintf("{[[ %v ]]}", stmt.Labels)
}

type CompletedNestedSectionStatement struct {
	Root AST
}

func newCompletedNestedSectionStatement(k AST) CompletedNestedSectionStatement {
	return CompletedNestedSectionStatement{
		Root: k,
	}
}

func (stmt CompletedNestedSectionStatement) Kind() ASTKind {
	return ASTKindCompletedNestedSectionStatement
}

func (stmt CompletedNestedSectionStatement) String() string {
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
