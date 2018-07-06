package ini

type ASTKind int

// ASTKind* is used in the parse table to transition between
// the different states
const (
	ASTKindNone = ASTKind(iota)
	ASTKindStart
	ASTKindExpr
	ASTKindStatement
	ASTKindSkipStatement
	ASTKindExprStatement
	ASTKindNestedSectionStatement
	ASTKindCompletedNestedSectionStatement
	ASTKindCommentStatement
	ASTKindCompletedSectionStatement
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
	case ASTKindCommentStatement:
		return "comment"
	case ASTKindNestedSectionStatement:
		return "nested_section_stmt"
	case ASTKindCompletedSectionStatement:
		return "completed_stmt"
	case ASTKindSkipStatement:
		return "skip"
	default:
		return ""
	}
}

// AST interface allows us to determine what kind of node we
// are on and casting may not need to be necessary.
type AST interface {
	Kind() ASTKind
}

// Start represents the starting parsing state.
type Start struct{}

func (node Start) Kind() ASTKind {
	return ASTKindStart
}
