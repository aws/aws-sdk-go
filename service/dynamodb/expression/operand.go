package expression

import (
	"fmt"
	"strings"

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

// ExprNode will be the nodes to the inward facing tree which all the
// deduplication and aliasing will work on
type ExprNode struct {
	names    []string
	values   []dynamodb.AttributeValue
	children []ExprNode
	fmtExpr  string
}

// OperandBuilder will be mainly satisfied by PathBuilder and ValueBuilder.
// Concrete types that satisfy this interface will be referred to as an Operand
// In select cases, other builders may satisfy this interface
type OperandBuilder interface {
	BuildOperand() (ExprNode, error)
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

// BuildOperand will create the ExprNode which will be recursively
// called in the BuildExpression operation
func (p PathBuilder) BuildOperand() (ExprNode, error) {
	if p.path == "" {
		return ExprNode{}, fmt.Errorf("BuildOperand Error: Path is empty")
	}

	ret := ExprNode{
		names: []string{},
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

		ret.names = append(ret.names, word)
		ret.fmtExpr += "$p" + substr
		if i != len(nameSplit)-1 {
			ret.fmtExpr += "."
		}
	}
	return ret, nil
}

// BuildOperand will create the ExprNode which will be recursively
// called in the BuildExpression operation
func (v ValueBuilder) BuildOperand() (ExprNode, error) {
	expr, err := dynamodbattribute.Marshal(v.value)
	if err != nil {
		return ExprNode{}, err
	}

	ret := ExprNode{
		values:  []dynamodb.AttributeValue{*expr},
		fmtExpr: "$v",
	}
	return ret, nil
}

// BuildOperand will create the ExprNode which will be recursively
// called in the BuildExpression operation
func (s SizeBuilder) BuildOperand() (ExprNode, error) {
	ret, err := s.pb.BuildOperand()
	ret.fmtExpr = "size (" + ret.fmtExpr + ")"

	return ret, err
}

// aliasList will keep track of all the names we need to alias in the nested
// struct of conditions and operands. This will allow each alias to be unique
// while deduplicating aliases.
type aliasList struct {
	namesList  []string
	valuesList []dynamodb.AttributeValue
}

// buildExpression returns an Expression with aliasing for paths/values specified
// by aliasList
func (en ExprNode) buildExprNodes(al *aliasList) (Expression, error) {
	if al == nil {
		return Expression{}, fmt.Errorf("buildExpression Error: aliasList is nil")
	}

	index := struct {
		name, value, children int
	}{}

	//index := [3]int{}
	expr := Expression{
		Expression: en.fmtExpr,
	}

	for i := 0; i < len(expr.Expression); {
		if expr.Expression[i] != '$' {
			i++
			continue
		}

		var alias string
		switch expr.Expression[i+1] {
		case 'p':
			str, err := al.aliasPath(en.names[index.name])
			if err != nil {
				return Expression{}, err
			}
			alias = str
			if expr.Names == nil {
				expr.Names = make(map[string]*string)
			}
			expr.Names[alias] = &en.names[index.name]
			index.name++

		case 'v':
			str, err := al.aliasValue(en.values[index.value])
			if err != nil {
				return Expression{}, err
			}
			alias = str
			if expr.Values == nil {
				expr.Values = make(map[string]*dynamodb.AttributeValue)
			}
			expr.Values[alias] = &en.values[index.value]
			index.value++

		case 'c':
			childExpr, err := en.children[index.children].buildExprNodes(al)
			if err != nil {
				return Expression{}, err
			}
			alias = childExpr.Expression
			tempExpr := expr.Expression
			expr = mergeExpressionMaps([]Expression{expr, childExpr})
			expr.Expression = tempExpr
			index.children++
		}
		expr.Expression = expr.Expression[:i] + alias + expr.Expression[i+2:]
		i += len(alias)
	}

	return expr, nil
}

func (al *aliasList) aliasValue(dav dynamodb.AttributeValue) (string, error) {
	// for ind, attrval := range al.valuesList {
	// 	if reflect.DeepEqual(dav, attrval) {
	// 		return fmt.Sprintf(":%d", ind), nil
	// 	}
	// }

	// If deduplicating, uncomment above and there should be an error message here
	// since all the aliases should be taken care of beforehand in another tree
	// traversal
	al.valuesList = append(al.valuesList, dav)
	return fmt.Sprintf(":%d", len(al.valuesList)-1), nil
}

func (al *aliasList) aliasPath(nm string) (string, error) {
	for ind, name := range al.namesList {
		if nm == name {
			return fmt.Sprintf("#%d", ind), nil
		}
	}
	al.namesList = append(al.namesList, nm)
	return fmt.Sprintf("#%d", len(al.namesList)-1), nil
}

func mergeExpressionMaps(lists ...[]Expression) Expression {
	ret := Expression{}
	for _, list := range lists {
		for _, expr := range list {
			for alias, name := range expr.Names {
				if ret.Names == nil {
					ret.Names = make(map[string]*string)
				}
				ret.Names[alias] = name
			}

			for alias, value := range expr.Values {
				if ret.Values == nil {
					ret.Values = make(map[string]*dynamodb.AttributeValue)
				}
				ret.Values[alias] = value
			}
		}
	}
	return ret
}
