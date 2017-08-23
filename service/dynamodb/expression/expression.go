package expression

import (
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ErrEmptyFactoryBuilder is an error that is returned if BuildFactory() is
// called on an empty FactoryBuilder.
var ErrEmptyFactoryBuilder = awserr.New("EmptyFactoryBuilder", "BuildFactory error: the argument FactoryBuilder is empty", nil)

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

func (list typeList) Len() int {
	return len(list)
}

func (list typeList) Less(i, j int) bool {
	return string(list[i]) < string(list[j])
}

func (list typeList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

// FactoryBuilder contains a list of structs fulfilling the TreeBuilder
// interface, representing the different types of Expressions that make up the
// total Expression as a whole. FactoryBuilder will have methods corresponding
// to different types of expressions (Update(), Condition(), Filter(), etc) that
// will allow users to add expressions to the input member Factory.
// FactoryBuilder will have a method Build() which will build the Factory struct
// which can be used to produce members of DynamoDB input structs.
//
// Example:
//
//     factoryBuilder := expression.KeyCondition(
//       expression.Key("someKey").Equal(expression.Value("someValue"))
//     ).Projection(
//       expression.NamesList("aName", "anotherName", "oneOtherName")
//     )
//     factory := factoryBuilder.BuildFactory()
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    factory.KeyCondition(),
//       ProjectionExpression:      factory.Projection(),
//       ExpressionAttributeNames:  factory.Names(),
//       ExpressionAttributeValues: factory.Values(),
//       TableName: aws.String("SomeTable"),
//     }
type FactoryBuilder struct {
	expressionMap map[expressionType]TreeBuilder
}

// BuildFactory builds a Factory struct with the same expressionMap as the
// argument FactoryBuilder. The aim of this method is to check the child
// TreeBuilders for their formats and return an error. This makes sure that any
// Factory has the right format or is an empty struct
func (factoryBuilder FactoryBuilder) BuildFactory() (Factory, error) {
	if factoryBuilder.expressionMap == nil {
		return Factory{}, ErrEmptyFactoryBuilder
	}

	factory := Factory{
		expressionMap: map[expressionType]TreeBuilder{},
	}
	for expressionType, treeBuilder := range factoryBuilder.expressionMap {
		_, err := treeBuilder.BuildTree()
		if err != nil {
			return Factory{}, err
		}
		factory.expressionMap[expressionType] = treeBuilder
	}

	return factory, nil
}

// Condition method will add the argument ConditionBuilder as a TreeBuilder to
// the argument FactoryBuilder. If the argument FactoryBuilder already has a
// ConditionBuilder, Condition() will overwrite the existing ConditionBuilder.
// Users will able to add other TreeBuilders to the FactoryBuilder or call
// BuildFactory() to build a Factory struct.
//
// Example:
//
//     // let factoryBuilder be an existing FactoryBuilder{} and
//     // conditionBuilder be an existing ConditionBuilder{}
//     factoryBuilder = factoryBuilder.Condition(conditionBuilder)
//
//     factory := factoryBuilder.BuildFactory()
func (factoryBuilder FactoryBuilder) Condition(conditionBuilder ConditionBuilder) FactoryBuilder {
	if factoryBuilder.expressionMap == nil {
		factoryBuilder.expressionMap = map[expressionType]TreeBuilder{}
	}
	factoryBuilder.expressionMap[condition] = conditionBuilder
	return factoryBuilder
}

// Condition function will create a FactoryBuilder with the argument
// conditionBuilder as a child TreeBuilder. Users will able to add other
// TreeBuilders to the FactoryBuilder or call BuildFactory() to build a Factory
// struct.
//
// Example:
//
//     // let conditionBuilder and projectionBuilder be an existing
//     // ConditionBuilder and ProjectionBuilder respectively.
//     factoryBuilder := expression.Condition(conditionBuilder)
//
//     factoryBuilder = factoryBuilder.Projection(projectionBuilder) // Adding a ProjectionBuilder
//     factory := factoryBuilder.BuildFactory()                      // Creating a Factory
func Condition(conditionBuilder ConditionBuilder) FactoryBuilder {
	ret := FactoryBuilder{}
	return ret.Condition(conditionBuilder)
}

// Projection method will add the argument ProjectionBuilder as a TreeBuilder to
// the argument FactoryBuilder. If the argument FactoryBuilder already has a
// ProjectionBuilder, Projection() will overwrite the existing ProjectionBuilder.
// Users will able to add other TreeBuilders to the FactoryBuilder or call
// BuildFactory() to build a Factory struct.
//
// Example:
//
//     // let factoryBuilder be an existing FactoryBuilder{} and
//     // projectionBuilder be an existing projectionBuilder{}
//     factoryBuilder = factoryBuilder.Projection(projectionBuilder)
//
//     factory := factoryBuilder.BuildFactory()
func (factoryBuilder FactoryBuilder) Projection(projectionBuilder ProjectionBuilder) FactoryBuilder {
	if factoryBuilder.expressionMap == nil {
		factoryBuilder.expressionMap = map[expressionType]TreeBuilder{}
	}
	factoryBuilder.expressionMap[projection] = projectionBuilder
	return factoryBuilder
}

// Projection function will create a FactoryBuilder with the argument
// projectionBuilder as a child TreeBuilder. Users will able to add other
// TreeBuilders to the FactoryBuilder or call BuildFactory() to build a Factory
// struct.
//
// Example:
//
//     // let projectionBuilder and conditionBuilder be an existing
//     // ProjectionBuilder and ConditionBuilder respectively.
//     factoryBuilder := expression.Projection(projectionBuilder)
//
//     factoryBuilder = factoryBuilder.Condition(conditionBuilder)   // Adding a ConditionBuilder
//     factory := factoryBuilder.BuildFactory()                      // Creating a Factory
func Projection(projectionBuilder ProjectionBuilder) FactoryBuilder {
	ret := FactoryBuilder{}
	return ret.Projection(projectionBuilder)
}

// // KeyCondition method will add the argument KeyConditionBuilder as a TreeBuilder to
// // the argument FactoryBuilder. If the argument FactoryBuilder already has a
// // KeyConditionBuilder, KeyCondition() will overwrite the existing KeyConditionBuilder.
// // Users will able to add other TreeBuilders to the FactoryBuilder or call
// // BuildFactory() to build a Factory struct.
// //
// // Example:
// //
// //     // let factoryBuilder be an existing FactoryBuilder{} and
// //     // keyConditionBuilder be an existing keyConditionBuilder{}
// //     factoryBuilder = factoryBuilder.KeyCondition(keyConditionBuilder)
// //
// //     factory := factoryBuilder.BuildFactory()
// func (factoryBuilder FactoryBuilder) KeyCondition(keyConditionBuilder KeyConditionBuilder) FactoryBuilder {
// 	if factoryBuilder.expressionMap == nil {
// 		factoryBuilder.expressionMap = map[expressionType]TreeBuilder{}
// 	}
// 	factoryBuilder.expressionMap[keyCondition] = keyConditionBuilder
// 	return factoryBuilder
// }
//
// // KeyCondition function will create a FactoryBuilder with the argument
// // keyConditionBuilder as a child TreeBuilder. Users will able to add other
// // TreeBuilders to the FactoryBuilder or call BuildFactory() to build a Factory
// // struct.
// //
// // Example:
// //
// //     // let keyConditionBuilder and conditionBuilder be an existing
// //     // KeyConditionBuilder and ConditionBuilder respectively.
// //     factoryBuilder := expression.KeyCondition(keyConditionBuilder)
// //
// //     factoryBuilder = factoryBuilder.Condition(conditionBuilder)   // Adding a ConditionBuilder
// //     factory := factoryBuilder.BuildFactory()                      // Creating a Factory
// func KeyCondition(keyConditionBuilder KeyConditionBuilder) FactoryBuilder {
// 	ret := FactoryBuilder{}
// 	return ret.KeyCondition(keyConditionBuilder)
// }

// Filter method will add the argument ConditionBuilder as a TreeBuilder to
// the argument FactoryBuilder. If the argument FactoryBuilder already has a
// ConditionBuilder, Filter() will overwrite the existing ConditionBuilder.
// Users will able to add other TreeBuilders to the FactoryBuilder or call
// BuildFactory() to build a Factory struct.
//
// Example:
//
//     // let factoryBuilder be an existing FactoryBuilder{} and
//     // filterBuilder be an existing filterBuilder{}
//     factoryBuilder = factoryBuilder.Filter(filterBuilder)
//
//     factory := factoryBuilder.BuildFactory()
func (factoryBuilder FactoryBuilder) Filter(filterBuilder ConditionBuilder) FactoryBuilder {
	if factoryBuilder.expressionMap == nil {
		factoryBuilder.expressionMap = map[expressionType]TreeBuilder{}
	}
	factoryBuilder.expressionMap[filter] = filterBuilder
	return factoryBuilder
}

// Filter function will create a FactoryBuilder with the argument
// filterBuilder as a child TreeBuilder. Users will able to add other
// TreeBuilders to the FactoryBuilder or call BuildFactory() to build a Factory
// struct.
//
// Example:
//
//     // let filterBuilder and conditionBuilder be an existing
//     // ConditionBuilder and ConditionBuilder respectively.
//     factoryBuilder := expression.Filter(filterBuilder)
//
//     factoryBuilder = factoryBuilder.Condition(conditionBuilder)   // Adding a ConditionBuilder
//     factory := factoryBuilder.BuildFactory()                      // Creating a Factory
func Filter(filterBuilder ConditionBuilder) FactoryBuilder {
	ret := FactoryBuilder{}
	return ret.Filter(filterBuilder)
}

// // Update method will add the argument UpdateBuilder as a TreeBuilder to
// // the argument FactoryBuilder. If the argument FactoryBuilder already has a
// // UpdateBuilder, Update() will overwrite the existing UpdateBuilder.
// // Users will able to add other TreeBuilders to the FactoryBuilder or call
// // BuildFactory() to build a Factory struct.
// //
// // Example:
// //
// //     // let factoryBuilder be an existing FactoryBuilder{} and
// //     // updateBuilder be an existing updateBuilder{}
// //     factoryBuilder = factoryBuilder.Update(updateBuilder)
// //
// //     factory := factoryBuilder.BuildFactory()
// func (factoryBuilder FactoryBuilder) Update(updateBuilder UpdateBuilder) FactoryBuilder {
// 	if factoryBuilder.expressionMap == nil {
// 		factoryBuilder.expressionMap = map[expressionType]TreeBuilder{}
// 	}
// 	factoryBuilder.expressionMap[update] = updateBuilder
// 	return factoryBuilder
// }
//
// // Update function will create a FactoryBuilder with the argument
// // updateBuilder as a child TreeBuilder. Users will able to add other
// // TreeBuilders to the FactoryBuilder or call BuildFactory() to build a Factory
// // struct.
// //
// // Example:
// //
// //     // let updateBuilder and conditionBuilder be an existing
// //     // UpdateBuilder and ConditionBuilder respectively.
// //     factoryBuilder := expression.Update(updateBuilder)
// //
// //     factoryBuilder = factoryBuilder.Condition(conditionBuilder)   // Adding a ConditionBuilder
// //     factory := factoryBuilder.BuildFactory()                      // Creating a Factory
// func Update(updateBuilder UpdateBuilder) FactoryBuilder {
// 	ret := FactoryBuilder{}
// 	return ret.Update(updateBuilder)
// }

// Factory will be a struct that will be able to generate members to DynamoDB
// inputs.
// The idea of FactoryBuilder and Factory is separated to be able to check the
// format of the child TreeBuilders and to return an error at the BuildFactory()
// step.
//
// Example:
//
//     factoryBuilder := expression.KeyCondition(
//       expression.Key("someKey").Equal(expression.Value("someValue"))
//     ).Projection(
//       expression.NamesList("aName", "anotherName", "oneOtherName")
//     )
//     factory := factoryBuilder.BuildFactory()
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    factory.KeyCondition(),
//       ProjectionExpression:      factory.Projection(),
//       ExpressionAttributeNames:  factory.Names(),
//       ExpressionAttributeValues: factory.Values(),
//       TableName: aws.String("SomeTable"),
//     }
type Factory struct {
	expressionMap map[expressionType]TreeBuilder
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

// Condition will return the *string corresponding to the Condition Expression
// of the argument Factory. This method is used to satisfy the members of
// DynamoDB input structs.
//
// Example:
//
//     // let factory be an instance of Factory{}
//
//     deleteInput := dynamodb.DeleteItemInput{
//       ConditionExpression:       factory.Condition(),
//       ExpressionAttributeNames:  factory.Names(),
//       ExpressionAttributeValues: factory.Values(),
//       Key: map[string]*dynamodb.AttributeValue{
//         "PartitionKey": &dynamodb.AttributeValue{
//           S: aws.String("SomeKey"),
//         },
//       },
//       TableName: aws.String("SomeTable"),
//     }
func (factory Factory) Condition() *string {
	return factory.returnExpression(condition)
}

// Filter will return the *string corresponding to the Filter Expression of the
// argument Factory. This method is used to satisfy the members of DynamoDB
// input structs.
//
// Example:
//
//     // let factory be an instance of Factory{}
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    factory.KeyCondition(),
//       FilterExpression:          factory.Filter(),
//       ExpressionAttributeNames:  factory.Names(),
//       ExpressionAttributeValues: factory.Values(),
//       TableName: aws.String("SomeTable"),
//     }
func (factory Factory) Filter() *string {
	return factory.returnExpression(filter)
}

// Projection will return the *string corresponding to the Projection Expression
// of the argument Factory. This method is used to satisfy the members of
// DynamoDB input structs.
//
// Example:
//
//     // let factory be an instance of Factory{}
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    factory.KeyCondition(),
//       ProjectionExpression:      factory.Projection(),
//       ExpressionAttributeNames:  factory.Names(),
//       ExpressionAttributeValues: factory.Values(),
//       TableName: aws.String("SomeTable"),
//     }
func (factory Factory) Projection() *string {
	return factory.returnExpression(projection)
}

// Names will return the map[string]*string corresponding to the
// ExpressionAttributeNames of the argument Factory. This method is used to
// satisfy the members of DynamoDB input structs.
//
// Example:
//
//     // let factory be an instance of Factory{}
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    factory.KeyCondition(),
//       ProjectionExpression:      factory.Projection(),
//       ExpressionAttributeNames:  factory.Names(),
//       ExpressionAttributeValues: factory.Values(),
//       TableName: aws.String("SomeTable"),
//     }
func (factory Factory) Names() map[string]*string {
	aliasList, _ := factory.buildChildTrees()

	namesMap := map[string]*string{}
	for ind, val := range aliasList.namesList {
		namesMap[fmt.Sprintf("#%v", ind)] = aws.String(val)
	}

	return namesMap
}

// Values will return the map[string]*dynamodb.AttributeValue corresponding to
// the ExpressionAttributeValues of the argument Factory. This method is used to
// satisfy the members of DynamoDB input structs.
//
// Example:
//
//     // let factory be an instance of Factory{}
//
//     queryInput := dynamodb.QueryInput{
//       KeyConditionExpression:    factory.KeyCondition(),
//       ProjectionExpression:      factory.Projection(),
//       ExpressionAttributeNames:  factory.Names(),
//       ExpressionAttributeValues: factory.Values(),
//       TableName: aws.String("SomeTable"),
//     }
func (factory Factory) Values() map[string]*dynamodb.AttributeValue {
	aliasList, _ := factory.buildChildTrees()

	valuesMap := map[string]*dynamodb.AttributeValue{}
	for i := 0; i < len(aliasList.valuesList); i++ {
		valuesMap[fmt.Sprintf(":%v", i)] = &aliasList.valuesList[i]
	}

	return valuesMap
}

// returnExpression will return *string corresponding to the type of Expression
// string specified by the expressionType. If there is no corresponding
// expression available in Factory, the method will return nil
func (factory Factory) returnExpression(expressionType expressionType) *string {
	if factory.expressionMap == nil {
		return nil
	}
	_, formattedExpressions := factory.buildChildTrees()

	return aws.String(formattedExpressions[expressionType])
}

// buildChildTrees will compile the list of ExpressionTreeBuilders that
// are the children of the argument Factory. The returned AliasList
// will represent all the alias tokens used in the expression strings. The
// returned map[string]string will map the type of expression (i.e. "condition",
// "update") to the appropriate expression string. buildChildTrees() assumes
// that the BuildTree() and BuildExpressionString() will not return an error
// because the error check should have been done at the BuildFactory() step.
func (factory Factory) buildChildTrees() (*AliasList, map[expressionType]string) {
	aliasList := &AliasList{}
	formattedExpressions := map[expressionType]string{}
	keys := typeList{}

	for expressionType := range factory.expressionMap {
		keys = append(keys, expressionType)
	}

	sort.Sort(keys)

	for _, key := range keys {
		exprNode, _ := factory.expressionMap[key].BuildTree()
		formattedExpression, _ := exprNode.BuildExpressionString(aliasList)
		formattedExpressions[key] = formattedExpression
	}

	return aliasList, formattedExpressions
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
