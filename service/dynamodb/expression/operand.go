package expression

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// OperandMode will specify the type of Operand the OperandExpression represents
type OperandMode int

const (
	// PathOpe will define an OperandExpression as a Path
	PathOpe OperandMode = iota + 1
	// ValueOpe will define an OperandExpression as a Value
	ValueOpe
	// SizeOpe will define an OperandExpression as a Size
	SizeOpe
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

// SizeBuilder will implement OperandBuilder thus being an Operand. This
// reflects the fact that the function Size() returns type that is used in place
// of an Operand
type SizeBuilder struct {
	pb PathBuilder
}

// Expression implements the expressions in DynamoDB. DynamoDB operation inputs
// take maps of aliases to pointers and strings to represent expressions.
type Expression struct {
	Names      map[string]*string
	Values     map[string]*dynamodb.AttributeValue
	Expression string
}

// OperandExpression will be the nodes to the inward facing tree which all the
// deduplication and aliasing will work on
type OperandExpression struct {
	Mode  OperandMode
	path  string
	value dynamodb.AttributeValue
}

// OperandBuilder will be mainly satisfied by PathBuilder and ValueBuilder.
// Concrete types that satisfy this interface will be referred to as an Operand
// In select cases, other builders may satisfy this interface
type OperandBuilder interface {
	BuildOperand() (OperandExpression, error)
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

// Size returns a SizeBuilder which satisfies the OperandBuilder interface.
func (p PathBuilder) Size() SizeBuilder {
	return SizeBuilder{
		pb: p,
	}
}

// BuildOperand will create the OperandExpression which will be recursively
// called in the BuildExpression operation
func (p PathBuilder) BuildOperand() (OperandExpression, error) {
	return OperandExpression{
		Mode: PathOpe,
		path: p.path,
	}, nil
}

// BuildOperand will create the OperandExpression which will be recursively
// called in the BuildExpression operation
func (v ValueBuilder) BuildOperand() (OperandExpression, error) {
	expr, err := dynamodbattribute.Marshal(v.value)

	return OperandExpression{
		Mode:  ValueOpe,
		value: *expr,
	}, err
}

// BuildOperand will create the OperandExpression which will be recursively
// called in the BuildExpression operation
func (s SizeBuilder) BuildOperand() (OperandExpression, error) {
	return OperandExpression{
		Mode: SizeOpe,
		path: s.pb.path,
	}, nil
}

// buildList lists all the names that must be aliased and returns the values
// in aliasMap.
func (oe OperandExpression) buildList() (aliasList, error) {
	switch oe.Mode {
	case PathOpe, SizeOpe:
		return pathBuildList(oe)
	case ValueOpe:
		return valueBuildList(oe)
	default:
		return aliasList{}, fmt.Errorf("OperandExpression buildList Error: Undefined OperandMode %#v", oe.Mode)
	}
}

// buildExpression returns an Expression with aliasing for paths/values specified
// by aliasList
func (oe OperandExpression) buildExpression(al aliasList) (Expression, error) {
	switch oe.Mode {
	case PathOpe:
		return pathBuildExpression(oe, al)
	case ValueOpe:
		return valueBuildExpression(oe, al)
	case SizeOpe:
		return sizeBuildExpression(oe, al)
	default:
		return Expression{}, fmt.Errorf("OperandExpression buildExpression Error: Undefined OperandMode %#v", oe.Mode)
	}
}

// Since we are not deduplicating values yet, the value will not be added to the
// aliasList.ValuesList and the alias for values will be done by looking at the
// length of aliasList.ValuesList during buildExpression
func valueBuildList(oe OperandExpression) (aliasList, error) {
	if reflect.DeepEqual(oe.value, (dynamodb.AttributeValue{})) {
		return aliasList{}, fmt.Errorf("valueBuildList Error: OperandExpression value is empty")
	}
	// return aliasList{
	// 	ValuesList: []dynamodb.AttributeValue{oe.value},
	// }, nil
	return aliasList{}, nil // deduplicating values will be implemented later
}

// in order to have unique aliases, we will use length of the
// aliasList.ValuesList to alias and add the value to the aliasList.ValuesList
func valueBuildExpression(oe OperandExpression, al aliasList) (Expression, error) {
	if reflect.DeepEqual(oe.value, (dynamodb.AttributeValue{})) {
		return Expression{}, fmt.Errorf("valueBuildExpression Error: OperandExpression value is empty")
	}

	alias := fmt.Sprintf(":%v", len(al.ValuesList))
	al.ValuesList = append(al.ValuesList, oe.value)

	return Expression{
		Values: map[string]*dynamodb.AttributeValue{
			alias: &oe.value,
		},
		Expression: alias,
	}, nil
}

// for path, we want to make sure to parse "." and "[]" so that we only alias
// item attributes, not path identifiers and list indexes
func pathBuildList(oe OperandExpression) (aliasList, error) {
	if oe.path == "" {
		return aliasList{}, fmt.Errorf("pathBuildList Error: Path is empty")
	}
	al := aliasList{
		NamesList: make([]string, 0),
	}

	nameSplit := strings.Split(oe.path, ".")
	for _, word := range nameSplit {
		if word == "" {
			return aliasList{}, fmt.Errorf("pathBuildList Error: Path is incomplete")
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

// we want to deduplicate names, so BuildExpression for path will search
// aliasList.NamesList for words that need to be aliased. If the alias is not
// found in the aliasList.NamesList, it will return an error.
func pathBuildExpression(oe OperandExpression, al aliasList) (Expression, error) {
	if oe.path == "" {
		return Expression{}, fmt.Errorf("pathBuildExpression Error: Path is empty")
	}

	ret := Expression{
		Names: make(map[string]*string),
	}

	nameSplit := strings.Split(oe.path, ".")
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
			return Expression{}, fmt.Errorf("pathBuildExpression could not find an alias for path")
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

// for size, the only difference from the pathBuildExpression should be the
// Expression string. We will take advantage of the existing pathBuildExpression
func sizeBuildExpression(oe OperandExpression, al aliasList) (Expression, error) {
	expr, err := pathBuildExpression(oe, al)
	if err != nil {
		return Expression{}, err
	}
	expr.Expression = "size (" + expr.Expression + ")"
	return expr, nil
}

// aliasList will keep track of all the names we need to alias in the nested
// struct of conditions and operands. This will allow each alias to be unique
// while deduplicating aliases.
type aliasList struct {
	NamesList  []string
	ValuesList []dynamodb.AttributeValue
}
