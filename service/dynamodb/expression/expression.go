package expression

import (
	"fmt"

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
		var err error
		// if an escaped character is found, substitute it with the proper alias
		// TODO consider AST instead of string in the future
		switch expr.Expression[i+1] {
		case 'p':
			alias, err = substitutePath(index.name, en, &expr, al)
			if err != nil {
				return Expression{}, err
			}
			index.name++

		case 'v':
			alias, err = substituteValue(index.value, en, &expr, al)
			if err != nil {
				return Expression{}, err
			}
			index.value++

		case 'c':
			alias, err = substituteChild(index.children, en, &expr, al)
			if err != nil {
				return Expression{}, err
			}
			index.children++

		default:
			return Expression{}, fmt.Errorf("buildExprNode error: invalid escape rune %#v", expr.Expression[i+1])
		}
		expr.Expression = expr.Expression[:i] + alias + expr.Expression[i+2:]
		i += len(alias)
	}

	return expr, nil
}

// substitutePath will substitute the escaped character $p with the appropriate
// alias.
func substitutePath(index int, en ExprNode, expr *Expression, al *aliasList) (string, error) {
	if index >= len(en.names) {
		return "", fmt.Errorf("substitutePath error: ExprNode []names out of range")
	}
	str, err := al.aliasPath(en.names[index])
	if err != nil {
		return "", err
	}
	if expr.Names == nil {
		expr.Names = map[string]*string{}
	}
	expr.Names[str] = &en.names[index]
	return str, nil
}

// substituteValue will substitute the escaped character $v with the appropriate
// alias.
func substituteValue(index int, en ExprNode, expr *Expression, al *aliasList) (string, error) {
	if index >= len(en.values) {
		return "", fmt.Errorf("substituteValue error: ExprNode []values out of range")
	}
	str, err := al.aliasValue(en.values[index])
	if err != nil {
		return "", err
	}
	if expr.Values == nil {
		expr.Values = map[string]*dynamodb.AttributeValue{}
	}
	expr.Values[str] = &en.values[index]
	return str, nil
}

// substituteChild will substitute the escaped character $c with the appropriate
// alias.
func substituteChild(index int, en ExprNode, expr *Expression, al *aliasList) (string, error) {
	if index >= len(en.children) {
		return "", fmt.Errorf("substituteChild error: ExprNode []children out of range")
	}
	childExpr, err := en.children[index].buildExprNodes(al)
	if err != nil {
		return "", err
	}
	str := childExpr.Expression
	tempExpr := expr.Expression
	*expr = MergeMaps(*expr, childExpr)
	expr.Expression = tempExpr
	return str, nil
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
func MergeMaps(list ...Expression) Expression {
	ret := Expression{}
	for _, expr := range list {
		for alias, name := range expr.Names {
			if ret.Names == nil {
				ret.Names = map[string]*string{}
			}
			ret.Names[alias] = name
		}

		for alias, value := range expr.Values {
			if ret.Values == nil {
				ret.Values = map[string]*dynamodb.AttributeValue{}
			}
			ret.Values[alias] = value
		}
	}
	return ret
}
