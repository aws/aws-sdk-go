package ini

type ContainerExpr struct {
}

func newContainerExpr() ContainerExpr {
	return ContainerExpr{}
}

func (expr ContainerExpr) Kind() ASTKind {
	return ASTKindContainerExpr
}
