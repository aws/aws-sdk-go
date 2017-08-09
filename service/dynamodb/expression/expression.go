package expression

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

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
