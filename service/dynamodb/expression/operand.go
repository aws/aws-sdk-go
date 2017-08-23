package expression

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	// ErrEmptyName is an error that is returned if the name specified is or has
	// an empty string. (i.e. "", "foo..bar")
	ErrEmptyName = awserr.New("EmptyName", "BuildOperand error: name is an empty string", nil)

	// ErrInvalidNameIndex is an error that is returned if the index used in the
	// name either does not have an item attribute associated with it or is an
	// empty string.
	// (i.e. "[2]", "foo[]")
	ErrInvalidNameIndex = awserr.New("InvalidNameIndex", "BuildOperand error: invalid name index", nil)
)

// ValueBuilder represents a value operand and will implement the OperandBuilder
// interface. It will have various methods corresponding to the operations
// supported by DynamoDB operations. (i.e. AND, BETWEEN, EQUALS) The underlying
// undefined type member variable will be converted into dynamodb.AttributeValue
// by dynamodbattribute.Marshal(in) when Builder are created.
type ValueBuilder struct {
	value interface{}
}

// NameBuilder represents a name of a top level item attribute or a nested
// attribute. It will implement the OperandBuilder interface. It will
// have various methods corresponding to the operations supported by DynamoDB
// operations. (i.e. AND, BETWEEN, EQUALS)
type NameBuilder struct {
	name string
}

// SizeBuilder represents the output of the function size ("someName"), which
// evaluates to the size of the item attribute defined by "someName".
// SizeBuilder will implement OperandBuilder interface. It will have various
// methods corresponding to the operations supported by DynamoDB operations.
// (i.e. AND, BETWEEN, EQUALS)
type SizeBuilder struct {
	nameBuilder NameBuilder
}

// OperandBuilder represents the idea of Operand which are building blocks to
// DynamoDB Expressions. OperandBuilders will be children of ConditionBuilders
// to represent a tree like structure of Expression dependencies. The method
// BuildNode() will create an instance of ExprNode, which is an generic
// representation of both Operands and Conditions. BuildNode() will mainly be
// called recursively by the BuildTree() method call when Builder is built from
// ConditionBuilders
type OperandBuilder interface {
	BuildNode() (ExprNode, error)
}

// Name creates a NameBuilder, which implements the OperandBuilder interface.
// Name will mainly be called in a pattern in order to create
// ConditionBuilders.
//
// Example:
//
//     condition := Name("foo").Equal(Name("bar"))
func Name(name string) NameBuilder {
	return NameBuilder{
		name: name,
	}
}

// Value creates a ValueBuilder, which implements the OperandBuilder
// interface. Value will mainly be called in a pattern in order to create
// ConditionBuilders.
//
// Example:
//
//     condition := Name("foo").Equal(Value(10))
func Value(value interface{}) ValueBuilder {
	return ValueBuilder{
		value: value,
	}
}

// Size creates a SizeBuilder, which implements the OperandBuilder interface.
// Size will mainly be called in a pattern in order to create ConditionBuilders.
//
// Example:
//
//     condition := Name("foo").Size().Equal(Value(10))
func (nameBuilder NameBuilder) Size() SizeBuilder {
	return SizeBuilder{
		nameBuilder: nameBuilder,
	}
}

// BuildNode will create the ExprNode which is a generic representation of
// Operands and Conditions. BuildNode() is mainly for the BuildTree() method to
// call on, not for users to invoke. BuildOperand aliases all strings to avoid
// stepping over DynamoDB's reserved words.
// More information on reserved words at http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html
func (nameBuilder NameBuilder) BuildNode() (ExprNode, error) {
	if nameBuilder.name == "" {
		return ExprNode{}, ErrEmptyName
	}

	ret := ExprNode{
		names: []string{},
	}

	nameSplit := strings.Split(nameBuilder.name, ".")
	fmtNames := make([]string, 0, len(nameSplit))

	for _, word := range nameSplit {
		var substr string
		if word == "" {
			return ExprNode{}, ErrEmptyName
		}

		if word[len(word)-1] == ']' {
			for j, char := range word {
				if char == '[' {
					substr = word[j:]
					word = word[:j]
					break
				}
			}
		}

		if word == "" {
			return ExprNode{}, ErrInvalidNameIndex
		}

		// Create a string with special characters that can be substituted later: $p
		ret.names = append(ret.names, word)
		fmtNames = append(fmtNames, "$n"+substr)
	}
	ret.fmtExpr = strings.Join(fmtNames, ".")
	return ret, nil
}

// BuildNode will create the ExprNode which is a generic representation of
// Operands and Conditions. BuildNode() is mainly for the BuildTree() method to
// call on, not for users to invoke.
func (valueBuilder ValueBuilder) BuildNode() (ExprNode, error) {
	expr, err := dynamodbattribute.Marshal(valueBuilder.value)
	if err != nil {
		return ExprNode{}, err
	}

	// Create a string with special characters that can be substituted later: $v
	ret := ExprNode{
		values:  []dynamodb.AttributeValue{*expr},
		fmtExpr: "$v",
	}
	return ret, nil
}

// BuildNode will create the ExprNode which is a generic representation of
// Operands and Conditions. BuildNode() is mainly for the BuildTree() method to
// call on, not for users to invoke.
func (sizeBuilder SizeBuilder) BuildNode() (ExprNode, error) {
	ret, err := sizeBuilder.nameBuilder.BuildNode()
	ret.fmtExpr = "size (" + ret.fmtExpr + ")"

	return ret, err
}
