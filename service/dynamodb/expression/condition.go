package expression

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// ErrUnsetCondition is an error that is returned if BuildExpression is called
// on an empty ConditionBuilder.
var ErrUnsetCondition = awserr.New("UnsetCondition", "buildCondition error: the argument ConditionBuilder's mode is unset", nil)

// conditionMode will specify the types of the struct conditionBuilder,
// representing the different types of Conditions (i.e. And, Or, Between, ...)
type conditionMode int

const (
	// unsetCond will catch errors if users make an empty ConditionBuilder
	unsetCond conditionMode = iota
	// equalCond will represent the Equals Condition
	equalCond
	// notEqualCond will represent the Not Equals Condition
	notEqualCond
	// lessCond will represent the Less Than Condition
	lessCond
	// lessEqualCond will represent the Less Than Or Equal To Condition
	lessEqualCond
	// greaterCond will represent the Greater Than Condition
	greaterCond
	// greaterEqualCond will represent the Greater Than Or Equal To Condition
	greaterEqualCond
	// andCond will represent the Logical And Condition
	andCond
	// orCond will represent the Logical Or Condition
	orCond
	// notCond will represent the Logical Not Condition
	notCond
	// betweenCond will represent the Between Condition
	betweenCond
	// inCond will represent the In Condition
	inCond
	// attrExistsCond will represent the Attribute Exists Condition
	attrExistsCond
	// attrNotExistsCond will represent the Attribute Not Exists Condition
	attrNotExistsCond
	// attrTypeCond will represent the Attribute Type Condition
	attrTypeCond
	// beginsWithCond will represent the Begins With Condition
	beginsWithCond
	// containsCond will represent the Contains Condition
	containsCond
)

func (cm conditionMode) String() string {
	switch cm {
	case unsetCond:
		return "unsetCond"
	default:
		return "no matching conditionMode"
	}
}

// DynamoDBAttributeType will specify the type of an DynamoDB item attribute.
// This enum will be used in the AttributeType() function in order to be
// explicit about the DynamoDB type that is being checked and ensure compile
// time checks
// More Informatin at http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.OperatorsAndFunctions.html#Expressions.OperatorsAndFunctions.Functions
type DynamoDBAttributeType string

const (
	// String will represent the DynamoDB String type
	String DynamoDBAttributeType = "S"
	// StringSet will represent the DynamoDB String Set type
	StringSet = "SS"
	// Number will represent the DynamoDB Number type
	Number = "N"
	// NumberSet will represent the DynamoDB Number Set type
	NumberSet = "NS"
	// Binary will represent the DynamoDB Binary type
	Binary = "B"
	// BinarySet will represent the DynamoDB Binary Set type
	BinarySet = "BS"
	// Boolean will represent the DynamoDB Boolean type
	Boolean = "BOOL"
	// Null will represent the DynamoDB Null type
	Null = "NULL"
	// List will represent the DynamoDB List type
	List = "L"
	// Map will represent the DynamoDB Map type
	Map = "M"
)

// ConditionBuilder will represent Condition Expressions and Filter Expressions
// in DynamoDB. It is composed of operands (OperandBuilder) and other conditions
// (ConditionBuilder). There are many different types of conditions, specified
// by ConditionMode. Users will be able to call the BuildExpression() method on
// a ConditionBuilder to create an Expression which can then be used for
// operation inputs into DynamoDB. Since Filter Expressions support all the same
// functions and formats as Condition Expressions, ConditionBuilders will
// satisfy both Expressions.
// More Information at: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ConditionExpressions.html
// More Information on Filter Expressions: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Query.html#Query.FilterExpression
type ConditionBuilder struct {
	operandList   []OperandBuilder
	conditionList []ConditionBuilder
	mode          conditionMode
}

// Equal will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Equal(expression.Name("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func Equal(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        equalCond,
	}
}

// Equal will create a ConditionBuilder. This will be the method for NameBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Equal(expression.Name("foo"), expression.Value(5))
//     condition := expression.Name("foo").Equal(expression.Value(5))
func (nameBuilder NameBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(nameBuilder, right)
}

// Equal will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Equal(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).Equal(expression.Value(5))
func (valueBuilder ValueBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(valueBuilder, right)
}

// Equal will create a ConditionBuilder. This will be the method for SizeBuilder
//
// Example:
//
//     The following produces equivalent conditions:
//     condition := expression.Equal(expression.Name("foo").Size(), expression.Value(5))
//     condition := expression.Name("foo").Size().Equal(expression.Value(5))
func (sizeBuilder SizeBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(sizeBuilder, right)
}

// NotEqual will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.NotEqual(expression.Name("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func NotEqual(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        notEqualCond,
	}
}

// NotEqual will create a ConditionBuilder. This will be the method for
// NameBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.NotEqual(expression.Name("foo"), expression.Value(5))
//     condition := expression.Name("foo").NotEqual(expression.Value(5))
func (nameBuilder NameBuilder) NotEqual(right OperandBuilder) ConditionBuilder {
	return NotEqual(nameBuilder, right)
}

// NotEqual will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.NotEqual(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).NotEqual(expression.Value(5))
func (valueBuilder ValueBuilder) NotEqual(right OperandBuilder) ConditionBuilder {
	return NotEqual(valueBuilder, right)
}

// NotEqual will create a ConditionBuilder. This will be the method for
// SizeBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.NotEqual(expression.Name("foo").Size(), expression.Value(5))
//     condition := expression.Name("foo").Size().NotEqual(expression.Value(5))
func (sizeBuilder SizeBuilder) NotEqual(right OperandBuilder) ConditionBuilder {
	return NotEqual(sizeBuilder, right)
}

// Less will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Less(expression.Name("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func Less(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        lessCond,
	}
}

// Less will create a ConditionBuilder. This will be the method for NameBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Less(expression.Name("foo"), expression.Value(5))
//     condition := expression.Name("foo").Less(expression.Value(5))
func (nameBuilder NameBuilder) Less(right OperandBuilder) ConditionBuilder {
	return Less(nameBuilder, right)
}

// Less will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Less(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).Less(expression.Value(5))
func (valueBuilder ValueBuilder) Less(right OperandBuilder) ConditionBuilder {
	return Less(valueBuilder, right)
}

// Less will create a ConditionBuilder. This will be the method for SizeBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Less(expression.Name("foo").Size(), expression.Value(5))
//     condition := expression.Name("foo").Size().Less(expression.Value(5))
func (sizeBuilder SizeBuilder) Less(right OperandBuilder) ConditionBuilder {
	return Less(sizeBuilder, right)
}

// LessEqual will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.LessEqual(expression.Name("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func LessEqual(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        lessEqualCond,
	}
}

// LessEqual will create a ConditionBuilder. This will be the method for NameBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.LessEqual(expression.Name("foo"), expression.Value(5))
//     condition := expression.Name("foo").LessEqual(expression.Value(5))
func (nameBuilder NameBuilder) LessEqual(right OperandBuilder) ConditionBuilder {
	return LessEqual(nameBuilder, right)
}

// LessEqual will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.LessEqual(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).LessEqual(expression.Value(5))
func (valueBuilder ValueBuilder) LessEqual(right OperandBuilder) ConditionBuilder {
	return LessEqual(valueBuilder, right)
}

// LessEqual will create a ConditionBuilder. This will be the method for SizeBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.LessEqual(expression.Name("foo").Size(), expression.Value(5))
//     condition := expression.Name("foo").Size().LessEqual(expression.Value(5))
func (sizeBuilder SizeBuilder) LessEqual(right OperandBuilder) ConditionBuilder {
	return LessEqual(sizeBuilder, right)
}

// Greater will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Greater(expression.Name("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func Greater(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        greaterCond,
	}
}

// Greater will create a ConditionBuilder. This will be the method for NameBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Greater(expression.Name("foo"), expression.Value(5))
//     condition := expression.Name("foo").Greater(expression.Value(5))
func (nameBuilder NameBuilder) Greater(right OperandBuilder) ConditionBuilder {
	return Greater(nameBuilder, right)
}

// Greater will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Greater(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).Greater(expression.Value(5))
func (valueBuilder ValueBuilder) Greater(right OperandBuilder) ConditionBuilder {
	return Greater(valueBuilder, right)
}

// Greater will create a ConditionBuilder. This will be the method for SizeBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Greater(expression.Name("foo").Size(), expression.Value(5))
//     condition := expression.Name("foo").Size().Greater(expression.Value(5))
func (sizeBuilder SizeBuilder) Greater(right OperandBuilder) ConditionBuilder {
	return Greater(sizeBuilder, right)
}

// GreaterEqual will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.GreaterEqual(expression.Name("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func GreaterEqual(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        greaterEqualCond,
	}
}

// GreaterEqual will create a ConditionBuilder. This will be the method for NameBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.GreaterEqual(expression.Name("foo"), expression.Value(5))
//     condition := expression.Name("foo").GreaterEqual(expression.Value(5))
func (nameBuilder NameBuilder) GreaterEqual(right OperandBuilder) ConditionBuilder {
	return GreaterEqual(nameBuilder, right)
}

// GreaterEqual will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.GreaterEqual(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).GreaterEqual(expression.Value(5))
func (valueBuilder ValueBuilder) GreaterEqual(right OperandBuilder) ConditionBuilder {
	return GreaterEqual(valueBuilder, right)
}

// GreaterEqual will create a ConditionBuilder. This will be the method for SizeBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.GreaterEqual(expression.Name("foo").Size(), expression.Value(5))
//     condition := expression.Name("foo").Size().GreaterEqual(expression.Value(5))
func (sizeBuilder SizeBuilder) GreaterEqual(right OperandBuilder) ConditionBuilder {
	return GreaterEqual(sizeBuilder, right)
}

// And will create a ConditionBuilder with more than two other Conditions as
// children, representing logical statements that will be logically ANDed
// together. The resulting ConditionBuilder can be used to build other
// Conditions or to create an Expression to be used in an operation input. This
// will be the function call.
//
// Example:
//
//     condition1 := expression.Equal(expression.Name("foo"), expression.Value(5))
//     condition2 := expression.Less(expression.Name("bar"), expression.Value(2010))
//     condition3 := expression.Name("baz").Between(expression.Value(2), expression.Value(10))
//     andCondition := expression.And(condition1, condition2, condition3)
//
//     anotherCondition := expression.Not(andCondition)  // Used in another condition
//     expression, err := andCondition.BuildExpression() // Used to make an Expression
func And(left, right ConditionBuilder, other ...ConditionBuilder) ConditionBuilder {
	other = append([]ConditionBuilder{left, right}, other...)
	return ConditionBuilder{
		conditionList: other,
		mode:          andCond,
	}
}

// And will create a ConditionBuilder. This will be the method signature
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.And(condition1, condition2, condition3)
//     condition := condition1.And(condition2, condition3)
func (conditionBuilder ConditionBuilder) And(right ConditionBuilder, other ...ConditionBuilder) ConditionBuilder {
	return And(conditionBuilder, right, other...)
}

// Or will create a ConditionBuilder with more than two other Conditions as
// children, representing logical statements that will be logically ORed
// together. The resulting ConditionBuilder can be used to build other
// Conditions or to create an Expression to be used in an operation input. This
// will be the function call.
//
// Example:
//
//     condition1 := expression.Equal(expression.Name("foo"), expression.Value(5))
//     condition2 := expression.Less(expression.Name("bar"), expression.Value(2010))
//     condition3 := expression.Name("baz").Between(expression.Value(2), expression.Value(10))
//     orCondition := expression.Or(condition1, condition2, condition3)
//
//     anotherCondition := expression.Not(orCondition)  // Used in another condition
//     expression, err := orCondition.BuildExpression() // Used to make an Expression
func Or(left, right ConditionBuilder, other ...ConditionBuilder) ConditionBuilder {
	other = append([]ConditionBuilder{left, right}, other...)
	return ConditionBuilder{
		conditionList: other,
		mode:          orCond,
	}
}

// Or will create a ConditionBuilder. This will be the method signature
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Or(condition1, condition2, condition3)
//     condition := condition1.Or(condition2, condition3)
func (conditionBuilder ConditionBuilder) Or(right ConditionBuilder, other ...ConditionBuilder) ConditionBuilder {
	return Or(conditionBuilder, right, other...)
}

// Not will create a ConditionBuilder with one Conditions as a child,
// representing the logical statements that will be logically negated. The
// resulting ConditionBuilder can be used to build other Conditions or to create
// an Expression to be used in an operation input. This will be the function
// call.
//
// Example:
//
//     condition := expression.Equal(expression.Name("foo"), expression.Value(5))
//     notCondition := expression.Or(condition)
//
//     anotherCondition := expression.Not(notCondition)  // Used in another condition
//     expression, err := notCondition.BuildExpression() // Used to make an Expression
func Not(conditionBuilder ConditionBuilder) ConditionBuilder {
	return ConditionBuilder{
		conditionList: []ConditionBuilder{conditionBuilder},
		mode:          notCond,
	}
}

// Not will create a ConditionBuilder. This will be the method signature
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Not(condition)
//     condition := condition.Not()
func (conditionBuilder ConditionBuilder) Not() ConditionBuilder {
	return Not(conditionBuilder)
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
//     condition := expression.Between(expression.Name("foo"), expression.Value(2), expression.Value(6))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func Between(ope, lower, upper OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{ope, lower, upper},
		mode:        betweenCond,
	}
}

// Between will create a ConditionBuilder. This will be the method signature for
// NameBuilders.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Between(operand1, operand2, operand3)
//     condition := operand1.Between(operand2, operand3)
func (nameBuilder NameBuilder) Between(lower, upper OperandBuilder) ConditionBuilder {
	return Between(nameBuilder, lower, upper)
}

// Between will create a ConditionBuilder. This will be the method signature for
// ValueBuilders.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Between(operand1, operand2, operand3)
//     condition := operand1.Between(operand2, operand3)
func (valueBuilder ValueBuilder) Between(lower, upper OperandBuilder) ConditionBuilder {
	return Between(valueBuilder, lower, upper)
}

// Between will create a ConditionBuilder. This will be the method signature for
// SizeBuilders.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Between(operand1, operand2, operand3)
//     condition := operand1.Between(operand2, operand3)
func (sizeBuilder SizeBuilder) Between(lower, upper OperandBuilder) ConditionBuilder {
	return Between(sizeBuilder, lower, upper)
}

// In will create a ConditionBuilder with two or more operands as children, the
// first operand representing the operand being compared and the rest of the
// operands representing a set in which the first operand either belongs to or
// not. The argument must have at least two operands. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Between(expression.Name("foo"), expression.Value(2), expression.Value(6))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func In(left, right OperandBuilder, other ...OperandBuilder) ConditionBuilder {
	other = append([]OperandBuilder{left, right}, other...)
	return ConditionBuilder{
		operandList: other,
		mode:        inCond,
	}
}

// In will create a ConditionBuilder. This will be the method signature for
// NameBuilders.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.In(operand1, operand2, operand3)
//     condition := operand1.In(operand2, operand3)
func (nameBuilder NameBuilder) In(right OperandBuilder, other ...OperandBuilder) ConditionBuilder {
	return In(nameBuilder, right, other...)
}

// In will create a ConditionBuilder. This will be the method signature for
// ValueBuilders.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.In(operand1, operand2, operand3)
//     condition := operand1.In(operand2, operand3)
func (valueBuilder ValueBuilder) In(right OperandBuilder, other ...OperandBuilder) ConditionBuilder {
	return In(valueBuilder, right, other...)
}

// In will create a ConditionBuilder. This will be the method signature for
// SizeBuilders.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.In(operand1, operand2, operand3)
//     condition := operand1.In(operand2, operand3)
func (sizeBuilder SizeBuilder) In(right OperandBuilder, other ...OperandBuilder) ConditionBuilder {
	return In(sizeBuilder, right, other...)
}

// AttributeExists will create a ConditionBuilder with a name as a child. The
// function will return true if the item attribute described by the name exists.
// The resulting ConditionBuilder can be used to build other Conditions or to
// create an Expression to be used in an operation input. This will be the
// function call.
//
// Example:
//
//     condition := expression.AttributeExists(Name("foo"))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func AttributeExists(nameBuilder NameBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{nameBuilder},
		mode:        attrExistsCond,
	}
}

// AttributeExists will create a ConditionBuilder. AttributeExists will only
// have a method for NameBuilders since that is the only valid operand that the
// function can be called on.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.AttributeExists(Name("foo"))
//     condition := Name("foo").AttributeExists()
func (nameBuilder NameBuilder) AttributeExists() ConditionBuilder {
	return AttributeExists(nameBuilder)
}

// AttributeNotExists will create a ConditionBuilder with a name as a child. The
// function will return true if the item attribute described by the name does
// not exist. The resulting ConditionBuilder can be used to build other
// Conditions or to create an Expression to be used in an operation input. This
// will be the function call.
//
// Example:
//
//     condition := expression.AttributeNotExists(expression.Name("foo"))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func AttributeNotExists(nameBuilder NameBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{nameBuilder},
		mode:        attrNotExistsCond,
	}
}

// AttributeNotExists will create a ConditionBuilder. AttributeNotExists will
// only have a method for NameBuilders since that is the only valid operand that
// the function can be called on.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.AttributeNotExists(expression.Name("foo"))
//     condition := expression.Name("foo").AttributeNotExists()
func (nameBuilder NameBuilder) AttributeNotExists() ConditionBuilder {
	return AttributeNotExists(nameBuilder)
}

// AttributeType will create a ConditionBuilder with a name and a value as a
// child. The name will represent the item attribute being compared. The value
// will be a string corresponding to the argument DynamoDBAttributeType. The
// function will return true if the item attribute described by the name is the
// type specified by DynamoDBAttributeType. The resulting ConditionBuilder can
// be used to build other Conditions or to create an Expression to be used in an
// operation input. This will be the function call.
//
// Example:
//
//     condition := expression.AttributeType(Name("foo"), expression.StringSet)
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func AttributeType(nameBuilder NameBuilder, attributeType DynamoDBAttributeType) ConditionBuilder {
	v := ValueBuilder{
		value: string(attributeType),
	}
	return ConditionBuilder{
		operandList: []OperandBuilder{nameBuilder, v},
		mode:        attrTypeCond,
	}
}

// AttributeType will create a ConditionBuilder. AttributeType will only have a
// method for NameBuilders since that is the only valid operand that the
// function can be called on.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.AttributeType(expression.Name("foo"), expression.Number)
//     condition := expression.Name("foo").AttributeType(expression.Number)
func (nameBuilder NameBuilder) AttributeType(attributeType DynamoDBAttributeType) ConditionBuilder {
	return AttributeType(nameBuilder, attributeType)
}

// BeginsWith will create a ConditionBuilder with a name and a value as
// children. The name will represent the name to the item attribute being
// compared. The value will represent the substring in which the item attribute
// will be compared with. The function will return true if the item attribute
// specified by the name starts with the substring. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.BeginsWith(Name("foo"), "bar")
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func BeginsWith(nameBuilder NameBuilder, substr string) ConditionBuilder {
	v := ValueBuilder{
		value: substr,
	}
	return ConditionBuilder{
		operandList: []OperandBuilder{nameBuilder, v},
		mode:        beginsWithCond,
	}
}

// BeginsWith will create a ConditionBuilder. BeginsWith will only have a method
// for NameBuilders since that is the only valid operand that the function can
// be called on.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.BeginsWith(expression.Name("foo"), "bar")
//     condition := expression.Name("foo").BeginsWith("bar")
func (nameBuilder NameBuilder) BeginsWith(substr string) ConditionBuilder {
	return BeginsWith(nameBuilder, substr)
}

// Contains will create a ConditionBuilder with a name and a value as
// children. The name will represent the name to the item attribute being
// compared. The item attribute MUST be a String or a Set. The value will
// represent the string in which the item attribute will be compared with. The
// function will return true if the item attribute specified by the name
// contains the substring specified by the value or if the item attribute is a
// set that contains the string specified by the value. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Contains(Name("foo"), "bar")
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func Contains(nameBuilder NameBuilder, substr string) ConditionBuilder {
	v := ValueBuilder{
		value: substr,
	}
	return ConditionBuilder{
		operandList: []OperandBuilder{nameBuilder, v},
		mode:        containsCond,
	}
}

// Contains will create a ConditionBuilder. Contains will only have a method
// for NameBuilders since that is the only valid operand that the function can
// be called on.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Contains(expression.Name("foo"), "bar")
//     condition := expression.Name("foo").Contains("bar")
func (nameBuilder NameBuilder) Contains(substr string) ConditionBuilder {
	return Contains(nameBuilder, substr)
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
//       ConditionExpression:       expr.ConditionExpression(),
// 	     ExpressionAttributeNames:  expr.Names(),
//       ExpressionAttributeValues: expr.Values(),
//       Key: map[string]*dynamodb.AttributeValue{
//         "PartitionKey": &dynamodb.AttributeValue{
//           S: aws.String("SomeKey"),
//         },
//       },
//       TableName: aws.String("SomeTable"),
//     }
func (conditionBuilder ConditionBuilder) BuildExpression() (Expression, error) {
	return Expression{
		expressionMap: map[string]TreeBuilder{
			"condition": conditionBuilder,
		},
	}, nil
}

// BuildTree will build a tree structure of ExprNodes based on the
// tree structure of the input ConditionBuilder's child ConditionBuilders and
// OperandBuilders.
func (conditionBuilder ConditionBuilder) BuildTree() (ExprNode, error) {
	childNodes, err := conditionBuilder.buildChildNodes()
	if err != nil {
		return ExprNode{}, err
	}
	ret := ExprNode{
		children: childNodes,
	}

	switch conditionBuilder.mode {
	case equalCond, notEqualCond, lessCond, lessEqualCond, greaterCond, greaterEqualCond:
		return compareBuildCondition(conditionBuilder.mode, ret)
	case andCond, orCond:
		return compoundBuildCondition(conditionBuilder, ret)
	case notCond:
		return notBuildCondition(ret)
	case betweenCond:
		return betweenBuildCondition(ret)
	case inCond:
		return inBuildCondition(conditionBuilder, ret)
	case attrExistsCond:
		return attrExistsBuildCondition(ret)
	case attrNotExistsCond:
		return attrNotExistsBuildCondition(ret)
	case attrTypeCond:
		return attrTypeBuildCondition(ret)
	case beginsWithCond:
		return beginsWithBuildCondition(ret)
	case containsCond:
		return containsBuildCondition(ret)
	case unsetCond:
		return ExprNode{}, ErrUnsetCondition
	default:
		return ExprNode{}, fmt.Errorf("buildCondition error: unsupported mode: %v", conditionBuilder.mode)
	}
}

// compareBuildCondition is the function to make ExprNodes from Compare
// ConditionBuilders. compareBuildCondition will only be called by the
// buildCondition method. This function assumes that the argument
// ConditionBuilder has the right format.
func compareBuildCondition(cm conditionMode, en ExprNode) (ExprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	switch cm {
	case equalCond:
		en.fmtExpr = "$c = $c"
	case notEqualCond:
		en.fmtExpr = "$c <> $c"
	case lessCond:
		en.fmtExpr = "$c < $c"
	case lessEqualCond:
		en.fmtExpr = "$c <= $c"
	case greaterCond:
		en.fmtExpr = "$c > $c"
	case greaterEqualCond:
		en.fmtExpr = "$c >= $c"
	}

	return en, nil
}

// compoundBuildCondition is the function to make ExprNodes from And/Or
// ConditionBuilders. compoundBuildCondition will only be called by the
// buildCondition method. This function assumes that the argument
// ConditionBuilder has the right format.
func compoundBuildCondition(c ConditionBuilder, en ExprNode) (ExprNode, error) {
	// create a string with escaped characters to substitute them with proper
	// aliases during runtime
	var mode string
	switch c.mode {
	case andCond:
		mode = " AND "
	case orCond:
		mode = " OR "
	}
	en.fmtExpr = "($c)" + strings.Repeat(mode+"($c)", len(c.conditionList)-1)

	return en, nil
}

// notBuildCondition is the function to make ExprNodes from Not
// ConditionBuilders. notBuildCondition will only be called by the
// buildCondition method. This function assumes that the argument
// ConditionBuilder has the right format.
func notBuildCondition(en ExprNode) (ExprNode, error) {
	// create a string with escaped characters to substitute them with proper
	// aliases during runtime
	en.fmtExpr = "NOT ($c)"

	return en, nil
}

// betweenBuildCondition is the function to make ExprNodes from Between
// ConditionBuilders. BuildCondition will only be called by the
// buildCondition method. This function assumes that the argument
// ConditionBuilder has the right format.
func betweenBuildCondition(en ExprNode) (ExprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	en.fmtExpr = "$c BETWEEN $c AND $c"

	return en, nil
}

// inBuildCondition is the function to make ExprNodes from In
// ConditionBuilders. inBuildCondition will only be called by the
// buildCondition method. This function assumes that the argument
// ConditionBuilder has the right format.
func inBuildCondition(c ConditionBuilder, en ExprNode) (ExprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	en.fmtExpr = "$c IN ($c" + strings.Repeat(", $c", len(c.operandList)-2) + ")"

	return en, nil
}

// attrExistsBuildCondition is the function to make ExprNodes from
// AttrExistsCond ConditionBuilders. attrExistsBuildCondition will only be
// called by the buildCondition method. This function assumes that the argument
// ConditionBuilder has the right format.
func attrExistsBuildCondition(en ExprNode) (ExprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	en.fmtExpr = "attribute_exists ($c)"

	return en, nil
}

// attrNotExistsBuildCondition is the function to make ExprNodes from
// AttrNotExistsCond ConditionBuilders. attrNotExistsBuildCondition will only be
// called by the buildCondition method. This function assumes that the argument
// ConditionBuilder has the right format.
func attrNotExistsBuildCondition(en ExprNode) (ExprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	en.fmtExpr = "attribute_not_exists ($c)"

	return en, nil
}

// attrTypeBuildCondition is the function to make ExprNodes from AttrTypeCond
// ConditionBuilders. attrTypeBuildCondition will only be called by the
// buildCondition method. This function assumes that the argument
// ConditionBuilder has the right format.
func attrTypeBuildCondition(en ExprNode) (ExprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	en.fmtExpr = "attribute_type ($c, $c)"

	return en, nil
}

// beginsWithBuildCondition is the function to make ExprNodes from
// BeginsWithCond ConditionBuilders. beginsWithBuildCondition will only be
// called by the buildCondition method. This function assumes that the argument
// ConditionBuilder has the right format.
func beginsWithBuildCondition(en ExprNode) (ExprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	en.fmtExpr = "begins_with ($c, $c)"

	return en, nil
}

// containsBuildCondition is the function to make ExprNodes from
// ContainsCond ConditionBuilders. containsBuildCondition will only be
// called by the buildCondition method. This function assumes that the argument
// ConditionBuilder has the right format.
func containsBuildCondition(en ExprNode) (ExprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	en.fmtExpr = "contains ($c, $c)"

	return en, nil
}

// buildChildNodes will create the list of the child ExprNodes. This avoids
// duplication of code amongst the various buildConditions.
func (conditionBuilder ConditionBuilder) buildChildNodes() ([]ExprNode, error) {
	childNodes := make([]ExprNode, 0, len(conditionBuilder.conditionList)+len(conditionBuilder.operandList))
	for _, condition := range conditionBuilder.conditionList {
		en, err := condition.BuildTree()
		if err != nil {
			return []ExprNode{}, err
		}
		childNodes = append(childNodes, en)
	}
	for _, ope := range conditionBuilder.operandList {
		en, err := ope.BuildOperand()
		if err != nil {
			return []ExprNode{}, err
		}
		childNodes = append(childNodes, en)
	}

	return childNodes, nil
}
