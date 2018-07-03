package ini

type ASTKind int

const (
	ASTKindNone = ASTKind(iota)
	ASTKindStart
	ASTKindExpr
	ASTKindStatement
	ASTKindExprStatement
	ASTKindNestedTableStatement
	ASTKindCompletedNestedTableStatement
	ASTKindContainerExpr
	ASTKindCommentStatement
)

func (k ASTKind) String() string {
	switch k {
	case ASTKindNone:
		return "none"
	case ASTKindStart:
		return "start"
	case ASTKindExpr:
		return "expr"
	case ASTKindStatement:
		return "stmt"
	case ASTKindExprStatement:
		return "expr_stmt"
	case ASTKindNestedTableStatement:
		return "nested_table_stmt"
	default:
		return ""
	}
}

type AST interface {
	Kind() ASTKind
}

type Start struct{}

func (node Start) Kind() ASTKind {
	return ASTKindStart
}
