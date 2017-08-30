package expression

import (
	"fmt"
)

// keyConditionMode will specify the types of the struct KeyConditionBuilder,
// representing the different types of KeyConditions (i.e. And, Or, Between, ...)
type keyConditionMode int

const (
	// unsetKeyCond will catch errors if users make an empty KeyConditionBuilder
	unsetKeyCond keyConditionMode = iota
	// equalKeyCond will represent the Equals KeyCondition
	equalKeyCond
	// lessThanKeyCond will represent the Less Than KeyCondition
	lessThanKeyCond
	// lessThanEqualKeyCond will represent the Less Than Or Equal To KeyCondition
	lessThanEqualKeyCond
	// greaterThanKeyCond will represent the Greater Than KeyCondition
	greaterThanKeyCond
	// greaterThanEqualKeyCond will represent the Greater Than Or Equal To KeyCondition
	greaterThanEqualKeyCond
	// andKeyCond will represent the Logical And KeyCondition
	andKeyCond
	// betweenKeyCond will represent the Between KeyCondition
	betweenKeyCond
	// beginsWithKeyCond will represent the Begins With KeyCondition
	beginsWithKeyCond
)

// KeyConditionBuilder will represent Key Condition Expressions in DynamoDB. It
// is composed of operands (OperandBuilder) and other key conditions
// (KeyConditionBuilder). There are many different types of conditions,
// specified by keyConditionMode. KeyConditionBuilders will be the building
// blocks of Expressions.
// More Information at: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Query.html#Query.KeyConditionExpressions
type KeyConditionBuilder struct {
	operandList      []OperandBuilder
	keyConditionList []KeyConditionBuilder
	mode             keyConditionMode
}

// KeyEqual will create a KeyConditionBuilder with a KeyBuilder and a
// ValueBuilder as children. The KeyBuilder represents the key being compared
// and the ValueBuilder represents the value that the key is being compared to.
// The resulting KeyConditionBuilder can be used to build other
// KeyConditionBuilder or to create an Expression to be used in an operation
// input. This will be the function call.
//
// Example:
//
//     keyCondition := expression.KeyEqual(expression.Key("partitionKey"), expression.Value(someValue))
//
//     // Used to make another KeyConditionBuilder
//     anotherKeyCondition := keyCondition.And(expression.KeyEqual(expression.Key("sortKey"), expression.Value(someValue)))
//     // Used to make an Expression
//     expression := KeyCondition(keyCondition)
func KeyEqual(keyBuilder KeyBuilder, valueBuilder ValueBuilder) KeyConditionBuilder {
	return KeyConditionBuilder{
		operandList: []OperandBuilder{keyBuilder, valueBuilder},
		mode:        equalKeyCond,
	}
}

// Equal will create a KeyConditionBuilder if and only if the method is called
// on a KeyBuilder.
//
// Example:
//
//     // The following produces equivalent key conditions:
//     keyCondition := expression.KeyEqual(expression.Key("foo"), expression.Value(5))
//     keyCondition := expression.Key("foo").Equal(expression.Value(5))
func (kb KeyBuilder) Equal(valueBuilder ValueBuilder) KeyConditionBuilder {
	return KeyEqual(kb, valueBuilder)
}

// KeyLessThan will create a KeyConditionBuilder with a KeyBuilder and a
// ValueBuilder as children. The KeyBuilder represents the sort key being
// compared and the ValueBuilder represents the value that the sort key is being
// compared to. The resulting KeyConditionBuilder can be optionally provided in
// addition to an equal clause on the partition key to specify a subset of sort
// keys to retrieve. This will be the function call.
//
// Example:
//
//     partitionKeyCondition := expression.Key("partitionKey").Equal(expression.Value(someValue))
//     sortKeyCondition := expression.KeyLessThan(expression.Key("sortKey"), expression.Value(someValue))
//
//     expression := KeyCondition(partitionKeyCondition.And(sortKeyCondition))
func KeyLessThan(keyBuilder KeyBuilder, valueBuilder ValueBuilder) KeyConditionBuilder {
	return KeyConditionBuilder{
		operandList: []OperandBuilder{keyBuilder, valueBuilder},
		mode:        lessThanKeyCond,
	}
}

// LessThan will create a KeyConditionBuilder if and only if the method is
// called on a KeyBuilder.
//
// Example:
//
//     // The following produces equivalent key conditions:
//     keyCondition := expression.KeyLessThan(expression.Key("foo"), expression.Value(5))
//     keyCondition := expression.Key("foo").LessThan(expression.Value(5))
func (kb KeyBuilder) LessThan(valueBuilder ValueBuilder) KeyConditionBuilder {
	return KeyLessThan(kb, valueBuilder)
}

// KeyLessThanEqual will create a KeyConditionBuilder with a KeyBuilder and a
// ValueBuilder as children. The KeyBuilder represents the key being compared
// and the ValueBuilder represents the value that the key is being compared to.
// The resulting KeyConditionBuilder can be optionally provided in addition to
// an equal clause on the partition key to specify a subset of sort keys to
// retrieve. This will be the function call.
//
// Example:
//
//     partitionKeyCondition := expression.Key("partitionKey").Equal(expression.Value(someValue))
//     sortKeyCondition := expression.KeyLessThanEqual(expression.Key("sortKey"), expression.Value(someValue))
//
//     expression := KeyCondition(partitionKeyCondition.And(sortKeyCondition))
func KeyLessThanEqual(keyBuilder KeyBuilder, valueBuilder ValueBuilder) KeyConditionBuilder {
	return KeyConditionBuilder{
		operandList: []OperandBuilder{keyBuilder, valueBuilder},
		mode:        lessThanEqualKeyCond,
	}
}

// LessThanEqual will create a KeyConditionBuilder if and only if the method is
// called on a KeyBuilder.
//
// Example:
//
//     // The following produces equivalent key conditions:
//     keyCondition := expression.KeyLessThanEqual(expression.Key("foo"), expression.Value(5))
//     keyCondition := expression.Key("foo").LessThanEqual(expression.Value(5))
func (kb KeyBuilder) LessThanEqual(valueBuilder ValueBuilder) KeyConditionBuilder {
	return KeyLessThanEqual(kb, valueBuilder)
}

// KeyGreaterThan will create a KeyConditionBuilder with a KeyBuilder and a
// ValueBuilder as children. The KeyBuilder represents the key being compared
// and the ValueBuilder represents the value that the key is being compared to.
// The resulting KeyConditionBuilder can be optionally provided in addition to
// an equal clause on the partition key to specify a subset of sort keys to
// retrieve. This will be the function call.
//
// Example:
//
//     partitionKeyCondition := expression.Key("partitionKey").Equal(expression.Value(someValue))
//     sortKeyCondition := expression.KeyGreaterThan(expression.Key("sortKey"), expression.Value(someValue))
//
//     expression := KeyCondition(partitionKeyCondition.And(sortKeyCondition))
func KeyGreaterThan(keyBuilder KeyBuilder, valueBuilder ValueBuilder) KeyConditionBuilder {
	return KeyConditionBuilder{
		operandList: []OperandBuilder{keyBuilder, valueBuilder},
		mode:        greaterThanKeyCond,
	}
}

// GreaterThan will create a KeyConditionBuilder if and only if the method
// is called on a KeyBuilder.
//
// Example:
//
//     // The following produces equivalent key conditions:
//     keyCondition := expression.KeyGreaterThan(expression.Key("foo"), expression.Value(5))
//     keyCondition := expression.Key("foo").GreaterThan(expression.Value(5))
func (kb KeyBuilder) GreaterThan(valueBuilder ValueBuilder) KeyConditionBuilder {
	return KeyGreaterThan(kb, valueBuilder)
}

// KeyGreaterThanEqual will create a KeyConditionBuilder with a KeyBuilder and a
// ValueBuilder as children. The KeyBuilder represents the key being compared and
// the ValueBuilder represents the value that the key is being compared to.
// The resulting KeyConditionBuilder can be optionally provided in addition to
// an equal clause on the partition key to specify a subset of sort keys to
// retrieve. This will be the function call.
//
// Example:
//
//     partitionKeyCondition := expression.Key("partitionKey").Equal(expression.Value(someValue))
//     sortKeyCondition := expression.KeyGreaterThanEqual(expression.Key("sortKey"), expression.Value(someValue))
//
//     expression := KeyCondition(partitionKeyCondition.And(sortKeyCondition))
func KeyGreaterThanEqual(keyBuilder KeyBuilder, valueBuilder ValueBuilder) KeyConditionBuilder {
	return KeyConditionBuilder{
		operandList: []OperandBuilder{keyBuilder, valueBuilder},
		mode:        greaterThanEqualKeyCond,
	}
}

// GreaterThanEqual will create a KeyConditionBuilder if and only if the method
// is called on a KeyBuilder.
//
// Example:
//
//     // The following produces equivalent key conditions:
//     keyCondition := expression.KeyGreaterThanEqual(expression.Key("foo"), expression.Value(5))
//     keyCondition := expression.Key("foo").GreaterThanEqual(expression.Value(5))
func (kb KeyBuilder) GreaterThanEqual(valueBuilder ValueBuilder) KeyConditionBuilder {
	return KeyGreaterThanEqual(kb, valueBuilder)
}

// KeyAnd will create a KeyConditionBuilder with two KeyConditionBuilders as
// children. The first KeyConditionBuilder must be an equalKeyCond
// KeyConditionBuilder, representing the equal clause on the partition key of
// an item. The second KeyConditionBuilder must not be another andKeyCond
// KeyConditionBuilder since Key Conditions only allow the first partition key
// equal clause to be optionally extended with a rule on the sort key.
// The resulting KeyConditionBuilder is used to create a Expression.
// This will be the function call.
//
// Example:
//
//     partitionKeyCondition := expression.Key("partitionKey").Equal(expression.Value(someValue))
//     sortKeyCondition := expression.KeyGreaterThanEqual(expression.Key("sortKey"), expression.Value(someValue))
//     andKeyCondition := expression.KeyAnd(partitionKeyCondition, sortKeyCondition)
//
//     expression := KeyCondition(andKeyCondition)
func KeyAnd(left, right KeyConditionBuilder) KeyConditionBuilder {
	if left.mode != equalKeyCond {
		return KeyConditionBuilder{
			mode: andKeyCond,
		}
	}
	if right.mode == andKeyCond {
		return KeyConditionBuilder{
			mode: andKeyCond,
		}
	}
	return KeyConditionBuilder{
		keyConditionList: []KeyConditionBuilder{left, right},
		mode:             andKeyCond,
	}
}

// And will create a KeyConditionBuilder if and only if the method is called on
// a KeyConditionBuilder.
//
// Example:
//
//     // The following produces equivalent key conditions:
//     keyCondition := expression.KeyAnd(partitionKeyCondition, sortKeyCondition)
//     keyCondition := partitionKeyCondition.And(sortKeyCondition)
func (kcb KeyConditionBuilder) And(right KeyConditionBuilder) KeyConditionBuilder {
	return KeyAnd(kcb, right)
}

// KeyBetween will create a KeyConditionBuilder with three operands as children, the
// first operand representing the sort key being compared, the second operand
// representing the lower bound value of the first operand, and the third
// operand representing the upper bound value of the first operand. The
// resulting KeyConditionBuilder is used narrow down possible values of sort keys
// in conjunction with an equal clause on the partition key. This will be the
// function call.
//
// Example:
//
//     partitionKeyCondition := expression.Key("partitionKey").Equal(expression.Value(someValue))
//     sortKeyCondition := expression.KeyBetween(expression.Key("sortKey"), expression.Value(lower), expression.Value(upper))
//
//     expression := KeyCondition(partitionKeyCondition.And(sortKeyCondition))
func KeyBetween(keyBuilder KeyBuilder, lower, upper ValueBuilder) KeyConditionBuilder {
	return KeyConditionBuilder{
		operandList: []OperandBuilder{keyBuilder, lower, upper},
		mode:        betweenKeyCond,
	}
}

// Between will create a KeyConditionBuilder if and only if the method is called
// on KeyBuilder.
//
// Example:
//
//     // The following produces equivalent key conditions:
//     keyCondition := expression.KeyBetween(expression.Key("sortKey"), expression.Value(lower), expression.Value(upper))
//     keyCondition := expression.Key("sortKey").Between(expression.Value(lower), expression.Value(upper))
func (kb KeyBuilder) Between(lower, upper ValueBuilder) KeyConditionBuilder {
	return KeyBetween(kb, lower, upper)
}

// KeyBeginsWith will create a KeyConditionBuilder with a key and a value as
// children. The key will represent the sortKey of the item being compared. The
// value will represent the prefix in which the sortKey will be compared
// with. The function will return true if the sortKey starts with the prefix.
// The resulting KeyConditionBuilder can be optionally provided in addition to
// an equal clause on the partition key to specify a subset of sort keys to
// retrieve. This will be the function call.
//
// Example:
//
//     partitionKeyCondition := expression.Key("partitionKey").Equal(expression.Value(someValue))
//     sortKeyCondition := expression.KeyBeginsWith(expression.Key("sortKey"), "prefix")
//
//     expression := KeyCondition(partitionKeyCondition.And(sortKeyCondition)) // Used to make a Expression
func KeyBeginsWith(keyBuilder KeyBuilder, prefix string) KeyConditionBuilder {
	valueBuilder := ValueBuilder{
		value: prefix,
	}
	return KeyConditionBuilder{
		operandList: []OperandBuilder{keyBuilder, valueBuilder},
		mode:        beginsWithKeyCond,
	}
}

// BeginsWith will create a KeyConditionBuilder if and only if the method is called
// on KeyBuilder.
//
// Example:
//
//     // The following produces equivalent key conditions:
//     keyCondition := expression.KeyBeginsWith(expression.Key("sortKey"), "prefix")
//     keyCondition := expression.Key("sortKey").BeginsWith("prefix")
func (kb KeyBuilder) BeginsWith(prefix string) KeyConditionBuilder {
	return KeyBeginsWith(kb, prefix)
}

// buildTree will build a tree structure of exprNodes based on the tree
// structure of the input KeyConditionBuilder's child KeyConditions/Operands.
// buildTree() satisfies the treeBuilder interface so KeyConditionBuilder can be
// a part of Expression struct.
func (kcb KeyConditionBuilder) buildTree() (exprNode, error) {
	childNodes, err := kcb.buildChildNodes()
	if err != nil {
		return exprNode{}, err
	}
	ret := exprNode{
		children: childNodes,
	}

	switch kcb.mode {
	case equalKeyCond, lessThanKeyCond, lessThanEqualKeyCond, greaterThanKeyCond, greaterThanEqualKeyCond:
		return compareBuildKeyCondition(kcb.mode, ret)
	case andKeyCond:
		return andBuildKeyCondition(kcb, ret)
	case betweenKeyCond:
		return betweenBuildKeyCondition(ret)
	case beginsWithKeyCond:
		return beginsWithBuildKeyCondition(ret)
	case unsetKeyCond:
		return exprNode{}, newUnsetParameterError("buildTree", "KeyConditionBuilder")
	default:
		return exprNode{}, fmt.Errorf("buildKeyCondition error: unsupported mode: %v", kcb.mode)
	}
}

// compareBuildKeyCondition is the function to make exprNodes from Compare
// KeyConditionBuilders. compareBuildKeyCondition will only be called by the
// buildKeyCondition method. This function assumes that the argument
// KeyConditionBuilder has the right format.
func compareBuildKeyCondition(keyConditionMode keyConditionMode, node exprNode) (exprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	switch keyConditionMode {
	case equalKeyCond:
		node.fmtExpr = "$c = $c"
	case lessThanKeyCond:
		node.fmtExpr = "$c < $c"
	case lessThanEqualKeyCond:
		node.fmtExpr = "$c <= $c"
	case greaterThanKeyCond:
		node.fmtExpr = "$c > $c"
	case greaterThanEqualKeyCond:
		node.fmtExpr = "$c >= $c"
	default:
		return exprNode{}, fmt.Errorf("build compare key condition error: unsupported mode: %v", keyConditionMode)
	}

	return node, nil
}

// andBuildKeyCondition is the function to make exprNodes from And
// KeyConditionBuilders. andBuildKeyCondition will only be called by the
// buildKeyCondition method. This function assumes that the argument
// KeyConditionBuilder has the right format.
func andBuildKeyCondition(keyConditionBuilder KeyConditionBuilder, node exprNode) (exprNode, error) {
	if len(keyConditionBuilder.keyConditionList) == 0 && len(keyConditionBuilder.operandList) == 0 {
		return exprNode{}, newInvalidParameterError("andBuildKeyCondition", "KeyConditionBuilder")
	}
	// create a string with escaped characters to substitute them with proper
	// aliases during runtime
	node.fmtExpr = "($c) AND ($c)"

	return node, nil
}

// betweenBuildKeyCondition is the function to make exprNodes from Between
// KeyConditionBuilders. betweenBuildKeyCondition will only be called by the
// buildKeyCondition method. This function assumes that the argument
// KeyConditionBuilder has the right format.
func betweenBuildKeyCondition(node exprNode) (exprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	node.fmtExpr = "$c BETWEEN $c AND $c"

	return node, nil
}

// beginsWithBuildKeyCondition is the function to make exprNodes from
// BeginsWith KeyConditionBuilders. beginsWithBuildKeyCondition will only be
// called by the buildKeyCondition method. This function assumes that the argument
// KeyConditionBuilder has the right format.
func beginsWithBuildKeyCondition(node exprNode) (exprNode, error) {
	// Create a string with special characters that can be substituted later: $c
	node.fmtExpr = "begins_with ($c, $c)"

	return node, nil
}

// buildChildNodes will create the list of the child exprNodes. This avoids
// duplication of code amongst the various buildConditions.
func (kcb KeyConditionBuilder) buildChildNodes() ([]exprNode, error) {
	childNodes := make([]exprNode, 0, len(kcb.keyConditionList)+len(kcb.operandList))
	for _, keyCondition := range kcb.keyConditionList {
		node, err := keyCondition.buildTree()
		if err != nil {
			return []exprNode{}, err
		}
		childNodes = append(childNodes, node)
	}
	for _, operand := range kcb.operandList {
		ope, err := operand.BuildOperand()
		if err != nil {
			return []exprNode{}, err
		}
		childNodes = append(childNodes, ope.exprNode)
	}

	return childNodes, nil
}
