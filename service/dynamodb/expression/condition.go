package expression

import (
	"fmt"
	"strings"
)

// ConditionMode will specify the types of the struct ConditionBuilder,
// representing the different types of Conditions (i.e. And, Or, Between, ...)
type ConditionMode int

const (
	// UnsetCond will catch errors if users make an empty ConditionBuilder
	UnsetCond ConditionMode = iota
	// EqualCond will represent the Equal Clause ConditionBuilder
	EqualCond
	// AndCond will represent the And Clause ConditionBuilder
	AndCond
	// BetweenCond will represent the Between ConditionBuilder
	BetweenCond
)

// String will satisfy the Stringer interface in order for Error outputs to be
// more readable.
func (cm ConditionMode) String() string {
	switch cm {
	case UnsetCond:
		return "UnsetCond"
	case EqualCond:
		return "EqualCond"
	case AndCond:
		return "AndCond"
	case BetweenCond:
		return "BetweenCond"
	default:
		return "no matching ConditionMode"
	}
}

// ConditionBuilder will represent the ConditionExpressions in DynamoDB. It is
// composed of operands (OperandBuilder) and other conditions (ConditionBuilder)
// There are many different types of conditions, specified by ConditionMode.
// Users will be able to call the BuildExpression() method on a ConditionBuilder
// to create an Expression which can then be used for operation inputs into
// DynamoDB.
// More Information at: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ConditionExpressions.html
type ConditionBuilder struct {
	operandList   []OperandBuilder
	conditionList []ConditionBuilder
	Mode          ConditionMode
}

// Equal will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Equal(expression.Path("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)	// Used in another condition
//     expression, err := condition.BuildExpression()	// Used to make an Expression
func Equal(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		Mode:        EqualCond,
	}
}

// Equal will create a ConditionBuilder. This will be the method for PathBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Equal(expression.Path("foo"), expression.Value(5))
//     condition := expression.Path("foo").Equal(expression.Value(5))
func (p PathBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(p, right)
}

// Equal will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Equal(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).Equal(expression.Value(5))
func (v ValueBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(v, right)
}

// Equal will create a ConditionBuilder. This will be the method for SizeBuilder
//
// Example:
//
//     The following produce equivalent conditions:
//     condition := expression.Equal(expression.Path("foo").Size(), expression.Value(5))
//     condition := expression.Path("foo").Size().Equal(expression.Value(5))
func (s SizeBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(s, right)
}

// And will create a ConditionBuilder with more than two other Conditions as
// children, representing logical statements that will be logically ANDed
// together. The resulting ConditionBuilder can be used to build other
// Conditions or to create an Expression to be used in an operation input. This
// will be the function call.
//
// Example:
//
//     condition1 := expression.Equal(expression.Path("foo"), expression.Value(5))
//     condition2 := expression.Less(expression.Path("bar"), expression.Value(2010))
//     condition3 := expression.Path("baz").Between(expression.Value(2), expression.Value(10))
//     andCondition := expression.And(condition1, condition2, condition3)
//
//     anotherCondition := expression.Not(andCondition)		// Used in another condition
//     expression, err := andCondition.BuildExpression()	// Used to make an Expression
func And(left, right ConditionBuilder, other ...ConditionBuilder) ConditionBuilder {
	other = append([]ConditionBuilder{left, right}, other...)
	return ConditionBuilder{
		conditionList: other,
		Mode:          AndCond,
	}
}

// And will create a ConditionBuilder. This will be the method signature
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.And(condition1, condition2, condition3)
//     condition := condition1.And(condition2, condition3)
func (cond ConditionBuilder) And(right ConditionBuilder, other ...ConditionBuilder) ConditionBuilder {
	return And(cond, right, other...)
}

// Between will create a ConditionBuilder with three operands as children, the
// first operand representing the operand being compared, the second operand
// representing the lower bound value of the first operand, and the third
// operand representing the upper bound value of the first operand. The
// resulting ConditionBuilder can be used to build other Conditions or to create
// an Expression to be used in an operation input. This will be the function
// call.
//
// Example:
//
//     condition := expression.Between(expression.Path("foo"), expression.Value(2), expression.Value(6))
//
//     anotherCondition := expression.Not(condition)	// Used in another condition
//     expression, err := condition.BuildExpression()	// Used to make an Expression
func Between(ope, lower, upper OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{ope, lower, upper},
		Mode:        BetweenCond,
	}
}

// Between will create a ConditionBuilder. This will be the method signature for
// PathBuilders.
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Between(operand1, operand2, operand3)
//     condition := operand1.Between(operand2, operand3)
func (p PathBuilder) Between(lower, upper OperandBuilder) ConditionBuilder {
	return Between(p, lower, upper)
}

// Between will create a ConditionBuilder. This will be the method signature for
// ValueBuilders.
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Between(operand1, operand2, operand3)
//     condition := operand1.Between(operand2, operand3)
func (v ValueBuilder) Between(lower, upper OperandBuilder) ConditionBuilder {
	return Between(v, lower, upper)
}

// Between will create a ConditionBuilder. This will be the method signature for
// SizeBuilders.
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Between(operand1, operand2, operand3)
//     condition := operand1.Between(operand2, operand3)
func (s SizeBuilder) Between(lower, upper OperandBuilder) ConditionBuilder {
	return Between(s, lower, upper)
}

// BuildExpression will take an ConditionBuilder as input and output an
// Expression which can be used in DynamoDB operational inputs (i.e.
// UpdateItemInput, DeleteItemInput, etc) In the future, the Expression struct
// can be used in some injection method into the input structs.
//
// Example:
//
//     expr, err := someCondition.BuildExpression()
//
//     deleteInput := dynamodb.DeleteItemInput{
//       ConditionExpression:       aws.String(expr.Expression),
// 	     ExpressionAttributeNames:  expr.Names,
//       ExpressionAttributeValues: expr.Values,
//       Key: map[string]*dynamodb.AttributeValue{
//         "PartitionKey": &dynamodb.AttributeValue{
//           S: aws.String("SomeKey"),
//         },
//       },
//       TableName: aws.String("SomeTable"),
//     }
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

// buildCondition will build a tree structure of ExprNodes based on the tree
// structure of the input ConditionBuilder's child Conditions/Operands.
func (cond ConditionBuilder) buildCondition() (ExprNode, error) {
	switch cond.Mode {
	case EqualCond:
		return compareBuildCondition(cond)
	case AndCond:
		return compoundBuildCondition(cond)
	case BetweenCond:
		return betweenBuildCondition(cond)
	}
	return ExprNode{}, fmt.Errorf("buildCondition error: no matching ConditionMode to %v", cond.Mode)
}

// compareBuildCondition is the function to make ExprNodes from Compare
// ConditionBuilders. There will first be checks to make sure that the input
// ConditionBuilder has the correct format.
func compareBuildCondition(c ConditionBuilder) (ExprNode, error) {
	childNodes, err := c.buildChildNodes()
	if err != nil {
		return ExprNode{}, err
	}
	ret := ExprNode{
		children: childNodes,
	}

	// Create a string with special characters that can be substituted later: $c
	switch c.Mode {
	case EqualCond:
		ret.fmtExpr = "$c = $c"
	}

	return ret, nil
}

// compoundBuildCondition is the function to make ExprNodes from And/Or
// ConditionBuilders. There will first be checks to make sure that the input
// ConditionBuilder has the correct format.
func compoundBuildCondition(c ConditionBuilder) (ExprNode, error) {
	childNodes, err := c.buildChildNodes()
	if err != nil {
		return ExprNode{}, err
	}
	ret := ExprNode{
		children: childNodes,
	}

	// create a string with escaped characters to substitute them with proper
	// aliases during runtime
	var mode string
	switch c.Mode {
	case AndCond:
		mode = " AND "
	}

	ret.fmtExpr = "($c)" + strings.Repeat(mode+"($c)", len(c.conditionList)-1)

	return ret, nil
}

// betweenBuildCondition is the function to make ExprNodes from Between
// ConditionBuilders. There will first be checks to make sure that the input
// ConditionBuilder has the correct format.
func betweenBuildCondition(c ConditionBuilder) (ExprNode, error) {
	childNodes, err := c.buildChildNodes()
	if err != nil {
		return ExprNode{}, err
	}
	ret := ExprNode{
		children: childNodes,
	}

	// Create a string with special characters that can be substituted later: $c
	ret.fmtExpr = "$c BETWEEN $c AND $c"

	return ret, nil
}

// buildChildNodes will create the list of the child ExprNodes. This avoids
// duplication of code amongst the various buildConditions.
func (cond ConditionBuilder) buildChildNodes() ([]ExprNode, error) {
	var childNodes []ExprNode

	childNodes = make([]ExprNode, 0, len(cond.conditionList)+len(cond.operandList))
	for _, condition := range cond.conditionList {
		en, err := condition.buildCondition()
		if err != nil {
			return []ExprNode{}, err
		}
		childNodes = append(childNodes, en)
	}

	for _, ope := range cond.operandList {
		en, err := ope.BuildOperand()
		if err != nil {
			return []ExprNode{}, err
		}
		childNodes = append(childNodes, en)
	}

	return childNodes, nil
}
