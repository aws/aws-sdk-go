package ini

import "fmt"

// Statement is an empty AST mostly used
// for transitioning states.
type Statement struct{}

func newStatement() Statement {
	return Statement{}
}

// Kind returns the AST kind
func (stmt Statement) Kind() ASTKind {
	return ASTKindStatement
}

// SectionStatement represents a section AST
type SectionStatement struct {
	Name string
}

func newSectionStatement(tok Token) SectionStatement {
	return SectionStatement{
		Name: string(tok.Raw()),
	}
}

// Kind returns the AST kind
func (stmt SectionStatement) Kind() ASTKind {
	return ASTKindStatement
}

func (stmt SectionStatement) String() string {
	return fmt.Sprintf("{%s}", stmt.Name)
}

// ExprStatement represents a completed expression AST
type ExprStatement struct {
	V AST
}

func newExprStatement(v AST) ExprStatement {
	return ExprStatement{
		V: v,
	}
}

// Kind returns the AST kind
func (stmt ExprStatement) Kind() ASTKind {
	return ASTKindExprStatement
}

func (stmt ExprStatement) String() string {
	return fmt.Sprintf("{%v}", stmt.V)
}

// CommentStatement represents a comment in the ini defintion.
//
//	grammar:
//	comment -> #comment' | ;comment' | /comment_slash
//	comment_slash -> /comment'
//	comment' -> value
type CommentStatement struct {
	Comment Token
}

func newCommentStatement(tok Token) CommentStatement {
	return CommentStatement{
		Comment: tok,
	}
}

// Kind returns the AST kind
func (stmt CommentStatement) Kind() ASTKind {
	return ASTKindCommentStatement
}

func (stmt CommentStatement) String() string {
	return string(stmt.Comment.Raw())
}

// CompletedSectionStatement represents a completed section
type CompletedSectionStatement struct {
	V AST
}

func newCompletedSectionStatement(ast AST) CompletedSectionStatement {
	return CompletedSectionStatement{
		V: ast,
	}
}

// Kind returns the AST kind
func (stmt CompletedSectionStatement) Kind() ASTKind {
	return ASTKindCompletedSectionStatement
}

// SkipStatement is used to skip whole statements
type SkipStatement struct {
	V AST
}

func newSkipStatement(ast AST) SkipStatement {
	return SkipStatement{
		V: ast,
	}
}

// Kind returns the AST kind
func (stmt SkipStatement) Kind() ASTKind {
	return ASTKindSkipStatement
}
