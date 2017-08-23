package expression

import (
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Builder contains a list of structs fulfilling the
// ExpressionTreeBuilder interface, representing the different types of
// Expressions that make up the total Expression as a whole. Builder
// will be used to fill the members of various DynamoDB input structs.
//
// Example:
//
//     // let expr be an instance of Builder{}
//
//     deleteInput := dynamodb.DeleteItemInput{
//       ConditionExpression:       expr.Condition(),
//       ExpressionAttributeNames:  expr.Names(),
//       ExpressionAttributeValues: expr.Values(),
//       Key: map[string]*dynamodb.AttributeValue{
//         "PartitionKey": &dynamodb.AttributeValue{
//           S: aws.String("SomeKey"),
//         },
//       },
//       TableName: aws.String("SomeTable"),
//     }
type Builder struct {
	expressionMap map[string]TreeBuilder
}

// TreeBuilder interface will be fulfilled by builder structs that
// represent different types of Expressions.
type TreeBuilder interface {
	// BuildTree will create the tree structure of ExprNodes. The tree structure
	// of ExprNodes will be traversed in order to build the string representing
	// different types of Expressions as well as the maps that represent
	// ExpressionAttributeNames and ExpressionAttributeValues.
	BuildTree() (ExprNode, error)
}

// returnExpression will return *string corresponding to the type of Expression
// string specified by the expressionType.
func (expression Builder) returnExpression(expressionType string) *string {
	_, formattedExpressions, err := expression.buildChildBuilders()
	if err != nil {
		return nil
	}

	return aws.String(formattedExpressions[expressionType])
}

// Condition will return the *string corresponding to the Condition
// Expression of the argument Builder. This method is used to satisfy
// the members of DynamoDB input structs.
//
// Example:
//
//     // let expr be an instance of Builder{}
//
//     deleteInput := dynamodb.DeleteItemInput{
//       ConditionExpression:       expr.Condition(),
//       ExpressionAttributeNames:  expr.Names(),
//       ExpressionAttributeValues: expr.Values(),
//       Key: map[string]*dynamodb.AttributeValue{
//         "PartitionKey": &dynamodb.AttributeValue{
//           S: aws.String("SomeKey"),
//         },
//       },
//       TableName: aws.String("SomeTable"),
//     }
func (expression Builder) Condition() *string {
	return expression.returnExpression("condition")
}

// Projection will return the *string corresponding to the Projection
// Expression of the argument Builder. This method is used to satisfy
// the members of DynamoDB input structs.
//
// Example:
//
//     // let expr be an instance of Builder{}
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    expr.KeyCondition(),
//       ProjectionExpression:      expr.Projection(),
//       ExpressionAttributeNames:  expr.Names(),
//       ExpressionAttributeValues: expr.Values(),
//       TableName: aws.String("SomeTable"),
//     }
func (expression Builder) Projection() *string {
	return expression.returnExpression("projection")
}

// Names will return the map[string]*string corresponding to the
// ExpressionAttributeNames of the argument Builder. This method is
// used to satisfy the members of DynamoDB input structs.
//
// Example:
//
//     // let expr be an instance of Builder{}
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    expr.KeyCondition(),
//       ProjectionExpression:      expr.Projection(),
//       ExpressionAttributeNames:  expr.Names(),
//       ExpressionAttributeValues: expr.Values(),
//       TableName: aws.String("SomeTable"),
//     }
func (expression Builder) Names() map[string]*string {
	aliasList, _, err := expression.buildChildBuilders()
	if err != nil {
		return nil
	}

	namesMap := map[string]*string{}
	for ind, val := range aliasList.namesList {
		namesMap[fmt.Sprintf("#%v", ind)] = aws.String(val)
	}

	return namesMap
}

// Values will return the map[string]*dynamodb.AttributeValue corresponding to
// the ExpressionAttributeValues of the argument Builder. This method
// is used to satisfy the members of DynamoDB input structs.
//
// Example:
//
//     // let expr be an instance of Builder{}
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    expr.KeyCondition(),
//       ProjectionExpression:      expr.Projection(),
//       ExpressionAttributeNames:  expr.Names(),
//       ExpressionAttributeValues: expr.Values(),
//       TableName: aws.String("SomeTable"),
//     }
func (expression Builder) Values() map[string]*dynamodb.AttributeValue {
	aliasList, _, err := expression.buildChildBuilders()
	if err != nil {
		return nil
	}

	valuesMap := map[string]*dynamodb.AttributeValue{}
	for i := 0; i < len(aliasList.valuesList); i++ {
		valuesMap[fmt.Sprintf(":%v", i)] = &aliasList.valuesList[i]
	}

	return valuesMap
}

// ExprNode will be the generic nodes that will represent both Operands and
// Conditions. The purpose of ExprNode is to be able to call an generic
// recursive function on the top level ExprNode to be able to determine a root
// node in order to deduplicate name aliases.
// fmtExpr is a string that has escaped characters to refer to
// names/values/children which needs to be aliased at runtime in order to avoid
// duplicate values. The rules are as follows:
//     $n: Indicates that an alias of a name needs to be inserted. The corresponding
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

// BuildExpressionString returns a string with aliasing for names/values
// specified by AliasList. The string corresponds to the expression that the
// ExprNode tree represents.
func (exprNode ExprNode) BuildExpressionString(aliasList *AliasList) (string, error) {
	if aliasList == nil {
		return "", fmt.Errorf("buildExprNodes error: AliasList is nil")
	}

	// Since each ExprNode contains a slice of names, values, and children that
	// correspond to the escaped characters, we an index to traverse the slices
	index := struct {
		name, value, children int
	}{}

	formattedExpression := exprNode.fmtExpr

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
		case 'n':
			alias, err = substitutePath(index.name, exprNode, aliasList)
			if err != nil {
				return "", err
			}
			index.name++

		case 'v':
			alias, err = substituteValue(index.value, exprNode, aliasList)
			if err != nil {
				return "", err
			}
			index.value++

		case 'c':
			alias, err = substituteChild(index.children, exprNode, aliasList)
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

// substitutePath will substitute the escaped character $n with the appropriate
// alias.
func substitutePath(index int, exprNode ExprNode, aliasList *AliasList) (string, error) {
	if index >= len(exprNode.names) {
		return "", fmt.Errorf("substitutePath error: ExprNode []names out of range")
	}
	str, err := aliasList.aliasPath(exprNode.names[index])
	if err != nil {
		return "", err
	}
	return str, nil
}

// substituteValue will substitute the escaped character $v with the appropriate
// alias.
func substituteValue(index int, exprNode ExprNode, aliasList *AliasList) (string, error) {
	if index >= len(exprNode.values) {
		return "", fmt.Errorf("substituteValue error: ExprNode []values out of range")
	}
	str, err := aliasList.aliasValue(exprNode.values[index])
	if err != nil {
		return "", err
	}
	return str, nil
}

// substituteChild will substitute the escaped character $c with the appropriate
// alias.
func substituteChild(index int, exprNode ExprNode, aliasList *AliasList) (string, error) {
	if index >= len(exprNode.children) {
		return "", fmt.Errorf("substituteChild error: ExprNode []children out of range")
	}
	return exprNode.children[index].BuildExpressionString(aliasList)
}

// aliasValue returns the corresponding alias to the dav value argument. Since
// values are not deduplicated as of now, all values are just appended to the
// AliasList and given the index as the alias.
func (aliasList *AliasList) aliasValue(dav dynamodb.AttributeValue) (string, error) {
	if aliasList == nil {
		return "", fmt.Errorf("aliasValue error: AliasList is nil")
	}

	aliasList.valuesList = append(aliasList.valuesList, dav)
	return fmt.Sprintf(":%d", len(aliasList.valuesList)-1), nil
}

// aliasPath returns the corresponding alias to the argument string. The
// argument is checked against all existing AliasList names in order to avoid
// duplicate strings getting two different aliases.
func (aliasList *AliasList) aliasPath(nm string) (string, error) {
	if aliasList == nil {
		return "", fmt.Errorf("aliasValue error: AliasList is nil")
	}

	for ind, name := range aliasList.namesList {
		if nm == name {
			return fmt.Sprintf("#%d", ind), nil
		}
	}
	aliasList.namesList = append(aliasList.namesList, nm)
	return fmt.Sprintf("#%d", len(aliasList.namesList)-1), nil
}

// buildChildBuilders will compile the list of ExpressionTreeBuilders that
// are the children of the argument Builder. The returned AliasList
// will represent all the alias tokens used in the expression strings. The
// returned map[string]string will map the type of expression (i.e. "condition",
// "update") to the appropriate expression string.
func (expression Builder) buildChildBuilders() (*AliasList, map[string]string, error) {
	aliasList := &AliasList{}
	formattedExpressions := map[string]string{}
	keys := []string{}

	for key := range expression.expressionMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		exprNode, err := expression.expressionMap[key].BuildTree()
		if err != nil {
			return nil, nil, err
		}
		formattedExpression, err := exprNode.BuildExpressionString(aliasList)
		if err != nil {
			return nil, nil, err
		}
		formattedExpressions[key] = formattedExpression
	}

	return aliasList, formattedExpressions, nil
}
