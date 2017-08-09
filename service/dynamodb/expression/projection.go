package expression

import "strings"

type ProjectionBuilder struct {
	paths []PathBuilder
}

func Projection(p PathBuilder, pl ...PathBuilder) ProjectionBuilder {
	pl = append([]PathBuilder{p}, pl...)
	return ProjectionBuilder{
		paths: pl,
	}
}

func (p PathBuilder) Projection(pl ...PathBuilder) ProjectionBuilder {
	return Projection(p, pl...)
}

func AddPaths(proj ProjectionBuilder, pl ...PathBuilder) ProjectionBuilder {
	proj.paths = append(proj.paths, pl...)
	return proj
}

func (proj ProjectionBuilder) AddPaths(pl ...PathBuilder) ProjectionBuilder {
	return AddPaths(proj, pl...)
}

func (proj ProjectionBuilder) BuildExpression() (Expression, error) {
	en, err := proj.buildProjection()
	if err != nil {
		return Expression{}, err
	}

	expr, err := en.buildExprNodes(&aliasList{})
	if err != nil {
		return Expression{}, err
	}

	return expr, nil
}

func (proj ProjectionBuilder) buildProjection() (ExprNode, error) {
	childNodes, err := proj.buildChildNodes()
	if err != nil {
		return ExprNode{}, err
	}
	ret := ExprNode{
		children: childNodes,
	}

	ret.fmtExpr = "$c" + strings.Repeat(", $c", len(proj.paths)-1)

	return ret, nil
}

func (proj ProjectionBuilder) buildChildNodes() ([]ExprNode, error) {
	childNodes := make([]ExprNode, 0, len(proj.paths))
	for _, path := range proj.paths {
		en, err := path.BuildOperand()
		if err != nil {
			return []ExprNode{}, err
		}
		childNodes = append(childNodes, en)
	}

	return childNodes, nil
}
