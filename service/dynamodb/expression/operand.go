package expression

import (
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
	BuildOperand(AliasList) (Expression, error)
	ListOperand() (AliasList, error)
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

// ListOperand lists all the names that must be aliased and returns the values
// in AliasMap. Since  we are not deduplicating Values, we don't need to make
// a list of the values, we will have just a counter
func (v ValueBuilder) ListOperand() (AliasList, error) {
	return AliasList{}, nil
}

// BuildOperand will create an instance of an Expression. BuildOperand will be
// called whenever Build___() for any Expression Builders
func (v ValueBuilder) BuildOperand(al AliasList) (Expression, error) {
	attrval, err := dynamodbattribute.Marshal(v.value)
	if err != nil {
		return Expression{}, err
	}

	if al.ValuesCounter == nil {
		return Expression{}, fmt.Errorf("Value Counter is nil")
	}

	alias := fmt.Sprintf(":%v", *al.ValuesCounter)
	*al.ValuesCounter++

	return Expression{
		Values: map[string]*dynamodb.AttributeValue{
			alias: attrval,
		},
		Expression: alias,
	}, nil
}

// ListOperand returns a list of names that must be aliased in the struct
// AliasMap.
func (p PathBuilder) ListOperand() (AliasList, error) {
	if p.path == "" {
		return AliasList{}, fmt.Errorf("ListOperand received an unexpected argument, Path is empty")
	}
	al := AliasList{
		NamesList: make([]string, 0),
	}

	nameSplit := strings.Split(p.path, ".")
	for _, word := range nameSplit {
		if word == "" {
			return AliasList{}, fmt.Errorf("ListOperand received an unexpected argument, Path is incomplete")
		}
		if string(word[len(word)-1]) == "]" {
			for j, char := range word {
				if string(char) == "[" {
					word = word[:j]
				}
			}
		}
		al.NamesList = append(al.NamesList, word)
	}
	return al, nil
}

// BuildOperand will create an instance of an Expression. BuildOperand will be
// called whenever Build___() for any Expression Builders
func (p PathBuilder) BuildOperand(al AliasList) (Expression, error) {
	ret := Expression{
		Names: make(map[string]*string),
	}

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

		index := -1
		for ind, val := range al.NamesList {
			if word == val {
				index = ind
				break
			}
		}

		if index == -1 {
			return Expression{}, fmt.Errorf("BuildOperand could not find an alias for path")
		}

		alias := fmt.Sprintf("#%v", index)
		ret.Names[alias] = aws.String(word)
		ret.Expression += alias + substr
		if i != len(nameSplit)-1 {
			ret.Expression += "."
		}
	}
	return ret, nil
}

// AliasList will keep track of all the names we need to alias in the nested
// struct of conditions and operands. This will allow each alias to be unique
// while deduplicating aliases.
type AliasList struct {
	NamesList []string
	//ValuesList []dynamodb.AttributeValue
	ValuesCounter *int
}
