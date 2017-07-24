package expression

import (
	"encoding/base32"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// ValueBuilder will be the concrete struct that satisfies the OperandBuilder
// interface. It will have various methods corresponding to the operations
// supported in ConditionExpressions
type ValueBuilder struct {
	value interface{}
}

// PathBuilder will be the concrete struct that satisfies the OperandBuilder
// interface. It will have various methods corresponding to the operations
// supported in ConditionExpressions
type PathBuilder struct {
	path string
}

// Expression implements the expressions in DynamoDB. DynamoDB operation inputs
// take maps of aliases to pointers and strings to represent expressions.
type Expression struct {
	Names      map[string]*string
	Values     map[string]*dynamodb.AttributeValue
	Expression string
}

// OperandBuilder will be mainly satisfied by PathBuilder and ValueBuilder.
// Concrete types that satisfy this interface will be referred to as an Operand
// In select cases, other builders may satisfy this interface
type OperandBuilder interface {
	BuildOperand() (Expression, error)
}

// NewPath creates an Operand based off of the path entered
func NewPath(p string) PathBuilder {
	return PathBuilder{
		path: p,
	}
}

// NewValue creates an Operand based of the value entered
func NewValue(v interface{}) ValueBuilder {
	return ValueBuilder{
		value: v,
	}
}

// BuildOperand will create an instance of an Expression. BuildOperand will be
// called whenever Build___() for any Expression Builders
func (v ValueBuilder) BuildOperand() (Expression, error) {
	attrval, err := dynamodbattribute.Marshal(v.value)
	expr := ":" + encode(fmt.Sprint(*attrval))

	return Expression{
		Values: map[string]*dynamodb.AttributeValue{
			expr: attrval,
		},
		Expression: expr,
	}, err
}

// BuildOperand will create an instance of an Expression. BuildOperand will be
// called whenever Build___() for any Expression Builders
func (p PathBuilder) BuildOperand() (Expression, error) {
	nameMap := make(map[string]*string)
	expr := ""

	nameSplit := strings.Split(p.path, ".")
	for i, word := range nameSplit {
		var substr string
		if string(word[len(word)-1]) == "]" {
			for j, char := range word {
				if string(char) == "[" {
					substr = word[j:]
					word = word[:j]
				}
			}
		}

		token := "#" + encode(word)
		nameMap[token] = aws.String(word)
		expr += token + substr
		if i != len(nameSplit)-1 {
			expr += "."
		}
	}
	return Expression{
		Names:      nameMap,
		Expression: expr,
	}, nil
}

// encodeName consistently encodes a name.
// The consistency is important.
// Taken from github.com/guregu/dynamo
func encode(name string) string {
	name = base32.StdEncoding.EncodeToString([]byte(name))
	return strings.TrimRight(name, "=")
}
