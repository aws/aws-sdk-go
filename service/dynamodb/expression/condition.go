package expression

import "fmt"

// ConditionMode will specify the types of the struct ConditionBuilder
type ConditionMode int

const (
	// EqualCond will represent the Equal Clause ConditionBuilder
	EqualCond ConditionMode = iota + 1
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

// buildCondition will iterate over the tree of ConditionBuilders and
// OperandBuilders and build a tree of ExprNodes
func (cond ConditionBuilder) buildCondition() (ExprNode, error) {
	switch cond.Mode {
	case EqualCond:
		return compareBuildCondition(cond)
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

	operandExprNodes := make([]ExprNode, 0)
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

	switch c.Mode {
	case EqualCond:
		ret.fmtExpr = "(%c) = (%c)"
	}

	return ret, nil
}
