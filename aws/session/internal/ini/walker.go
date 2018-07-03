package ini

func Walk(tree []AST, v Visitor) error {
	for _, node := range tree {
		switch node.Kind() {
		case ASTKindExpr,
			ASTKindExprStatement:

			if err := v.VisitExpr(node); err != nil {
				return err
			}
		case ASTKindStatement,
			ASTKindNestedTableStatement,
			ASTKindCompletedNestedTableStatement:

			if err := v.VisitStatement(node); err != nil {
				return err
			}
		}
	}

	return nil
}
