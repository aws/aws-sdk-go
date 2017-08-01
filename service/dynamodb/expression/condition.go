package expression

import "fmt"

// ConditionMode will specify the types of the struct ConditionBuilder
type ConditionMode int

const (
	// UnsetCond will catch errors if users make an empty ConditionBuilder
	UnsetCond ConditionMode = iota
	// EqualCond will represent the Equal Clause ConditionBuilder
	EqualCond
	// AndCond will represent the And Clause ConditionBuilder
	AndCond
)

// ConditionBuilder will represent the ConditionExpressions
type ConditionBuilder struct {
	operandList   []OperandBuilder
	conditionList []ConditionBuilder
	Mode          ConditionMode
}

// Equal

// Equal will create a ConditionBuilder. This will be the function call
func Equal(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		Mode:        EqualCond,
	}
}

// Equal will create a ConditionBuilder. This will be the method for PathBuilder
func (p PathBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(p, right)
}

// Equal will create a ConditionBuilder. This will be the method for
// ValueBuilder
func (v ValueBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(v, right)
}

// Equal will create a ConditionBuilder. This will be the method for SizeBuilder
func (s SizeBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(s, right)
}

// And will create a ConditionBuilder. This will be the function call
func And(cond ...ConditionBuilder) ConditionBuilder {
	return ConditionBuilder{
		conditionList: cond,
		Mode:          AndCond,
	}
}

// And will create a ConditionBuilder. This will be the method signature
func (cond ConditionBuilder) And(right ...ConditionBuilder) ConditionBuilder {
	right = append(right, cond)
	return And(right...)
}

// BuildExpression will take an ConditionBuilder as input and output an
// Expression
func (cond ConditionBuilder) BuildExpression() (Expression, error) {
	en, err := cond.buildCondition()
	if err != nil {
		return Expression{}, err
	}

	expr, err := en.buildExprNodes(&aliasList{})
	if err != nil {
		return Expression{}, err
	}

	return expr, nil
}

// buildCondition will iterate over the tree of ConditionBuilders and
// OperandBuilders and build a tree of ExprNodes
func (cond ConditionBuilder) buildCondition() (ExprNode, error) {
	switch cond.Mode {
	case EqualCond:
		return compareBuildCondition(cond)
	case AndCond:
		return boolBuildCondition(cond)
	}
	return ExprNode{}, fmt.Errorf("No matching Mode to %v", cond.Mode)
}

// compareBuildCondition is the function to make ExprNodes from Compare
// ConditionBuilders
func compareBuildCondition(c ConditionBuilder) (ExprNode, error) {
	if len(c.conditionList) != 0 {
		return ExprNode{}, fmt.Errorf("Invalid ConditionBuilder. Expected 0 ConditionBuilders")
	}

	if len(c.operandList) != 2 {
		return ExprNode{}, fmt.Errorf("Invalid ConditionBuilder. Expected 2 Operands")
	}

	operandExprNodes := make([]ExprNode, 0, len(c.operandList))
	for _, ope := range c.operandList {
		exprNodes, err := ope.BuildOperand()
		if err != nil {
			return ExprNode{}, err
		}
		operandExprNodes = append(operandExprNodes, exprNodes)
	}

	ret := ExprNode{
		children: operandExprNodes,
	}

	// Create a string with special characters that can be substituted later: $c
	switch c.Mode {
	case EqualCond:
		ret.fmtExpr = "$c = $c"
	}

	return ret, nil
}

// boolBuildCondition is the function to make ExprNodes from And/Or
// ConditionBuilders
func boolBuildCondition(c ConditionBuilder) (ExprNode, error) {
	if len(c.conditionList) < 1 {
		return ExprNode{}, fmt.Errorf("Invalid ConditionBuilder. Expected at least 1 Condition")
	}

	if len(c.operandList) != 0 {
		return ExprNode{}, fmt.Errorf("Invalid ConditionBuilder. Expected 0 Operands")
	}

	conditionExprNodes := make([]ExprNode, 0, len(c.conditionList))
	for _, cond := range c.conditionList {
		exprNodes, err := cond.buildCondition()
		if err != nil {
			return ExprNode{}, err
		}
		conditionExprNodes = append(conditionExprNodes, exprNodes)
	}

	ret := ExprNode{
		children: conditionExprNodes,
	}

	// create a string with escaped characters to substitute them with proper
	// aliases during runtime
	for ind := range c.conditionList {
		ret.fmtExpr += "($c)"
		if ind != len(c.conditionList)-1 {
			switch c.Mode {
			case AndCond:
				ret.fmtExpr += " AND "
			}
		}
	}

	return ret, nil
}
