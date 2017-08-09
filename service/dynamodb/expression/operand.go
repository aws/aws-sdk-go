package expression

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	// ErrEmptyPath is an error that is returned if the path specified is or has
	// an empty string. (i.e. "", "foo..bar")
	ErrEmptyPath = awserr.New("EmptyPath", "BuildOperand error: path is an empty string", nil)

	// ErrInvalidPathIndex is an error that is returned if the index used in the
	// path either does not have a path associated with it or is an empty string.
	// (i.e. "[2]", "foo[]")
	ErrInvalidPathIndex = awserr.New("InvalidPathIndex", "BuildOperand error: invalid path index", nil)
)

// ValueBuilder represents a value operand and will implement the OperandBuilder
// interface. It will have various methods corresponding to the operations
// supported by DynamoDB operations. (i.e. AND, BETWEEN, EQUALS) The underlying
// undefined type member variable will be converted into dynamodb.AttributeValue
// by dynamodbattribute.Marshal(in) when Expressions are created.
type ValueBuilder struct {
	value interface{}
}

// PathBuilder represents a path to either a top level item attribute or a
// nested attribute. It will implement the OperandBuilder interface. It will
// have various methods corresponding to the operations supported by DynamoDB
// operations. (i.e. AND, BETWEEN, EQUALS)
type PathBuilder struct {
	path string
}

// SizeBuilder represents the output of the function size (path), which
// evaluates to the size of the item attribute defined by path. Size builder
// will implement OperandBuilder interface. It will have various methods
// corresponding to the operations supported by DynamoDB operations.
// (i.e. AND, BETWEEN, EQUALS)
type SizeBuilder struct {
	pb PathBuilder
}

// OperandBuilder represents the idea of Operand which are building blocks to
// DynamoDB Expressions. OperandBuilders will be children of ConditionBuilders
// to represent a tree like structure of Expression dependencies. The method
// BuildOperand() will create an instance of ExprNode, which is an generic
// representation of both Operands and Conditions. BuildOperand() will mainly
// be called recursively by the BuildExpression() method call when Expressions
// are built from ConditionBuilders
type OperandBuilder interface {
	BuildOperand() (ExprNode, error)
}

// Path creates a PathBuilder, which implements the OperandBuilder interface.
// Path will mainly be called in a pattern in order to create
// ConditionBuilders.
//
// Example:
//
//     condition := Path("foo").Equal(Path("bar"))
func Path(p string) PathBuilder {
	return PathBuilder{
		path: p,
	}
}

// Value creates a ValueBuilder, which implements the OperandBuilder
// interface. Value will mainly be called in a pattern in order to create
// ConditionBuilders.
//
// Example:
//
//     condition := Path("foo").Equal(Value(10))
func Value(v interface{}) ValueBuilder {
	return ValueBuilder{
		value: v,
	}
}

// Size creates a SizeBuilder, which implements the OperandBuilder interface.
// Size will mainly be called in a pattern in order to create ConditionBuilders.
//
// Example:
//
//     condition := Path("foo").Size().Equal(Value(10))
func (p PathBuilder) Size() SizeBuilder {
	return SizeBuilder{
		pb: p,
	}
}

// BuildOperand will create the ExprNode which is a generic representation of
// Operands and Conditions. BuildOperand() is mainly for the BuildExpression()
// method to call on, not for users to invoke. BuildOperand aliases all strings
// to avoid stepping over DynamoDB's reserved words.
// More information on reserved words at http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html
func (p PathBuilder) BuildOperand() (ExprNode, error) {
	if p.path == "" {
		return ExprNode{}, ErrEmptyPath
	}

	ret := ExprNode{
		names: []string{},
	}

	nameSplit := strings.Split(p.path, ".")
	fmtNames := make([]string, 0, len(nameSplit))

	for _, word := range nameSplit {
		var substr string
		if word == "" {
			return ExprNode{}, ErrEmptyPath
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
			return ExprNode{}, ErrInvalidPathIndex
		}

		// Create a string with special characters that can be substituted later: $p
		ret.names = append(ret.names, word)
		fmtNames = append(fmtNames, "$p"+substr)
	}
	ret.fmtExpr = strings.Join(fmtNames, ".")
	return ret, nil
}

// BuildOperand will create the ExprNode which is a generic representation of
// Operands and Conditions. BuildOperand() is mainly for the BuildExpression()
// method to call on, not for users to invoke.
func (v ValueBuilder) BuildOperand() (ExprNode, error) {
	expr, err := dynamodbattribute.Marshal(v.value)
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

// BuildOperand will create the ExprNode which is a generic representation of
// Operands and Conditions. BuildOperand() is mainly for the BuildExpression()
// method to call on, not for users to invoke.
func (s SizeBuilder) BuildOperand() (ExprNode, error) {
	ret, err := s.pb.BuildOperand()
	ret.fmtExpr = "size (" + ret.fmtExpr + ")"

	return ret, err
}
