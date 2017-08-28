package expression

import (
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ErrEmptyBuilder is an error that is returned if Build() is called on an empty
// Builder.
var ErrEmptyBuilder = awserr.New("EmptyBuilder", "Build error: the argument Builder is empty", nil)

// expressionType will specify the type of Expression. The const is used to
// eliminate magic strings
type expressionType string

const (
	projection   expressionType = "projection"
	keyCondition                = "keyCondition"
	condition                   = "condition"
	filter                      = "filter"
	update                      = "update"
)

// Implementing the Sort interface
type typeList []expressionType

func (l typeList) Len() int {
	return len(l)
}

func (l typeList) Less(i, j int) bool {
	return string(l[i]) < string(l[j])
}

func (l typeList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Builder contains a list of structs fulfilling the treeBuilder interface,
// representing the different types of Expressions that make up the total
// Expression as a whole. Builder will have methods corresponding to different
// types of expressions (WithUpdate(), WithCondition(), WithFilter(), etc) that
// will allow users to add expressions to the input member Expression that
// Builder creates. Builder will have a method Build() which will build the
// Expression struct which can be used to produce members of DynamoDB input
// structs.
// Builder is to be created with functions, not to be initialized.
//
// Example:
//
//     keyCond := expression.Key("someKey").Equal(expression.Value("someValue"))
//     proj := expression.NamesList(expression.Name("aName"), expression.Name("anotherName"), expression.Name("oneOtherName"))
//     builder := expression.NewBuilder().WithKeyCondition(keyCond).WithProjection(proj)
//     expression := builder.Build()
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    expression.KeyCondition(),
//       ProjectionExpression:      expression.Projection(),
//       ExpressionAttributeNames:  expression.Names(),
//       ExpressionAttributeValues: expression.Values(),
//       TableName: aws.String("SomeTable"),
//     }
type Builder struct {
	expressionMap map[expressionType]treeBuilder
}

// NewBuilder returns an empty Builder struct. The NewBuilder function can be
// used to functionally build the Builder.
//
// Example:
//
//     keyCond := expression.Key("someKey").Equal(expression.Value("someValue"))
//     proj := expression.NamesList(expression.Name("aName"), expression.Name("anotherName"), expression.Name("oneOtherName"))
//     builder := expression.NewBuilder().WithKeyCondition(keyCond).WithProjection(proj)
func NewBuilder() Builder {
	return Builder{}
}

// Build builds a Expression struct with the same expressionMap as the argument
// Builder. The aim of this method is to check the child treeBuilders for their
// formats and return an error. This makes sure that any Expression has the
// right format or is an empty struct
// Calling Build on an empty Builder will return a typed error ErrEmptyBuilder.
func (b Builder) Build() (Expression, error) {
	if b.expressionMap == nil {
		return Expression{}, ErrEmptyBuilder
	}

	aliasList, expressionMap, err := b.buildChildTrees()
	if err != nil {
		return Expression{}, err
	}

	expression := Expression{
		expressionMap: expressionMap,
	}

	if len(aliasList.namesList) != 0 {
		namesMap := map[string]*string{}
		for ind, val := range aliasList.namesList {
			namesMap[fmt.Sprintf("#%v", ind)] = aws.String(val)
		}
		expression.namesMap = namesMap
	}

	if len(aliasList.valuesList) != 0 {
		valuesMap := map[string]*dynamodb.AttributeValue{}
		for i := 0; i < len(aliasList.valuesList); i++ {
			valuesMap[fmt.Sprintf(":%v", i)] = &aliasList.valuesList[i]
		}
		expression.valuesMap = valuesMap
	}

	return expression, nil
}

// buildChildTrees will compile the list of treeBuilders that are the children
// of the argument Builder. The returned aliasList will represent all the
// alias tokens used in the expression strings. The returned map[string]string
// will map the type of expression (i.e. "condition", "update") to the
// appropriate expression string.
func (b Builder) buildChildTrees() (*aliasList, map[expressionType]string, error) {
	aliasList := &aliasList{}
	formattedExpressions := map[expressionType]string{}
	keys := typeList{}

	for expressionType := range b.expressionMap {
		keys = append(keys, expressionType)
	}

	sort.Sort(keys)

	for _, key := range keys {
		node, err := b.expressionMap[key].buildTree()
		if err != nil {
			return nil, nil, err
		}
		formattedExpression, err := node.buildExpressionString(aliasList)
		if err != nil {
			return nil, nil, err
		}
		formattedExpressions[key] = formattedExpression
	}

	return aliasList, formattedExpressions, nil
}

// WithCondition method will add the argument ConditionBuilder as a treeBuilder
// to the argument Builder. If the argument Builder already has a
// ConditionBuilder under the key condition, WithCondition() will overwrite the
// existing ConditionBuilder. Users will able to add other treeBuilders to the
// Builder or call Build() to build a Expression struct.
//
// Example:
//
//     // let builder be an existing Builder{} and
//     // conditionBuilder be an existing ConditionBuilder{}
//     builder = builder.WithCondition(conditionBuilder)
//
//     expression := builder.Build()
func (b Builder) WithCondition(conditionBuilder ConditionBuilder) Builder {
	if b.expressionMap == nil {
		b.expressionMap = map[expressionType]treeBuilder{}
	}
	b.expressionMap[condition] = conditionBuilder
	return b
}

// WithProjection method will add the argument ProjectionBuilder as a
// treeBuilder to the argument Builder. If the argument Builder already has a
// ProjectionBuilder, WithProjection() will overwrite the existing
// ProjectionBuilder. Users will able to add other treeBuilders to the Builder
// or call Build() to build a Expression struct.
//
// Example:
//
//     // let builder be an existing Builder{} and
//     // projectionBuilder be an existing projectionBuilder{}
//     builder = builder.WithProjection(projectionBuilder)
//
//     expression := builder.Build()
func (b Builder) WithProjection(projectionBuilder ProjectionBuilder) Builder {
	if b.expressionMap == nil {
		b.expressionMap = map[expressionType]treeBuilder{}
	}
	b.expressionMap[projection] = projectionBuilder
	return b
}

// // WithKeyCondition method will add the argument KeyConditionBuilder as a treeBuilder to
// // the argument Builder. If the argument Builder already has a
// // KeyConditionBuilder, WithKeyCondition() will overwrite the existing KeyConditionBuilder.
// // Users will able to add other treeBuilders to the Builder or call
// // Build() to build a Expression struct.
// //
// // Example:
// //
// //     // let builder be an existing Builder{} and
// //     // keyConditionBuilder be an existing keyConditionBuilder{}
// //     builder = builder.WithKeyCondition(keyConditionBuilder)
// //
// //     expression := builder.Build()
// func (b Builder) WithKeyCondition(keyConditionBuilder KeyConditionBuilder) Builder {
// 	if b.expressionMap == nil {
// 		b.expressionMap = map[expressionType]treeBuilder{}
// 	}
// 	b.expressionMap[keyCondition] = keyConditionBuilder
// 	return b
// }
//
// // WithKeyCondition function will create a Builder with the argument
// // keyConditionBuilder as a child treeBuilder. Users will able to add other
// // treeBuilders to the Builder or call Build() to build a Expression
// // struct.
// //
// // Example:
// //
// //     // let keyConditionBuilder and conditionBuilder be an existing
// //     // KeyConditionBuilder and ConditionBuilder respectively.
// //     builder := expression.WithKeyCondition(keyConditionBuilder)
// //
// //     builder = builder.WithCondition(conditionBuilder)   // Adding a ConditionBuilder
// //     expression := builder.Build()                      // Creating a Expression
// func WithKeyCondition(keyConditionBuilder KeyConditionBuilder) Builder {
// 	ret := Builder{}
// 	return ret.WithKeyCondition(keyConditionBuilder)
// }

// WithFilter method will add the argument ConditionBuilder as a treeBuilder to
// the argument Builder. If the argument Builder already has a
// ConditionBuilder under the key filter, WithFilter() will overwrite the
// existing ConditionBuilder. Users will able to add other treeBuilders to the
// Builder or call Build() to build a Expression struct.
//
// Example:
//
//     // let builder be an existing Builder{} and
//     // filterBuilder be an existing filterBuilder{}
//     builder = builder.WithFilter(filterBuilder)
//
//     expression := builder.Build()
func (b Builder) WithFilter(filterBuilder ConditionBuilder) Builder {
	if b.expressionMap == nil {
		b.expressionMap = map[expressionType]treeBuilder{}
	}
	b.expressionMap[filter] = filterBuilder
	return b
}

// // WithUpdate method will add the argument UpdateBuilder as a treeBuilder to
// // the argument Builder. If the argument Builder already has a
// // UpdateBuilder, WithUpdate() will overwrite the existing UpdateBuilder.
// // Users will able to add other treeBuilders to the Builder or call
// // Build() to build a Expression struct.
// //
// // Example:
// //
// //     // let builder be an existing Builder{} and
// //     // updateBuilder be an existing updateBuilder{}
// //     builder = builder.WithUpdate(updateBuilder)
// //
// //     expression := builder.Build()
// func (b Builder) WithUpdate(updateBuilder UpdateBuilder) Builder {
// 	if b.expressionMap == nil {
// 		b.expressionMap = map[expressionType]treeBuilder{}
// 	}
// 	b.expressionMap[update] = updateBuilder
// 	return b
// }
//
// // WithUpdate function will create a Builder with the argument
// // updateBuilder as a child treeBuilder. Users will able to add other
// // treeBuilders to the Builder or call Build() to build a Expression
// // struct.
// //
// // Example:
// //
// //     // let updateBuilder and conditionBuilder be an existing
// //     // UpdateBuilder and ConditionBuilder respectively.
// //     builder := expression.WithUpdate(updateBuilder)
// //
// //     builder = builder.WithCondition(conditionBuilder)   // Adding a ConditionBuilder
// //     expression := builder.Build()                      // Creating a Expression
// func WithUpdate(updateBuilder UpdateBuilder) Builder {
// 	ret := Builder{}
// 	return ret.WithUpdate(updateBuilder)
// }

// Expression will be a struct that will be able to generate members to DynamoDB
// inputs.
// The idea of Builder and Expression is separated to be able to check the
// format of the child treeBuilders and to return an error with the Build()
// method for ExpressionBuilder.
//
// Example:
//
//     keyCond := expression.Key("someKey").Equal(expression.Value("someValue"))
//     proj := expression.NamesList(expression.Name("aName"), expression.Name("anotherName"), expression.Name("oneOtherName"))
//     builder := expression.WithKeyCondition(keyCond).WithProjection(proj)
//     expression := builder.Build()
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    expression.KeyCondition(),
//       ProjectionExpression:      expression.Projection(),
//       ExpressionAttributeNames:  expression.Names(),
//       ExpressionAttributeValues: expression.Values(),
//       TableName: aws.String("SomeTable"),
//     }
type Expression struct {
	expressionMap map[expressionType]string
	namesMap      map[string]*string
	valuesMap     map[string]*dynamodb.AttributeValue
}

// treeBuilder interface will be fulfilled by builder structs that represent
// different types of Expressions.
type treeBuilder interface {
	// buildTree will create the tree structure of exprNodes. The tree structure
	// of exprNodes will be traversed in order to build the string representing
	// different types of Expressions as well as the maps that represent
	// ExpressionAttributeNames and ExpressionAttributeValues.
	buildTree() (exprNode, error)
}

// Condition will return the *string corresponding to the Condition Expression
// of the argument Expression. This method is used to satisfy the members of
// DynamoDB input structs. If the Expression does not have a condition
// expression this method will return nil.
//
// Example:
//
//     // let expression be an instance of Expression{}
//
//     deleteInput := dynamodb.DeleteItemInput{
//       ConditionExpression:       expression.Condition(),
//       ExpressionAttributeNames:  expression.Names(),
//       ExpressionAttributeValues: expression.Values(),
//       Key: map[string]*dynamodb.AttributeValue{
//         "PartitionKey": &dynamodb.AttributeValue{
//           S: aws.String("SomeKey"),
//         },
//       },
//       TableName: aws.String("SomeTable"),
//     }
func (e Expression) Condition() *string {
	return e.returnExpression(condition)
}

// Filter will return the *string corresponding to the Filter Expression of the
// argument Expression. This method is used to satisfy the members of DynamoDB
// input structs. If the Expression does not have a filter expression this
// method will return nil.
//
// Example:
//
//     // let expression be an instance of Expression{}
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    expression.KeyCondition(),
//       FilterExpression:          expression.Filter(),
//       ExpressionAttributeNames:  expression.Names(),
//       ExpressionAttributeValues: expression.Values(),
//       TableName: aws.String("SomeTable"),
//     }
func (e Expression) Filter() *string {
	return e.returnExpression(filter)
}

// Projection will return the *string corresponding to the Projection Expression
// of the argument Expression. This method is used to satisfy the members of
// DynamoDB input structs. If the Expression does not have a projection
// expression this method will return nil.
//
// Example:
//
//     // let expression be an instance of Expression{}
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    expression.KeyCondition(),
//       ProjectionExpression:      expression.Projection(),
//       ExpressionAttributeNames:  expression.Names(),
//       ExpressionAttributeValues: expression.Values(),
//       TableName: aws.String("SomeTable"),
//     }
func (e Expression) Projection() *string {
	return e.returnExpression(projection)
}

// Names will return the map[string]*string corresponding to the
// ExpressionAttributeNames of the argument Expression. This method is used to
// satisfy the members of DynamoDB input structs. If Expression does not use
// ExpressionAttributeNames, this method will return nil
//
// Example:
//
//     // let expression be an instance of Expression{}
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    expression.KeyCondition(),
//       ProjectionExpression:      expression.Projection(),
//       ExpressionAttributeNames:  expression.Names(),
//       ExpressionAttributeValues: expression.Values(),
//       TableName: aws.String("SomeTable"),
//     }
func (e Expression) Names() map[string]*string {
	return e.namesMap
}

// Values will return the map[string]*dynamodb.AttributeValue corresponding to
// the ExpressionAttributeValues of the argument Expression. This method is used
// to satisfy the members of DynamoDB input structs. If Expression does not use
// ExpressionAttributeValues, this method will return nil
// Example:
//
//     // let expression be an instance of Expression{}
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    expression.KeyCondition(),
//       ProjectionExpression:      expression.Projection(),
//       ExpressionAttributeNames:  expression.Names(),
//       ExpressionAttributeValues: expression.Values(),
//       TableName: aws.String("SomeTable"),
//     }
func (e Expression) Values() map[string]*dynamodb.AttributeValue {
	return e.valuesMap
}

// returnExpression will return *string corresponding to the type of Expression
// string specified by the expressionType. If there is no corresponding
// expression available in Expression, the method will return nil
func (e Expression) returnExpression(expressionType expressionType) *string {
	if e.expressionMap == nil {
		return nil
	}
	return aws.String(e.expressionMap[expressionType])
}

// exprNode will be the generic nodes that will represent both Operands and
// Conditions. The purpose of exprNode is to be able to call an generic
// recursive function on the top level exprNode to be able to determine a root
// node in order to deduplicate name aliases.
// fmtExpr is a string that has escaped characters to refer to
// names/values/children which needs to be aliased at runtime in order to avoid
// duplicate values. The rules are as follows:
//     $n: Indicates that an alias of a name needs to be inserted. The
//         corresponding name to be aliased will be in the []names slice.
//     $v: Indicates that an alias of a value needs to be inserted. The
//         corresponding value to be aliased will be in the []values slice.
//     $c: Indicates that the fmtExpr of a child exprNode needs to be inserted.
//         The corresponding child node is in the []children slice.
type exprNode struct {
	names    []string
	values   []dynamodb.AttributeValue
	children []exprNode
	fmtExpr  string
}

// aliasList will keep track of all the names we need to alias in the nested
// struct of conditions and operands. This will allow each alias to be unique.
// aliasList will be passed in as a pointer when buildChildTrees is called in
// order to deduplicate all names within the tree strcuture of the exprNodes.
type aliasList struct {
	namesList  []string
	valuesList []dynamodb.AttributeValue
}

// buildExpressionString returns a string with aliasing for names/values
// specified by aliasList. The string corresponds to the expression that the
// exprNode tree represents.
func (en exprNode) buildExpressionString(aliasList *aliasList) (string, error) {
	if aliasList == nil {
		return "", fmt.Errorf("buildExprNodes error: aliasList is nil")
	}

	// Since each exprNode contains a slice of names, values, and children that
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
			return "", fmt.Errorf("buildexprNode error: invalid escape character")
		}

		var alias string
		var err error
		// if an escaped character is found, substitute it with the proper alias
		// TODO consider AST instead of string in the future
		switch formattedExpression[i+1] {
		case 'n':
			alias, err = substitutePath(index.name, en, aliasList)
			if err != nil {
				return "", err
			}
			index.name++

		case 'v':
			alias, err = substituteValue(index.value, en, aliasList)
			if err != nil {
				return "", err
			}
			index.value++

		case 'c':
			alias, err = substituteChild(index.children, en, aliasList)
			if err != nil {
				return "", err
			}
			index.children++

		default:
			return "", fmt.Errorf("buildexprNode error: invalid escape rune %#v", formattedExpression[i+1])
		}
		formattedExpression = formattedExpression[:i] + alias + formattedExpression[i+2:]
		i += len(alias)
	}

	return formattedExpression, nil
}

// substitutePath will substitute the escaped character $n with the appropriate
// alias.
func substitutePath(index int, node exprNode, aliasList *aliasList) (string, error) {
	if index >= len(node.names) {
		return "", fmt.Errorf("substitutePath error: exprNode []names out of range")
	}
	str, err := aliasList.aliasPath(node.names[index])
	if err != nil {
		return "", err
	}
	return str, nil
}

// substituteValue will substitute the escaped character $v with the appropriate
// alias.
func substituteValue(index int, node exprNode, aliasList *aliasList) (string, error) {
	if index >= len(node.values) {
		return "", fmt.Errorf("substituteValue error: exprNode []values out of range")
	}
	str, err := aliasList.aliasValue(node.values[index])
	if err != nil {
		return "", err
	}
	return str, nil
}

// substituteChild will substitute the escaped character $c with the appropriate
// alias.
func substituteChild(index int, node exprNode, aliasList *aliasList) (string, error) {
	if index >= len(node.children) {
		return "", fmt.Errorf("substituteChild error: exprNode []children out of range")
	}
	return node.children[index].buildExpressionString(aliasList)
}

// aliasValue returns the corresponding alias to the dav value argument. Since
// values are not deduplicated as of now, all values are just appended to the
// aliasList and given the index as the alias.
func (al *aliasList) aliasValue(dav dynamodb.AttributeValue) (string, error) {
	if al == nil {
		return "", fmt.Errorf("aliasValue error: aliasList is nil")
	}

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
