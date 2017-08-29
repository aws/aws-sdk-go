package expression

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
// attribute. Since NameBuilder represents a DynamoDB Operand, it will implement
// the OperandBuilder interface. It will have various methods corresponding to
// the operations supported by DynamoDB operations. (i.e. AND, BETWEEN, EQUALS)
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

// Operand is a wrapper around the exprNode struct. BuildOperand returns Operand
// instead of exprNode so that the underlying structure of exprNodes is not
// accessible to external users.
type Operand struct {
	exprNode exprNode
}

// OperandBuilder represents the idea of Operand which are building blocks to
// DynamoDB Expressions. OperandBuilders will be children of ConditionBuilders
// to represent a tree like structure of Expression dependencies. The method
// BuildOperand() will create an instance of Operand with a child exprNode,
// which is an generic representation of both Operands and Conditions.
// BuildOperand() will mainly be called recursively by the buildTree() method
// call.
type OperandBuilder interface {
	BuildOperand() (Operand, error)
}

// Name creates a NameBuilder, which implements the OperandBuilder interface.
//
// Example:
//
//     condition := Name("foo").Equal(Name("bar"))
func Name(name string) NameBuilder {
	return NameBuilder{
		name: name,
	}
}

// Value creates a ValueBuilder, which implements the OperandBuilder interface.
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
func (nb NameBuilder) Size() SizeBuilder {
	return SizeBuilder{
		nameBuilder: nb,
	}
}

// BuildOperand will create an Operand struct with an exprNode as a member,
// which is a generic representation of Operands and Conditions. BuildOperand()
// is mainly for the buildTree() method to call on, not for users to invoke.
// BuildOperand aliases all strings to avoid stepping over DynamoDB's reserved
// words.
// More information on reserved words at http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html
func (nb NameBuilder) BuildOperand() (Operand, error) {
	if nb.name == "" {
		return Operand{}, newUnsetParameterError("BuildOperand", "NameBuilder")
	}

	node := exprNode{
		names: []string{},
	}

	nameSplit := strings.Split(nb.name, ".")
	fmtNames := make([]string, 0, len(nameSplit))

	for _, word := range nameSplit {
		var substr string
		if word == "" {
			return Operand{}, newInvalidParameterError("BuildOperand", "NameBuilder")
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
			return Operand{}, newInvalidParameterError("BuildOperand", "NameBuilder")
		}

		// Create a string with special characters that can be substituted later: $p
		node.names = append(node.names, word)
		fmtNames = append(fmtNames, "$n"+substr)
	}
	node.fmtExpr = strings.Join(fmtNames, ".")
	return Operand{
		exprNode: node,
	}, nil
}

// BuildOperand will create an Operand struct with an exprNode as a member,
// which is a generic representation of Operands and Conditions. BuildOperand()
// is mainly for the BuildTree() method to call on, not for users to invoke.
func (vb ValueBuilder) BuildOperand() (Operand, error) {
	expr, err := dynamodbattribute.Marshal(vb.value)
	if err != nil {
		return Operand{}, newInvalidParameterError("BuildOperand", "ValueBuilder")
	}

	// Create a string with special characters that can be substituted later: $v
	operand := Operand{
		exprNode: exprNode{
			values:  []dynamodb.AttributeValue{*expr},
			fmtExpr: "$v",
		},
	}
	return operand, nil
}

// BuildOperand will create an Operand struct with an exprNode as a member,
// which is a generic representation of Operands and Conditions. BuildOperand()
// is mainly for the BuildTree() method to call on, not for users to invoke.
func (sb SizeBuilder) BuildOperand() (Operand, error) {
	operand, err := sb.nameBuilder.BuildOperand()
	operand.exprNode.fmtExpr = "size (" + operand.exprNode.fmtExpr + ")"

	return operand, err
}
