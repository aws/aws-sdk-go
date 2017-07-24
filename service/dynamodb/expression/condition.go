package expression

import "github.com/aws/aws-sdk-go/service/dynamodb"

// ConditionBuilder blah
type ConditionBuilder interface {
	BuildCondition() (Expression, error)
}

// Compare

// CompareBuilder implements both FilterBuilder and ConditionBuilder
// It will be the output of the following
// - Equal()
// - NotEqual()
// - Less()
// - LessEqual()
// - Greater()
// - GreaterEqual()
type CompareBuilder struct {
	Left  OperandBuilder
	Right OperandBuilder
	Type  string
}

// Equal

// Equal will create a CompareBuilder. This will be the method PathBuilder.
func (p PathBuilder) Equal(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  p,
		Right: right,
		Type:  "=",
	}
}

// Equal will create a CompareBuilder. This will be the method ValueBuilder.
func (v ValueBuilder) Equal(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  v,
		Right: right,
		Type:  "=",
	}
}

// Equal will create a CompareBuilder. This will be the method SizeBuilder.
func (s SizeBuilder) Equal(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  s,
		Right: right,
		Type:  "=",
	}
}

// NotEqual

// NotEqual will create a CompareBuilder. This will be the method PathBuilder.
func (p PathBuilder) NotEqual(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  p,
		Right: right,
		Type:  "<>",
	}
}

// NotEqual will create a CompareBuilder. This will be the method ValueBuilder.
func (v ValueBuilder) NotEqual(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  v,
		Right: right,
		Type:  "<>",
	}
}

// NotEqual will create a CompareBuilder. This will be the method SizeBuilder.
func (s SizeBuilder) NotEqual(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  s,
		Right: right,
		Type:  "<>",
	}
}

// Less

// Less will create a CompareBuilder. This will be the method PathBuilder.
func (p PathBuilder) Less(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  p,
		Right: right,
		Type:  "<",
	}
}

// Less will create a CompareBuilder. This will be the method ValueBuilder.
func (v ValueBuilder) Less(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  v,
		Right: right,
		Type:  "<",
	}
}

// Less will create a CompareBuilder. This will be the method SizeBuilder.
func (s SizeBuilder) Less(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  s,
		Right: right,
		Type:  "<",
	}
}

// LessEqual

// LessEqual will create a CompareBuilder. This will be the method PathBuilder.
func (p PathBuilder) LessEqual(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  p,
		Right: right,
		Type:  "<=",
	}
}

// LessEqual will create a CompareBuilder. This will be the method ValueBuilder.
func (v ValueBuilder) LessEqual(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  v,
		Right: right,
		Type:  "<=",
	}
}

// LessEqual will create a CompareBuilder. This will be the method SizeBuilder.
func (s SizeBuilder) LessEqual(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  s,
		Right: right,
		Type:  "<=",
	}
}

// Greater

// Greater will create a CompareBuilder. This will be the method PathBuilder.
func (p PathBuilder) Greater(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  p,
		Right: right,
		Type:  ">",
	}
}

// Greater will create a CompareBuilder. This will be the method ValueBuilder.
func (v ValueBuilder) Greater(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  v,
		Right: right,
		Type:  ">",
	}
}

// Greater will create a CompareBuilder. This will be the method SizeBuilder.
func (s SizeBuilder) Greater(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  s,
		Right: right,
		Type:  ">",
	}
}

// GreaterEqual

// GreaterEqual will create a CompareBuilder. This will be the method PathBuilder.
func (p PathBuilder) GreaterEqual(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  p,
		Right: right,
		Type:  ">",
	}
}

// GreaterEqual will create a CompareBuilder. This will be the method ValueBuilder.
func (v ValueBuilder) GreaterEqual(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  v,
		Right: right,
		Type:  ">",
	}
}

// GreaterEqual will create a CompareBuilder. This will be the method SizeBuilder.
func (s SizeBuilder) GreaterEqual(right OperandBuilder) CompareBuilder {
	return CompareBuilder{
		Left:  s,
		Right: right,
		Type:  ">",
	}
}

// Size

// SizeBuilder will implement OperandBuilder thus being an Operand. This
// reflects the fact that the function Size() returns type that is used in place
// of an Operand
type SizeBuilder struct {
	path PathBuilder
}

// Size will
func (p PathBuilder) Size() SizeBuilder {
	return SizeBuilder{
		path: p,
	}
}

// BuildOperand will create
func (s SizeBuilder) BuildOperand() (Expression, error) {
	expr, err := s.path.BuildOperand()
	expr.Expression = "size (" + expr.Expression + ")"
	return expr, err
}

// BuildCondition will create the Expression represented by CompareBuilder
func (expr CompareBuilder) BuildCondition() (Expression, error) {
	left, err := expr.Left.BuildOperand()
	if err != nil {
		return Expression{}, err
	}
	right, err := expr.Right.BuildOperand()
	if err != nil {
		return Expression{}, err
	}

	if left.Names == nil && right.Names != nil {
		left.Names = make(map[string]*string)
	}
	for alias, name := range right.Names {
		left.Names[alias] = name
	}

	if left.Values == nil && right.Values != nil {
		left.Values = make(map[string]*dynamodb.AttributeValue)
	}
	for alias, value := range right.Values {
		left.Values[alias] = value
	}

	left.Expression = left.Expression + " " + expr.Type + " " + right.Expression

	return left, nil
}
