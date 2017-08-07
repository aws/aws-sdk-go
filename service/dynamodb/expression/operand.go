package expression

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

// Expression contains the map of aliases to names/values, representing the
// ExpressionAttributeNames and ExpressionAttributeValues, that is needed in
// order for the Expression string to be used in an operation input.
// (i.e. UpdateItemInput, DeleteItemInput, etc)
//
// Example:
//
//     // let expr be an instance of Expression{}
//
//     deleteInput := dynamodb.DeleteItemInput{
//       ConditionExpression:       aws.String(expr.Expression),
//       ExpressionAttributeNames:  expr.Names,
//       ExpressionAttributeValues: expr.Values,
//       Key: map[string]*dynamodb.AttributeValue{
//         "PartitionKey": &dynamodb.AttributeValue{
//           S: aws.String("SomeKey"),
//         },
//       },
//       TableName: aws.String("SomeTable"),
//     }
type Expression struct {
	Names      map[string]*string
	Values     map[string]*dynamodb.AttributeValue
	Expression string
}

// ExprNode will be the generic nodes that will represent both Operands and
// Conditions. The purpose of ExprNode is to be able to call an generic
// recursive function on the top level ExprNode to be able to determine a root
// node in order to deduplicate name aliases.
// fmtExpr is a string that has escaped characters to refer to
// names/values/children which needs to be aliased at runtime in order to avoid
// duplicate values. The rules are as follows:
//     $p: Indicates that an alias of a name needs to be inserted. The corresponding
//         name to be aliased will be in the []names slice.
//     $v: Indicates that an alias of a value needs to be inserted. The
//         corresponding value to be aliased will be in the []values slice.
//     $c: Indicates that the fmtExpr of a child ExprNode needs to be inserted. The
//         corresponding child node is in the []children slice.
type ExprNode struct {
	names    []string
	values   []dynamodb.AttributeValue
	children []ExprNode
	fmtExpr  string
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
// method to call on, not for users to invoke.
func (p PathBuilder) BuildOperand() (ExprNode, error) {
	if p.path == "" {
		return ExprNode{}, fmt.Errorf("BuildOperand error: path is empty")
	}

	ret := ExprNode{
		names: []string{},
	}

	nameSplit := strings.Split(p.path, ".")
	fmtNames := make([]string, 0, len(nameSplit))

	for _, word := range nameSplit {
		var substr string
		if word == "" {
			return ExprNode{}, fmt.Errorf("BuildOperand error: path is empty")
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
			return ExprNode{}, fmt.Errorf("BuildOperand error: invalid path index")
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

// aliasList will keep track of all the names we need to alias in the nested
// struct of conditions and operands. This will allow each alias to be unique.
// aliasList will be passed in as a pointer when buildExprNodes is called in
// order to deduplicate all names within the tree strcuture of the ExprNodes.
type aliasList struct {
	namesList  []string
	valuesList []dynamodb.AttributeValue
}

// buildExpression returns an Expression with aliasing for paths/values
// specified by aliasList
func (en ExprNode) buildExprNodes(al *aliasList) (Expression, error) {
	if al == nil {
		return Expression{}, fmt.Errorf("buildExprNodes error: aliasList is nil")
	}

	// Since each ExprNode contains a slice of names, values, and children that
	// correspond to the escaped characters, we an index to traverse the slices
	index := struct {
		name, value, children int
	}{}

	expr := Expression{
		Expression: en.fmtExpr,
	}

	for i := 0; i < len(expr.Expression); {
		if expr.Expression[i] != '$' {
			i++
			continue
		}

		if i == len(expr.Expression)-1 {
			return Expression{}, fmt.Errorf("buildExprNode error: invalid escape character")
		}

		var alias string
		// if an escaped character is found, substitute it with the proper alias
		// TODO consider AST instead of string in the future
		switch expr.Expression[i+1] {
		case 'p':
			if index.name >= len(en.names) {
				return Expression{}, fmt.Errorf("buildExprNodes error: ExprNode []names out of range")
			}
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
			if index.value >= len(en.values) {
				return Expression{}, fmt.Errorf("buildExprNodes error: ExprNode []values out of range")
			}
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
			if index.children >= len(en.children) {
				return Expression{}, fmt.Errorf("buildExprNodes error: ExprNode []children out of range")
			}
			childExpr, err := en.children[index.children].buildExprNodes(al)
			if err != nil {
				return Expression{}, err
			}
			alias = childExpr.Expression
			tempExpr := expr.Expression
			expr, err = mergeExpressionMaps([]Expression{expr, childExpr})
			if err != nil {
				return Expression{}, err
			}
			expr.Expression = tempExpr
			index.children++

		default:
			return Expression{}, fmt.Errorf("buildExprNode error: invalid escape rune %#v", expr.Expression[i+1])
		}
		expr.Expression = expr.Expression[:i] + alias + expr.Expression[i+2:]
		i += len(alias)
	}

	return expr, nil
}

// aliasValue returns the corresponding alias to the dav value argument. Since
// values are not deduplicated as of now, all values are just appended to the
// aliasList and given the index as the alias.
func (al *aliasList) aliasValue(dav dynamodb.AttributeValue) (string, error) {
	// for ind, attrval := range al.valuesList {
	// 	if reflect.DeepEqual(dav, attrval) {
	// 		return fmt.Sprintf(":%d", ind), nil
	// 	}
	// }

	if al == nil {
		return "", fmt.Errorf("aliasValue error: aliasList is nil")
	}

	// If deduplicating, uncomment above and there should be an error message here
	// since all the aliases should be taken care of beforehand in another tree
	// traversal
	al.valuesList = append(al.valuesList, dav)
	return fmt.Sprintf(":%d", len(al.valuesList)-1), nil
}

// aliasPath returns the corresponding alias to the argument string. The
// argument is checked against all existing aliasList names in order to avoid
// duplicate strings getting two different aliases.
func (al *aliasList) aliasPath(nm string) (string, error) {
	if al == nil {
		return "", fmt.Errorf("aliasValue error: aliasList is nil")
	}

	for ind, name := range al.namesList {
		if nm == name {
			return fmt.Sprintf("#%d", ind), nil
		}
	}
	al.namesList = append(al.namesList, nm)
	return fmt.Sprintf("#%d", len(al.namesList)-1), nil
}

// mergeExpressionMaps merges maps of multiple Expressions. This is used to
// combine the maps created by the child nodes
func mergeExpressionMaps(lists ...[]Expression) (Expression, error) {
	ret := Expression{}
	for _, list := range lists {
		for _, expr := range list {
			if reflect.DeepEqual(expr, (Expression{})) {
				return Expression{}, fmt.Errorf("mergeExpressionMaps error: expression is unset")
			}
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
	return ret, nil
}
