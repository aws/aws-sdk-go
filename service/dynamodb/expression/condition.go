package expression

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// ErrUnsetCondition is an error that is returned if an empty ConditionBuilder
// is used as a part of an ExpressionBuilder.
var ErrUnsetCondition = awserr.New("UnsetCondition", "buildTree error: the argument ConditionBuilder's mode is unset", nil)

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
	// lessThanCond will represent the LessThan Condition
	lessThanCond
	// lessThanEqualCond will represent the LessThanOrEqual Condition
	lessThanEqualCond
	// greaterThanCond will represent the GreaterThan Condition
	greaterThanCond
	// greaterThanEqualCond will represent the GreaterThanEqual Condition
	greaterThanEqualCond
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
// by conditionMode. ConditionBuilders will be the building blocks of Builders.
// Since Filter Expressions support all the same functions and formats as
// Condition Expressions, ConditionBuilders will represent both types of
// Expressions.
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
// Builder to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Equal(expression.Name("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
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
func (nb NameBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(nb, right)
}

// Equal will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Equal(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).Equal(expression.Value(5))
func (vb ValueBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(vb, right)
}

// Equal will create a ConditionBuilder. This will be the method for SizeBuilder
//
// Example:
//
//     The following produces equivalent conditions:
//     condition := expression.Equal(expression.Name("foo").Size(), expression.Value(5))
//     condition := expression.Name("foo").Size().Equal(expression.Value(5))
func (sb SizeBuilder) Equal(right OperandBuilder) ConditionBuilder {
	return Equal(sb, right)
}

// NotEqual will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Builder to be used in an operation input. This will be the function
// call.
//
// Example:
//
//     condition := expression.NotEqual(expression.Name("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
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
func (nb NameBuilder) NotEqual(right OperandBuilder) ConditionBuilder {
	return NotEqual(nb, right)
}

// NotEqual will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.NotEqual(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).NotEqual(expression.Value(5))
func (vb ValueBuilder) NotEqual(right OperandBuilder) ConditionBuilder {
	return NotEqual(vb, right)
}

// NotEqual will create a ConditionBuilder. This will be the method for
// SizeBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.NotEqual(expression.Name("foo").Size(), expression.Value(5))
//     condition := expression.Name("foo").Size().NotEqual(expression.Value(5))
func (sb SizeBuilder) NotEqual(right OperandBuilder) ConditionBuilder {
	return NotEqual(sb, right)
}

// LessThan will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Builder to be used in an operation input. This will be the function
// call.
//
// Example:
//
//     condition := expression.LessThan(expression.Name("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
func LessThan(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        lessThanCond,
	}
}

// LessThan will create a ConditionBuilder. This will be the method for
// NameBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.LessThan(expression.Name("foo"), expression.Value(5))
//     condition := expression.Name("foo").LessThan(expression.Value(5))
func (nb NameBuilder) LessThan(right OperandBuilder) ConditionBuilder {
	return LessThan(nb, right)
}

// LessThan will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.LessThan(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).LessThan(expression.Value(5))
func (vb ValueBuilder) LessThan(right OperandBuilder) ConditionBuilder {
	return LessThan(vb, right)
}

// LessThan will create a ConditionBuilder. This will be the method for
// SizeBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.LessThan(expression.Name("foo").Size(), expression.Value(5))
//     condition := expression.Name("foo").Size().LessThan(expression.Value(5))
func (sb SizeBuilder) LessThan(right OperandBuilder) ConditionBuilder {
	return LessThan(sb, right)
}

// LessThanEqual will create a ConditionBuilder with two OperandBuilders as
// children, representing the two operands that are being compared. The
// resulting ConditionBuilder can be used to build other Conditions or to create
// an Builder to be used in an operation input. This will be the
// function call.
//
// Example:
//
//     condition := expression.LessThanEqual(expression.Name("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
func LessThanEqual(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        lessThanEqualCond,
	}
}

// LessThanEqual will create a ConditionBuilder. This will be the method for
// NameBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.LessThanEqual(expression.Name("foo"), expression.Value(5))
//     condition := expression.Name("foo").LessThanEqual(expression.Value(5))
func (nb NameBuilder) LessThanEqual(right OperandBuilder) ConditionBuilder {
	return LessThanEqual(nb, right)
}

// LessThanEqual will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.LessThanEqual(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).LessThanEqual(expression.Value(5))
func (vb ValueBuilder) LessThanEqual(right OperandBuilder) ConditionBuilder {
	return LessThanEqual(vb, right)
}

// LessThanEqual will create a ConditionBuilder. This will be the method for
// SizeBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.LessThanEqual(expression.Name("foo").Size(), expression.Value(5))
//     condition := expression.Name("foo").Size().LessThanEqual(expression.Value(5))
func (sb SizeBuilder) LessThanEqual(right OperandBuilder) ConditionBuilder {
	return LessThanEqual(sb, right)
}

// GreaterThan will create a ConditionBuilder with two OperandBuilders as
// children, representing the two operands that are being compared. The
// resulting ConditionBuilder can be used to build other Conditions or to create
// an Builder to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.GreaterThan(expression.Name("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
func GreaterThan(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        greaterThanCond,
	}
}

// GreaterThan will create a ConditionBuilder. This will be the method for
// NameBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.GreaterThan(expression.Name("foo"), expression.Value(5))
//     condition := expression.Name("foo").GreaterThan(expression.Value(5))
func (nb NameBuilder) GreaterThan(right OperandBuilder) ConditionBuilder {
	return GreaterThan(nb, right)
}

// GreaterThan will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.GreaterThan(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).GreaterThan(expression.Value(5))
func (vb ValueBuilder) GreaterThan(right OperandBuilder) ConditionBuilder {
	return GreaterThan(vb, right)
}

// GreaterThan will create a ConditionBuilder. This will be the method for
// SizeBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.GreaterThan(expression.Name("foo").Size(), expression.Value(5))
//     condition := expression.Name("foo").Size().GreaterThan(expression.Value(5))
func (sb SizeBuilder) GreaterThan(right OperandBuilder) ConditionBuilder {
	return GreaterThan(sb, right)
}

// GreaterThanEqual will create a ConditionBuilder with two OperandBuilders as
// children, representing the two operands that are being compared. The
// resulting ConditionBuilder can be used to build other Conditions or to create
// an Builder to be used in an operation input. This will be the
// function call.
//
// Example:
//
//     condition := expression.GreaterThanEqual(expression.Name("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
func GreaterThanEqual(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        greaterThanEqualCond,
	}
}

// GreaterThanEqual will create a ConditionBuilder. This will be the method for
// NameBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.GreaterThanEqual(expression.Name("foo"), expression.Value(5))
//     condition := expression.Name("foo").GreaterThanEqual(expression.Value(5))
func (nb NameBuilder) GreaterThanEqual(right OperandBuilder) ConditionBuilder {
	return GreaterThanEqual(nb, right)
}

// GreaterThanEqual will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.GreaterThanEqual(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).GreaterThanEqual(expression.Value(5))
func (vb ValueBuilder) GreaterThanEqual(right OperandBuilder) ConditionBuilder {
	return GreaterThanEqual(vb, right)
}

// GreaterThanEqual will create a ConditionBuilder. This will be the method for
// SizeBuilder
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.GreaterThanEqual(expression.Name("foo").Size(), expression.Value(5))
//     condition := expression.Name("foo").Size().GreaterThanEqual(expression.Value(5))
func (sb SizeBuilder) GreaterThanEqual(right OperandBuilder) ConditionBuilder {
	return GreaterThanEqual(sb, right)
}

// And will create a ConditionBuilder with more than two other Conditions as
// children, representing logical statements that will be logically ANDed
// together. The resulting ConditionBuilder can be used to build other
// Conditions or to create an Builder to be used in an operation input. This
// will be the function call.
//
// Example:
//
//     condition1 := expression.Equal(expression.Name("foo"), expression.Value(5))
//     condition2 := expression.LessThan(expression.Name("bar"), expression.Value(2010))
//     condition3 := expression.Name("baz").Between(expression.Value(2), expression.Value(10))
//     andCondition := expression.And(condition1, condition2, condition3)
//
//     anotherCondition := expression.Not(andCondition)  // Used in another condition
//     builder := WithCondition(condition)               // Used to make an Builder
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
func (cb ConditionBuilder) And(right ConditionBuilder, other ...ConditionBuilder) ConditionBuilder {
	return And(cb, right, other...)
}

// Or will create a ConditionBuilder with more than two other Conditions as
// children, representing logical statements that will be logically ORed
// together. The resulting ConditionBuilder can be used to build other
// Conditions or to create an Builder to be used in an operation input. This
// will be the function call.
//
// Example:
//
//     condition1 := expression.Equal(expression.Name("foo"), expression.Value(5))
//     condition2 := expression.LessThan(expression.Name("bar"), expression.Value(2010))
//     condition3 := expression.Name("baz").Between(expression.Value(2), expression.Value(10))
//     orCondition := expression.Or(condition1, condition2, condition3)
//
//     anotherCondition := expression.Not(orCondition)  // Used in another condition
//     builder := WithCondition(condition)              // Used to make an Builder
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
func (cb ConditionBuilder) Or(right ConditionBuilder, other ...ConditionBuilder) ConditionBuilder {
	return Or(cb, right, other...)
}

// Not will create a ConditionBuilder with one Conditions as a child,
// representing the logical statements that will be logically negated. The
// resulting ConditionBuilder can be used to build other Conditions or to create
// an Builder to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Equal(expression.Name("foo"), expression.Value(5))
//     notCondition := expression.Or(condition)
//
//     anotherCondition := expression.Not(notCondition)  // Used in another condition
//     builder := WithCondition(condition)               // Used to make an Builder
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
func (cb ConditionBuilder) Not() ConditionBuilder {
	return Not(cb)
}

// Between will create a ConditionBuilder with three operands as children, the
// first operand representing the operand being compared, the second operand
// representing the lower bound value of the first operand, and the third
// operand representing the upper bound value of the first operand. The
// resulting ConditionBuilder can be used to build other Conditions or to create
// an Builder to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Between(expression.Name("foo"), expression.Value(2), expression.Value(6))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
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
func (nb NameBuilder) Between(lower, upper OperandBuilder) ConditionBuilder {
	return Between(nb, lower, upper)
}

// Between will create a ConditionBuilder. This will be the method signature for
// ValueBuilders.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Between(operand1, operand2, operand3)
//     condition := operand1.Between(operand2, operand3)
func (vb ValueBuilder) Between(lower, upper OperandBuilder) ConditionBuilder {
	return Between(vb, lower, upper)
}

// Between will create a ConditionBuilder. This will be the method signature for
// SizeBuilders.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.Between(operand1, operand2, operand3)
//     condition := operand1.Between(operand2, operand3)
func (sb SizeBuilder) Between(lower, upper OperandBuilder) ConditionBuilder {
	return Between(sb, lower, upper)
}

// In will create a ConditionBuilder with two or more operands as children, the
// first operand representing the operand being compared and the rest of the
// operands representing a set in which the first operand either belongs to or
// not. The argument must have at least two operands. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Builder to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Between(expression.Name("foo"), expression.Value(2), expression.Value(6))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
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
func (nb NameBuilder) In(right OperandBuilder, other ...OperandBuilder) ConditionBuilder {
	return In(nb, right, other...)
}

// In will create a ConditionBuilder. This will be the method signature for
// ValueBuilders.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.In(operand1, operand2, operand3)
//     condition := operand1.In(operand2, operand3)
func (vb ValueBuilder) In(right OperandBuilder, other ...OperandBuilder) ConditionBuilder {
	return In(vb, right, other...)
}

// In will create a ConditionBuilder. This will be the method signature for
// SizeBuilders.
//
// Example:
//
//     // The following produces equivalent conditions:
//     condition := expression.In(operand1, operand2, operand3)
//     condition := operand1.In(operand2, operand3)
func (sb SizeBuilder) In(right OperandBuilder, other ...OperandBuilder) ConditionBuilder {
	return In(sb, right, other...)
}

// AttributeExists will create a ConditionBuilder with a name as a child. The
// function will return true if the item attribute described by the name exists.
// The resulting ConditionBuilder can be used to build other Conditions or to
// create an Builder to be used in an operation input. This will be the function
// call.
//
// Example:
//
//     condition := expression.AttributeExists(Name("foo"))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
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
func (nb NameBuilder) AttributeExists() ConditionBuilder {
	return AttributeExists(nb)
}

// AttributeNotExists will create a ConditionBuilder with a name as a child. The
// function will return true if the item attribute described by the name does
// not exist. The resulting ConditionBuilder can be used to build other
// Conditions or to create an Builder to be used in an operation input. This
// will be the function call.
//
// Example:
//
//     condition := expression.AttributeNotExists(expression.Name("foo"))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
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
func (nb NameBuilder) AttributeNotExists() ConditionBuilder {
	return AttributeNotExists(nb)
}

// AttributeType will create a ConditionBuilder with a name and a value as a
// child. The name will represent the item attribute being compared. The value
// will be a string corresponding to the argument DynamoDBAttributeType. The
// function will return true if the item attribute described by the name is the
// type specified by DynamoDBAttributeType. The resulting ConditionBuilder can
// be used to build other Conditions or to create an Builder to be used in an
// operation input. This will be the function call.
//
// Example:
//
//     condition := expression.AttributeType(Name("foo"), expression.StringSet)
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
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
func (nb NameBuilder) AttributeType(attributeType DynamoDBAttributeType) ConditionBuilder {
	return AttributeType(nb, attributeType)
}

// BeginsWith will create a ConditionBuilder with a name and a value as
// children. The name will represent the name to the item attribute being
// compared. The value will represent the substring in which the item attribute
// will be compared with. The function will return true if the item attribute
// specified by the name starts with the substring. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Builder to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.BeginsWith(Name("foo"), "bar")
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
func BeginsWith(nameBuilder NameBuilder, prefix string) ConditionBuilder {
	v := ValueBuilder{
		value: prefix,
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
func (nb NameBuilder) BeginsWith(prefix string) ConditionBuilder {
	return BeginsWith(nb, prefix)
}

// Contains will create a ConditionBuilder with a name and a value as
// children. The name will represent the name to the item attribute being
// compared. The item attribute MUST be a String or a Set. The value will
// represent the string in which the item attribute will be compared with. The
// function will return true if the item attribute specified by the name
// contains the substring specified by the value or if the item attribute is a
// set that contains the string specified by the value. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Builder to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Contains(Name("foo"), "bar")
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     builder := WithCondition(condition)            // Used to make an Builder
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
func (nb NameBuilder) Contains(substr string) ConditionBuilder {
	return Contains(nb, substr)
}

// buildTree will build a tree structure of exprNodes based on the tree
// structure of the input ConditionBuilder's child ConditionBuilders and
// OperandBuilders. buildTree() satisfies the treeBuilder interface so
// ConditionBuilder can be a part of Builder and Expression struct.
func (cb ConditionBuilder) buildTree() (exprNode, error) {
	childNodes, err := cb.buildChildNodes()
	if err != nil {
		return exprNode{}, err
	}
	ret := exprNode{
		children: childNodes,
	}

	switch cb.mode {
	case equalCond, notEqualCond, lessThanCond, lessThanEqualCond, greaterThanCond, greaterThanEqualCond:
		return compareBuildCondition(cb.mode, ret)
	case andCond, orCond:
		return compoundBuildCondition(cb, ret)
	case notCond:
		return notBuildCondition(ret)
	case betweenCond:
		return betweenBuildCondition(ret)
	case inCond:
		return inBuildCondition(cb, ret)
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
		return exprNode{}, ErrUnsetCondition
	default:
		return exprNode{}, fmt.Errorf("build condition error: unsupported mode: %v", cb.mode)
	}
}

// compareBuildCondition is the function to make exprNodes from Compare
// ConditionBuilders. compareBuildCondition will only be called by the
// buildTree method. This function assumes that the argument ConditionBuilder
// has the right format.
func compareBuildCondition(conditionMode conditionMode, node exprNode) (exprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	switch conditionMode {
	case equalCond:
		node.fmtExpr = "$c = $c"
	case notEqualCond:
		node.fmtExpr = "$c <> $c"
	case lessThanCond:
		node.fmtExpr = "$c < $c"
	case lessThanEqualCond:
		node.fmtExpr = "$c <= $c"
	case greaterThanCond:
		node.fmtExpr = "$c > $c"
	case greaterThanEqualCond:
		node.fmtExpr = "$c >= $c"
	default:
		return exprNode{}, fmt.Errorf("build compare condition error: unsupported mode: %v", conditionMode)
	}

	return node, nil
}

// compoundBuildCondition is the function to make exprNodes from And/Or
// ConditionBuilders. compoundBuildCondition will only be called by the
// buildTree method. This function assumes that the argument ConditionBuilder
// has the right format.
func compoundBuildCondition(conditionBuilder ConditionBuilder, node exprNode) (exprNode, error) {
	// create a string with escaped characters to substitute them with proper
	// aliases during runtime
	var mode string
	switch conditionBuilder.mode {
	case andCond:
		mode = " AND "
	case orCond:
		mode = " OR "
	default:
		return exprNode{}, fmt.Errorf("build compound condition error: unsupported mode: %v", conditionBuilder.mode)
	}
	node.fmtExpr = "($c)" + strings.Repeat(mode+"($c)", len(conditionBuilder.conditionList)-1)

	return node, nil
}

// notBuildCondition is the function to make exprNodes from Not
// ConditionBuilders. notBuildCondition will only be called by the
// buildTree method. This function assumes that the argument ConditionBuilder
// has the right format.
func notBuildCondition(node exprNode) (exprNode, error) {
	// create a string with escaped characters to substitute them with proper
	// aliases during runtime
	node.fmtExpr = "NOT ($c)"

	return node, nil
}

// betweenBuildCondition is the function to make exprNodes from Between
// ConditionBuilders. BuildCondition will only be called by the
// buildTree method. This function assumes that the argument ConditionBuilder
// has the right format.
func betweenBuildCondition(node exprNode) (exprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	node.fmtExpr = "$c BETWEEN $c AND $c"

	return node, nil
}

// inBuildCondition is the function to make exprNodes from In
// ConditionBuilders. inBuildCondition will only be called by the
// buildTree method. This function assumes that the argument ConditionBuilder
// has the right format.
func inBuildCondition(conditionBuilder ConditionBuilder, node exprNode) (exprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	node.fmtExpr = "$c IN ($c" + strings.Repeat(", $c", len(conditionBuilder.operandList)-2) + ")"

	return node, nil
}

// attrExistsBuildCondition is the function to make exprNodes from
// AttrExistsCond ConditionBuilders. attrExistsBuildCondition will only be
// called by the buildTree method. This function assumes that the argument
// ConditionBuilder has the right format.
func attrExistsBuildCondition(node exprNode) (exprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	node.fmtExpr = "attribute_exists ($c)"

	return node, nil
}

// attrNotExistsBuildCondition is the function to make exprNodes from
// AttrNotExistsCond ConditionBuilders. attrNotExistsBuildCondition will only be
// called by the buildTree method. This function assumes that the argument
// ConditionBuilder has the right format.
func attrNotExistsBuildCondition(node exprNode) (exprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	node.fmtExpr = "attribute_not_exists ($c)"

	return node, nil
}

// attrTypeBuildCondition is the function to make exprNodes from AttrTypeCond
// ConditionBuilders. attrTypeBuildCondition will only be called by the
// buildTree method. This function assumes that the argument
// ConditionBuilder has the right format.
func attrTypeBuildCondition(node exprNode) (exprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	node.fmtExpr = "attribute_type ($c, $c)"

	return node, nil
}

// beginsWithBuildCondition is the function to make exprNodes from
// BeginsWithCond ConditionBuilders. beginsWithBuildCondition will only be
// called by the buildTree method. This function assumes that the argument
// ConditionBuilder has the right format.
func beginsWithBuildCondition(node exprNode) (exprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	node.fmtExpr = "begins_with ($c, $c)"

	return node, nil
}

// containsBuildCondition is the function to make exprNodes from
// ContainsCond ConditionBuilders. containsBuildCondition will only be
// called by the buildTree method. This function assumes that the argument
// ConditionBuilder has the right format.
func containsBuildCondition(node exprNode) (exprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	node.fmtExpr = "contains ($c, $c)"

	return node, nil
}

// buildChildNodes will create the list of the child exprNodes. This avoids
// duplication of code amongst the various buildTree functions.
func (cb ConditionBuilder) buildChildNodes() ([]exprNode, error) {
	childNodes := make([]exprNode, 0, len(cb.conditionList)+len(cb.operandList))
	for _, condition := range cb.conditionList {
		node, err := condition.buildTree()
		if err != nil {
			return []exprNode{}, err
		}
		childNodes = append(childNodes, node)
	}
	for _, ope := range cb.operandList {
		operand, err := ope.BuildOperand()
		if err != nil {
			return []exprNode{}, err
		}
		childNodes = append(childNodes, operand.exprNode)
	}

	return childNodes, nil
}
