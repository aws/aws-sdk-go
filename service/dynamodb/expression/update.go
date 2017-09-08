package expression

import (
	"fmt"
	"sort"
	"strings"
)

// operationMode will specify the types of update operations that the
// updateBuilder is going to represent. The const is in a string to use the
// const value as a map key and as a string when creating the formatted
// expression for the exprNodes.
type operationMode string

const (
	setOperation    operationMode = "SET"
	removeOperation               = "REMOVE"
	addOperation                  = "ADD"
	deleteOperation               = "DELETE"
)

// Implementing the Sort interface
type modeList []operationMode

func (ml modeList) Len() int {
	return len(ml)
}

func (ml modeList) Less(i, j int) bool {
	return string(ml[i]) < string(ml[j])
}

func (ml modeList) Swap(i, j int) {
	ml[i], ml[j] = ml[j], ml[i]
}

// UpdateBuilder will represent Update Expressions in DynamoDB. The
// operationList represents the different update operations that are available
// to DynamoDB. Each slice of operationBuilders will represent an update action,
// specified by the map key. Each operationBuilder in the slice will represent
// the specific update action. UpdateBuilder will be a building block of the
// Builder struct.
// More Information at: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.UpdateExpressions.html
type UpdateBuilder struct {
	operationList map[operationMode][]operationBuilder
}

// operationBuilder will represent specific update actions (SET, REMOVE, ADD,
// DELETE). The mode will specify what type of update action the
// operationBuilder represents.
type operationBuilder struct {
	name  NameBuilder
	value OperandBuilder
	mode  operationMode
}

// buildOperation builds an exprNode from an operationBuilder. buildOperation
// will be called recursively by buildTree in order to create a tree structure
// of exprNodes representing the parent/child relationships between
// UpdateBuilders and operationBuilders.
func (ob operationBuilder) buildOperation() (exprNode, error) {
	pathChild, err := ob.name.BuildOperand()
	if err != nil {
		return exprNode{}, err
	}

	node := exprNode{
		children: []exprNode{pathChild.exprNode},
		fmtExpr:  "$c",
	}

	if ob.mode == removeOperation {
		return node, nil
	}

	valueChild, err := ob.value.BuildOperand()
	if err != nil {
		return exprNode{}, err
	}
	node.children = append(node.children, valueChild.exprNode)

	switch ob.mode {
	case setOperation:
		node.fmtExpr += " = $c"
	case addOperation, deleteOperation:
		node.fmtExpr += " $c"
	default:
		return exprNode{}, fmt.Errorf("build update error: build operation error: unsupported mode: %v", ob.mode)
	}

	return node, nil
}

// Delete as a function call will create an empty UpdateBuilder and call the
// Delete method call.
// The DELETE action only supports DynamoDB Set types. The argument name must be
// a name to an item attribute of the Set type. The argument value must be a Set
// type representing the subset of the set described by the name that is going
// to be deleted.
// The resulting UpdateBuilder can be used to create a Builder to be used
// as a part of an Update Input. Users can also chain other update methods on
// the UpdateBuilder.
//
// Example:
//
//     update := expression.Delete(expression.Name("pathToList"), expression.Value("subsetToDelete"))
//
//     // Adding more update methods
//     anotherUpdate := update.Remove(expression.Name("someName"))
//     // Creating a Builder
//     builder := Update(update)
func Delete(name NameBuilder, value ValueBuilder) UpdateBuilder {
	emptyUpdateBuilder := UpdateBuilder{}
	return emptyUpdateBuilder.Delete(name, value)
}

// Delete as a method call will add a delete action to the argument
// UpdateBuilder in the form of adding an operationBuilder.
// The DELETE action only supports DynamoDB Set types. The argument name must be
// a name to an item attribute of the Set type. The argument value must be a Set
// type representing the subset of the set described by the name that is going
// to be deleted.
//
// Example:
//
//     // newUpdate sets foo to 15 and deletes 5 from barList
//     oldUpdate := expression.Set(expression.Name("foo"), expression.Value(15))
//     newUpdate := oldUpdate.Delete(expression.Name("barList"), expression.Value([]int{5}))
func (ub UpdateBuilder) Delete(name NameBuilder, value ValueBuilder) UpdateBuilder {
	if ub.operationList == nil {
		ub.operationList = map[operationMode][]operationBuilder{}
	}
	ub.operationList[deleteOperation] = append(ub.operationList[deleteOperation], operationBuilder{
		name:  name,
		value: value,
		mode:  deleteOperation,
	})
	return ub
}

// Add as a function call will create an empty UpdateBuilder and call the Add
// method call.
// The ADD action only supports the DynamoDB Set and Number types. The argument
// value must be a set type or a number type. The argument name must specify the
// item attribute that the value will be added to. If there is no existing item
// attribute, the ADD action will create a new item attribute with the value
// specified. (Similar to SET) If there is an existing item, the types of the
// item attribute and value must match. If both are Numbers, the ADD action will
// numerically add the value to the item attribute. If both are Sets, the ADD
// action will append the value to the item attribute.
// The resulting UpdateBuilder can be used to create a Builder to be used
// as a part of an Update Input. Users can also chain other update methods on
// the UpdateBuilder.
//
// Example:
//
//     update := expression.Add(expression.Name("pathToList"), expression.Value("valueToAdd"))
//
//     // Adding more update methods
//     anotherUpdate := update.Remove(expression.Name("someName"))
//     // Creating a Builder
//     builder := Update(update)
func Add(name NameBuilder, value ValueBuilder) UpdateBuilder {
	emptyUpdateBuilder := UpdateBuilder{}
	return emptyUpdateBuilder.Add(name, value)
}

// Add as a method call will add an add action to the argument UpdateBuilder by
// adding an operationBuilder.
// The ADD action only supports the DynamoDB Set and Number types. The argument
// value must be a set type or a number type. The argument name must specify the
// item attribute that the value will be added to. If there is no existing item
// attribute, the ADD action will create a new item attribute with the value
// specified. (Similar to SET) If there is an existing item, the types of the
// item attribute and value must match. If both are Numbers, the ADD action will
// numerically add the value to the item attribute. If both are Sets, the ADD
// action will append the value to the item attribute.
//
// Example:
//
//     // newUpdate sets foo to 15 and adds 5 to barList
//     oldUpdate := expression.Set(expression.Name("foo"), expression.Value(15))
//     newUpdate := oldUpdate.Add(expression.Name("barList"), expression.Value([]int{5}))
func (ub UpdateBuilder) Add(name NameBuilder, value ValueBuilder) UpdateBuilder {
	if ub.operationList == nil {
		ub.operationList = map[operationMode][]operationBuilder{}
	}
	ub.operationList[addOperation] = append(ub.operationList[addOperation], operationBuilder{
		name:  name,
		value: value,
		mode:  addOperation,
	})
	return ub
}

// Remove as a function call will create an empty UpdateBuilder and call the
// Remove method call.
// The argument name must represent a name to the item attribute that will be
// removed.
// The resulting UpdateBuilder can be used to create a Builder to be used as
// a part of an Update Input. Users can also chain other update methods on the
// UpdateBuilder. This will be the function call.
//
// Example:
//
//     update := expression.Remove(expression.Name("pathToItem"))
//
//     // Adding more update methods
//     anotherUpdate := update.Remove(expression.Name("someName"))
//     // Creating a Builder
//     builder := Update(update)
func Remove(name NameBuilder) UpdateBuilder {
	emptyUpdateBuilder := UpdateBuilder{}
	return emptyUpdateBuilder.Remove(name)
}

// Remove as a method call will add an remove action to the argument
// UpdateBuilder by adding an operationBuilder.
// The argument name must represent a name to the item attribute that will be
// removed.
//
// Example:
//
//     // newUpdate sets foo to 15 and removes attribute bar from the item
//     oldUpdate := expression.Set(expression.Name("foo"), expression.Value(15))
//     newUpdate := oldUpdate.Remove(expression.Name("bar"))
func (ub UpdateBuilder) Remove(name NameBuilder) UpdateBuilder {
	if ub.operationList == nil {
		ub.operationList = map[operationMode][]operationBuilder{}
	}
	ub.operationList[removeOperation] = append(ub.operationList[removeOperation], operationBuilder{
		name: name,
		mode: removeOperation,
	})
	return ub
}

// Set as a function call will create an empty UpdateBuilder and call the
// Set method call.
// The argument name will represent the name to the item attribute that is being
// modified. If an item attribute already exists at the specified name, it will
// be overwritten unless otherwise specified (See IfNotExists()). If there are
// no item attributes at the specified name, a new item attribute will be
// created. setValue will represent the value that the item attribute will be
// set to.
// The setValue can be any of the following:
//     NameBuilder
//     ValueBuilder
//     MinusBuilder
//     PlusBuilder
//     ListAppendBuilder
//     IfNotExistsBuilder
// The resulting UpdateBuilder can be used to create a Builder to be used as a
// part of an Update Input. Users can also chain other update methods on the
// UpdateBuilder. This will be the function call.
//
// Example:
//
//     update := expression.Set(expression.Name("pathToItem"), expression.Value("item"))
//
//     // Adding more update methods
//     anotherUpdate := update.Remove(expression.Name("someName"))
//     // Creating a Builder
//     builder := Update(update)
func Set(name NameBuilder, operandBuilder OperandBuilder) UpdateBuilder {
	emptyUpdateBuilder := UpdateBuilder{}
	return emptyUpdateBuilder.Set(name, operandBuilder)
}

// Set as a method call will add an set action to the argument UpdateBuilder by
// adding an operationBuilder.
// The argument name will represent the name to the item attribute that is being
// modified. If an item attribute already exists at the specified name, it will
// be overwritten unless otherwise specified (See IfNotExists()). If there are
// no item attributes at the specified name, a new item attribute will be
// created. setValue will represent the value that the item attribute will be
// set to.
// The setValue can be any of the following:
//     NameBuilder
//     ValueBuilder
//     MinusBuilder
//     PlusBuilder
//     ListAppendBuilder
//     IfNotExistsBuilder
//
// Example:
//
//     // newUpdate sets foo to 15 and also sets bar to "baz"
//     oldUpdate := expression.Set(expression.Name("foo"), expression.Value(15))
//     newUpdate := oldUpdate.Set(expression.Name("bar"), expression.Name("baz"))
func (ub UpdateBuilder) Set(name NameBuilder, operandBuilder OperandBuilder) UpdateBuilder {
	if ub.operationList == nil {
		ub.operationList = map[operationMode][]operationBuilder{}
	}
	ub.operationList[setOperation] = append(ub.operationList[setOperation], operationBuilder{
		name:  name,
		value: operandBuilder,
		mode:  setOperation,
	})
	return ub
}

// buildTree will build a tree structure of exprNodes based on the tree
// structure of the input UpdateBuilder's child UpdateBuilders/Operands.
// buildTree() satisfies the TreeBuilder interface so ProjectionBuilder can be a
// part of Expression struct.
func (ub UpdateBuilder) buildTree() (exprNode, error) {
	if ub.operationList == nil {
		return exprNode{}, newUnsetParameterError("buildTree", "UpdateBuilder")
	}
	ret := exprNode{
		children: []exprNode{},
	}

	modes := modeList{}

	for mode := range ub.operationList {
		modes = append(modes, mode)
	}

	sort.Sort(modes)

	for _, key := range modes {
		ret.fmtExpr += string(key) + " $c\n"

		childNode, err := buildChildNodes(ub.operationList[key])
		if err != nil {
			return exprNode{}, err
		}

		ret.children = append(ret.children, childNode)
	}

	return ret, nil
}

// buildChildNodes will create the list of the child exprNodes.
func buildChildNodes(operationBuilderList []operationBuilder) (exprNode, error) {
	if len(operationBuilderList) == 0 {
		return exprNode{}, fmt.Errorf("buildChildNodes error: operationBuilder list is empty")
	}

	node := exprNode{
		children: make([]exprNode, 0, len(operationBuilderList)),
		fmtExpr:  "$c" + strings.Repeat(", $c", len(operationBuilderList)-1),
	}

	for _, val := range operationBuilderList {
		valNode, err := val.buildOperation()
		if err != nil {
			return exprNode{}, err
		}
		node.children = append(node.children, valNode)
	}

	return node, nil
}
