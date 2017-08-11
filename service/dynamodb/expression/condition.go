package expression

import (
	"fmt"
	"strings"
)

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

// ConditionBuilder will represent the ConditionExpressions in DynamoDB. It is
// composed of operands (OperandBuilder) and other conditions (ConditionBuilder)
// There are many different types of conditions, specified by ConditionMode.
// Users will be able to call the BuildExpression() method on a ConditionBuilder
// to create an Expression which can then be used for operation inputs into
// DynamoDB. Only the Mode of the ConditionBuilder will be exported for users to
// check
// More Information at: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ConditionExpressions.html
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
//     condition := expression.Equal(expression.Path("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func Equal(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        equalCond,
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

// NotEqual will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.NotEqual(expression.Path("foo"), expression.Value(5))
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
// PathBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.NotEqual(expression.Path("foo"), expression.Value(5))
//     condition := expression.Path("foo").NotEqual(expression.Value(5))
func (p PathBuilder) NotEqual(right OperandBuilder) ConditionBuilder {
	return NotEqual(p, right)
}

// NotEqual will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.NotEqual(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).NotEqual(expression.Value(5))
func (v ValueBuilder) NotEqual(right OperandBuilder) ConditionBuilder {
	return NotEqual(v, right)
}

// NotEqual will create a ConditionBuilder. This will be the method for
// SizeBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.NotEqual(expression.Path("foo").Size(), expression.Value(5))
//     condition := expression.Path("foo").Size().NotEqual(expression.Value(5))
func (s SizeBuilder) NotEqual(right OperandBuilder) ConditionBuilder {
	return NotEqual(s, right)
}

// Less will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Less(expression.Path("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func Less(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        lessCond,
	}
}

// Less will create a ConditionBuilder. This will be the method for PathBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Less(expression.Path("foo"), expression.Value(5))
//     condition := expression.Path("foo").Less(expression.Value(5))
func (p PathBuilder) Less(right OperandBuilder) ConditionBuilder {
	return Less(p, right)
}

// Less will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Less(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).Less(expression.Value(5))
func (v ValueBuilder) Less(right OperandBuilder) ConditionBuilder {
	return Less(v, right)
}

// Less will create a ConditionBuilder. This will be the method for SizeBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Less(expression.Path("foo").Size(), expression.Value(5))
//     condition := expression.Path("foo").Size().Less(expression.Value(5))
func (s SizeBuilder) Less(right OperandBuilder) ConditionBuilder {
	return Less(s, right)
}

// LessEqual will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.LessEqual(expression.Path("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func LessEqual(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        lessEqualCond,
	}
}

// LessEqual will create a ConditionBuilder. This will be the method for PathBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.LessEqual(expression.Path("foo"), expression.Value(5))
//     condition := expression.Path("foo").LessEqual(expression.Value(5))
func (p PathBuilder) LessEqual(right OperandBuilder) ConditionBuilder {
	return LessEqual(p, right)
}

// LessEqual will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.LessEqual(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).LessEqual(expression.Value(5))
func (v ValueBuilder) LessEqual(right OperandBuilder) ConditionBuilder {
	return LessEqual(v, right)
}

// LessEqual will create a ConditionBuilder. This will be the method for SizeBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.LessEqual(expression.Path("foo").Size(), expression.Value(5))
//     condition := expression.Path("foo").Size().LessEqual(expression.Value(5))
func (s SizeBuilder) LessEqual(right OperandBuilder) ConditionBuilder {
	return LessEqual(s, right)
}

// Greater will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Greater(expression.Path("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func Greater(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        greaterCond,
	}
}

// Greater will create a ConditionBuilder. This will be the method for PathBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Greater(expression.Path("foo"), expression.Value(5))
//     condition := expression.Path("foo").Greater(expression.Value(5))
func (p PathBuilder) Greater(right OperandBuilder) ConditionBuilder {
	return Greater(p, right)
}

// Greater will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Greater(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).Greater(expression.Value(5))
func (v ValueBuilder) Greater(right OperandBuilder) ConditionBuilder {
	return Greater(v, right)
}

// Greater will create a ConditionBuilder. This will be the method for SizeBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Greater(expression.Path("foo").Size(), expression.Value(5))
//     condition := expression.Path("foo").Size().Greater(expression.Value(5))
func (s SizeBuilder) Greater(right OperandBuilder) ConditionBuilder {
	return Greater(s, right)
}

// GreaterEqual will create a ConditionBuilder with two OperandBuilders as children,
// representing the two operands that are being compared. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.GreaterEqual(expression.Path("foo"), expression.Value(5))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func GreaterEqual(left, right OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{left, right},
		mode:        greaterEqualCond,
	}
}

// GreaterEqual will create a ConditionBuilder. This will be the method for PathBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.GreaterEqual(expression.Path("foo"), expression.Value(5))
//     condition := expression.Path("foo").GreaterEqual(expression.Value(5))
func (p PathBuilder) GreaterEqual(right OperandBuilder) ConditionBuilder {
	return GreaterEqual(p, right)
}

// GreaterEqual will create a ConditionBuilder. This will be the method for
// ValueBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.GreaterEqual(expression.Value(10), expression.Value(5))
//     condition := expression.Value(10).GreaterEqual(expression.Value(5))
func (v ValueBuilder) GreaterEqual(right OperandBuilder) ConditionBuilder {
	return GreaterEqual(v, right)
}

// GreaterEqual will create a ConditionBuilder. This will be the method for SizeBuilder
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.GreaterEqual(expression.Path("foo").Size(), expression.Value(5))
//     condition := expression.Path("foo").Size().GreaterEqual(expression.Value(5))
func (s SizeBuilder) GreaterEqual(right OperandBuilder) ConditionBuilder {
	return GreaterEqual(s, right)
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
//     // The following produce equivalent conditions:
//     condition := expression.And(condition1, condition2, condition3)
//     condition := condition1.And(condition2, condition3)
func (cond ConditionBuilder) And(right ConditionBuilder, other ...ConditionBuilder) ConditionBuilder {
	return And(cond, right, other...)
}

// Or will create a ConditionBuilder with more than two other Conditions as
// children, representing logical statements that will be logically ORed
// together. The resulting ConditionBuilder can be used to build other
// Conditions or to create an Expression to be used in an operation input. This
// will be the function call.
//
// Example:
//
//     condition1 := expression.Equal(expression.Path("foo"), expression.Value(5))
//     condition2 := expression.Less(expression.Path("bar"), expression.Value(2010))
//     condition3 := expression.Path("baz").Between(expression.Value(2), expression.Value(10))
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
//     // The following produce equivalent conditions:
//     condition := expression.Or(condition1, condition2, condition3)
//     condition := condition1.Or(condition2, condition3)
func (cond ConditionBuilder) Or(right ConditionBuilder, other ...ConditionBuilder) ConditionBuilder {
	return Or(cond, right, other...)
}

// Not will create a ConditionBuilder with one Conditions as a child,
// representing the logical statements that will be logically negated. The
// resulting ConditionBuilder can be used to build other Conditions or to create
// an Expression to be used in an operation input. This will be the function
// call.
//
// Example:
//
//     condition := expression.Equal(expression.Path("foo"), expression.Value(5))
//     notCondition := expression.Or(condition)
//
//     anotherCondition := expression.Not(notCondition)  // Used in another condition
//     expression, err := notCondition.BuildExpression() // Used to make an Expression
func Not(cond ConditionBuilder) ConditionBuilder {
	return ConditionBuilder{
		conditionList: []ConditionBuilder{cond},
		mode:          notCond,
	}
}

// Not will create a ConditionBuilder. This will be the method signature
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Not(condition)
//     condition := condition.Not()
func (cond ConditionBuilder) Not() ConditionBuilder {
	return Not(cond)
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
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func Between(ope, lower, upper OperandBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{ope, lower, upper},
		mode:        betweenCond,
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

// In will create a ConditionBuilder with two or more operands as children, the
// first operand representing the operand being compared and the rest of the
// operands representing a set in which the first operand either belongs to or
// not. The argument must have at least two operands. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Between(expression.Path("foo"), expression.Value(2), expression.Value(6))
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
// PathBuilders.
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.In(operand1, operand2, operand3)
//     condition := operand1.In(operand2, operand3)
func (p PathBuilder) In(right OperandBuilder, other ...OperandBuilder) ConditionBuilder {
	return In(p, right, other...)
}

// In will create a ConditionBuilder. This will be the method signature for
// ValueBuilders.
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.In(operand1, operand2, operand3)
//     condition := operand1.In(operand2, operand3)
func (v ValueBuilder) In(right OperandBuilder, other ...OperandBuilder) ConditionBuilder {
	return In(v, right, other...)
}

// In will create a ConditionBuilder. This will be the method signature for
// SizeBuilders.
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.In(operand1, operand2, operand3)
//     condition := operand1.In(operand2, operand3)
func (s SizeBuilder) In(right OperandBuilder, other ...OperandBuilder) ConditionBuilder {
	return In(s, right, other...)
}

// AttributeExists will create a ConditionBuilder with a path as a child. The
// function will return true if the item attribute described by the path exists.
// The resulting ConditionBuilder can be used to build other Conditions or to
// create an Expression to be used in an operation input. This will be the
// function call.
//
// Example:
//
//     condition := expression.AttributeExists(Path("foo"))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func AttributeExists(p PathBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{p},
		mode:        attrExistsCond,
	}
}

// AttributeExists will create a ConditionBuilder. AttributeExists will only
// have a method for PathBuilders since that is the only valid operand that the
// function can be called on.
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.AttributeExists(Path("foo"))
//     condition := Path("foo").AttributeExists()
func (p PathBuilder) AttributeExists() ConditionBuilder {
	return AttributeExists(p)
}

// AttributeNotExists will create a ConditionBuilder with a path as a child. The
// function will return true if the item attribute described by the path does
// not exist. The resulting ConditionBuilder can be used to build other
// Conditions or to create an Expression to be used in an operation input. This
// will be the function call.
//
// Example:
//
//     condition := expression.AttributeNotExists(expression.Path("foo"))
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func AttributeNotExists(p PathBuilder) ConditionBuilder {
	return ConditionBuilder{
		operandList: []OperandBuilder{p},
		mode:        attrNotExistsCond,
	}
}

// AttributeNotExists will create a ConditionBuilder. AttributeNotExists will
// only have a method for PathBuilders since that is the only valid operand that
// the function can be called on.
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.AttributeNotExists(expression.Path("foo"))
//     condition := expression.Path("foo").AttributeNotExists()
func (p PathBuilder) AttributeNotExists() ConditionBuilder {
	return AttributeNotExists(p)
}

// AttributeType will create a ConditionBuilder with a path and a value as a
// child. The path will represent the item attribute being compared. The value
// will be a string corresponding to the argument DynamoDBAttributeType. The
// function will return true if the item attribute described by the path is the
// type specified by DynamoDBAttributeType. The resulting ConditionBuilder can
// be used to build other Conditions or to create an Expression to be used in an
// operation input. This will be the function call.
//
// Example:
//
//     condition := expression.AttributeType(Path("foo"), expression.StringSet)
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func AttributeType(p PathBuilder, at DynamoDBAttributeType) ConditionBuilder {
	v := ValueBuilder{
		value: string(at),
	}
	return ConditionBuilder{
		operandList: []OperandBuilder{p, v},
		mode:        attrTypeCond,
	}
}

// AttributeType will create a ConditionBuilder. AttributeType will only have a
// method for PathBuilders since that is the only valid operand that the
// function can be called on.
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.AttributeType(expression.Path("foo"), expression.Number)
//     condition := expression.Path("foo").AttributeType(expression.Number)
func (p PathBuilder) AttributeType(at DynamoDBAttributeType) ConditionBuilder {
	return AttributeType(p, at)
}

// BeginsWith will create a ConditionBuilder with a path and a value as
// children. The path will represent the path to the item attribute being
// compared. The value will represent the substring in which the item attribute
// will be compared with. The function will return true if the item attribute
// specified by the path starts with the substring. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.BeginsWith(Path("foo"), "bar")
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func BeginsWith(p PathBuilder, s string) ConditionBuilder {
	v := ValueBuilder{
		value: s,
	}
	return ConditionBuilder{
		operandList: []OperandBuilder{p, v},
		mode:        beginsWithCond,
	}
}

// BeginsWith will create a ConditionBuilder. BeginsWith will only have a method
// for PathBuilders since that is the only valid operand that the function can
// be called on.
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.BeginsWith(expression.Path("foo"), "bar")
//     condition := expression.Path("foo").BeginsWith("bar")
func (p PathBuilder) BeginsWith(s string) ConditionBuilder {
	return BeginsWith(p, s)
}

// Contains will create a ConditionBuilder with a path and a value as
// children. The path will represent the path to the item attribute being
// compared. The item attribute MUST be a String or a Set. The value will
// represent the string in which the item attribute will be compared with. The
// function will return true if the item attribute specified by the path
// contains the substring specified by the value or if the item attribute is a
// set that contains the string specified by the value. The resulting
// ConditionBuilder can be used to build other Conditions or to create an
// Expression to be used in an operation input. This will be the function call.
//
// Example:
//
//     condition := expression.Contains(Path("foo"), "bar")
//
//     anotherCondition := expression.Not(condition)  // Used in another condition
//     expression, err := condition.BuildExpression() // Used to make an Expression
func Contains(p PathBuilder, s string) ConditionBuilder {
	v := ValueBuilder{
		value: s,
	}
	return ConditionBuilder{
		operandList: []OperandBuilder{p, v},
		mode:        containsCond,
	}
}

// Contains will create a ConditionBuilder. Contains will only have a method
// for PathBuilders since that is the only valid operand that the function can
// be called on.
//
// Example:
//
//     // The following produce equivalent conditions:
//     condition := expression.Contains(expression.Path("foo"), "bar")
//     condition := expression.Path("foo").Contains("bar")
func (p PathBuilder) Contains(s string) ConditionBuilder {
	return Contains(p, s)
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
	childNodes, err := cond.buildChildNodes()
	if err != nil {
		return ExprNode{}, err
	}
	ret := ExprNode{
		children: childNodes,
	}

	switch cond.mode {
	case equalCond, notEqualCond, lessCond, lessEqualCond, greaterCond, greaterEqualCond:
		return compareBuildCondition(cond.mode, ret)
	case andCond, orCond:
		return compoundBuildCondition(cond, ret)
	case notCond:
		return notBuildCondition(ret)
	case betweenCond:
		return betweenBuildCondition(ret)
	case inCond:
		return inBuildCondition(cond, ret)
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
	}
	return ExprNode{}, fmt.Errorf("buildCondition error: unsupported mode: %v", cond.mode)
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
