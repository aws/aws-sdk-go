package expression

import (
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
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
// type Expression struct {
// 	Names      map[string]*string
// 	Values     map[string]*dynamodb.AttributeValue
// 	Expression string
// }
type Expression struct {
	expressionMap map[string]ExpressionTreeBuilder
}

type ExpressionTreeBuilder interface {
	BuildExpressionTree() (ExprNode, error)
}

func (expression Expression) ConditionExpression() *string {
	aliasList := &AliasList{}
	formattedExpressions := map[string]string{}
	keys := []string{}

	for key := range expression.expressionMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		exprNode, err := expression.expressionMap[key].BuildExpressionTree()
		if err != nil {
			return nil
		}
		formattedExpression, err := exprNode.BuildExpression(aliasList)
		if err != nil {
			return nil
		}
		formattedExpressions[key] = formattedExpression
	}

	return aws.String(formattedExpressions["condition"])
}

func (expression Expression) Names() map[string]*string {
	aliasList := &AliasList{}
	keys := []string{}

	for key := range expression.expressionMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		exprNode, err := expression.expressionMap[key].BuildExpressionTree()
		if err != nil {
			return nil
		}
		_, err = exprNode.BuildExpression(aliasList)
		if err != nil {
			return nil
		}
	}

	namesMap := map[string]*string{}
	for ind, val := range aliasList.namesList {
		namesMap[fmt.Sprintf("#%v", ind)] = aws.String(val)
	}

	return namesMap
}

func (expression Expression) Values() map[string]*dynamodb.AttributeValue {
	aliasList := &AliasList{}
	keys := []string{}

	for key := range expression.expressionMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		exprNode, err := expression.expressionMap[key].BuildExpressionTree()
		if err != nil {
			return nil
		}
		_, err = exprNode.BuildExpression(aliasList)
		if err != nil {
			return nil
		}
	}

	valuesMap := map[string]*dynamodb.AttributeValue{}
	for i := 0; i < len(aliasList.valuesList); i++ {
		valuesMap[fmt.Sprintf(":%v", i)] = &aliasList.valuesList[i]
	}
	// for ind, val := range aliasList.valuesList {
	// 	// fmt.Printf("%#v\n", val)
	// 	fmt.Printf("at ind %#v, got %#v\n", ind, &val)
	// 	valuesMap[fmt.Sprintf(":%v", ind)] = &val
	// }

	return valuesMap
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

// AliasList will keep track of all the names we need to alias in the nested
// struct of conditions and operands. This will allow each alias to be unique.
// AliasList will be passed in as a pointer when buildExprNodes is called in
// order to deduplicate all names within the tree strcuture of the ExprNodes.
type AliasList struct {
	namesList  []string
	valuesList []dynamodb.AttributeValue
}

// buildExpression returns an Expression with aliasing for paths/values
// specified by AliasList
func (en ExprNode) BuildExpression(al *AliasList) (string, error) {
	if al == nil {
		return "", fmt.Errorf("buildExprNodes error: AliasList is nil")
	}

	// Since each ExprNode contains a slice of names, values, and children that
	// correspond to the escaped characters, we an index to traverse the slices
	index := struct {
		name, value, children int
	}{}

	formattedExpression := en.fmtExpr

	for i := 0; i < len(formattedExpression); {
		if formattedExpression[i] != '$' {
			i++
			continue
		}

		if i == len(formattedExpression)-1 {
			return "", fmt.Errorf("buildExprNode error: invalid escape character")
		}

		var alias string
		var err error
		// if an escaped character is found, substitute it with the proper alias
		// TODO consider AST instead of string in the future
		switch formattedExpression[i+1] {
		case 'p':
			alias, err = substitutePath(index.name, en, al)
			if err != nil {
				return "", err
			}
			index.name++

		case 'v':
			alias, err = substituteValue(index.value, en, al)
			if err != nil {
				return "", err
			}
			index.value++

		case 'c':
			alias, err = substituteChild(index.children, en, al)
			if err != nil {
				return "", err
			}
			index.children++

		default:
			return "", fmt.Errorf("buildExprNode error: invalid escape rune %#v", formattedExpression[i+1])
		}
		formattedExpression = formattedExpression[:i] + alias + formattedExpression[i+2:]
		i += len(alias)
	}

	return formattedExpression, nil
}

// substitutePath will substitute the escaped character $p with the appropriate
// alias.
func substitutePath(index int, en ExprNode, al *AliasList) (string, error) {
	if index >= len(en.names) {
		return "", fmt.Errorf("substitutePath error: ExprNode []names out of range")
	}
	str, err := al.aliasPath(en.names[index])
	if err != nil {
		return "", err
	}
	return str, nil
}

// substituteValue will substitute the escaped character $v with the appropriate
// alias.
func substituteValue(index int, en ExprNode, al *AliasList) (string, error) {
	if index >= len(en.values) {
		return "", fmt.Errorf("substituteValue error: ExprNode []values out of range")
	}
	str, err := al.aliasValue(en.values[index])
	if err != nil {
		return "", err
	}
	return str, nil
}

// substituteChild will substitute the escaped character $c with the appropriate
// alias.
func substituteChild(index int, en ExprNode, al *AliasList) (string, error) {
	if index >= len(en.children) {
		return "", fmt.Errorf("substituteChild error: ExprNode []children out of range")
	}
	return en.children[index].BuildExpression(al)
	// if err != nil {
	// 	return "", err
	// }
	// str := childExpr.Expression
	// tempExpr := expr.Expression
	// *expr = MergeMaps(*expr, childExpr)
	// expr.Expression = tempExpr
	// return str, nil
}

// aliasValue returns the corresponding alias to the dav value argument. Since
// values are not deduplicated as of now, all values are just appended to the
// AliasList and given the index as the alias.
func (al *AliasList) aliasValue(dav dynamodb.AttributeValue) (string, error) {
	// for ind, attrval := range al.valuesList {
	// 	if reflect.DeepEqual(dav, attrval) {
	// 		return fmt.Sprintf(":%d", ind), nil
	// 	}
	// }

	if al == nil {
		return "", fmt.Errorf("aliasValue error: AliasList is nil")
	}

	// If deduplicating, uncomment above and there should be an error message here
	// since all the aliases should be taken care of beforehand in another tree
	// traversal
	al.valuesList = append(al.valuesList, dav)
	return fmt.Sprintf(":%d", len(al.valuesList)-1), nil
}

// aliasPath returns the corresponding alias to the argument string. The
// argument is checked against all existing AliasList names in order to avoid
// duplicate strings getting two different aliases.
func (al *AliasList) aliasPath(nm string) (string, error) {
	if al == nil {
		return "", fmt.Errorf("aliasValue error: AliasList is nil")
	}

	for ind, name := range al.namesList {
		if nm == name {
			return fmt.Sprintf("#%d", ind), nil
		}
	}
	al.namesList = append(al.namesList, nm)
	return fmt.Sprintf("#%d", len(al.namesList)-1), nil
}

// MergeMaps merges maps of multiple Expressions. This is used to
// combine the maps created by the child nodes. It is also used to combine maps
// in order to inject the resulting maps into operation inputs. MergeMaps
// assumes that the inputs are all valid.
//
// Example:
//
//     filterExpr, err := expression.Path("foo").Equal(expression.Value(5))
//     projExpr, err := expression.Projection(expression.Path("bar"), expression.Path("baz"))
//
//     scanInput := dynamodb.ScanInput{
//       FilterExpression:          aws.String(filterExpr.Expression),
//       ProjectionExpression:      aws.String(projExpr.Expression),
//       ExpressionAttributeNames:  MergeMaps(filterExpr, projExpr).Names,
//       ExpressionAttributeValues: MergeMaps(filterExpr, projExpr).Values,
//       Key: map[string]*dynamodb.AttributeValue{
//         "PartitionKey": &dynamodb.AttributeValue{
//           S: aws.String("SomeKey"),
//         },
//       },
//       TableName: aws.String("SomeTable"),
//     }
// func MergeMaps(list ...Expression) Expression {
// 	ret := Expression{}
// 	for _, expr := range list {
// 		for alias, name := range expr.Names {
// 			if ret.Names == nil {
// 				ret.Names = map[string]*string{}
// 			}
// 			ret.Names[alias] = name
// 		}
//
// 		for alias, value := range expr.Values {
// 			if ret.Values == nil {
// 				ret.Values = map[string]*dynamodb.AttributeValue{}
// 			}
// 			ret.Values[alias] = value
// 		}
// 	}
// 	return ret
// }
