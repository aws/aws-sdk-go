package expression

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// ValueBuilder represents a value operand and will implement the OperandBuilder
// interface. It will have various methods corresponding to the operations
// supported by DynamoDB operations. (i.e. AND, BETWEEN, EQUALS) The underlying
// undefined type member variable will be converted into dynamodb.AttributeValue
// by dynamodbattribute.Marshal(in) when Builder are created.
type ValueBuilder struct {
	value interface{}
}

// NameBuilder represents a name of a top level item attribute or a nested
// attribute. Since NameBuilder represents a DynamoDB Operand, it will implement
// the OperandBuilder interface. It will have various methods corresponding to
// the operations supported by DynamoDB operations. (i.e. AND, BETWEEN, EQUALS)
type NameBuilder struct {
	name string
}

// SizeBuilder represents the output of the function size ("someName"), which
// evaluates to the size of the item attribute defined by "someName".
// SizeBuilder will implement OperandBuilder interface. It will have various
// methods corresponding to the operations supported by DynamoDB operations.
// (i.e. AND, BETWEEN, EQUALS)
type SizeBuilder struct {
	nameBuilder NameBuilder
}

// Operand is a wrapper around the exprNode struct. BuildOperand returns Operand
// instead of exprNode so that the underlying structure of exprNodes is not
// accessible to external users.
type Operand struct {
	exprNode exprNode
}

// KeyBuilder represents either the partition key or the sort key, both of which
// are top level attributes to some item in DynamoDB. KeyBuilder have a member
// named key, which is a string that identifies the key. Keybuilder will
// implement OperandBuilder interface. It will have various methods
// corresponding to the operations supported by DynamoDB operations.
// (i.e. AND, BETWEEN, EQUALS) KeyBuilder will only be used in the context of
// Key Condition Expressions (KeyConditionBuilder)
type KeyBuilder struct {
	key string
}

// setValueMode will specify the type of SetValueBuilder. The default value will
// be unsetValue so if a User were to create an empty SetValueBuilder, we can
// return ErrUnsetSetValue when BuildOperand() is called.
type setValueMode int

const (
	unsetValue setValueMode = iota
	plus
	minus
	listAppend
	ifNotExists
)

// SetValueBuilder represents the result of the Plus() function/method for DynamoDB
// Update SET operations. The members represent the two values that will be
// added. The PlusBuilder will only be used as an argument in the Set() function
type SetValueBuilder struct {
	leftOperand  OperandBuilder
	rightOperand OperandBuilder
	mode         setValueMode
}

// OperandBuilder represents the idea of Operand which are building blocks to
// DynamoDB Expressions. OperandBuilders will be children of ConditionBuilders
// to represent a tree like structure of Expression dependencies. The method
// BuildOperand() will create an instance of Operand with a child exprNode,
// which is an generic representation of both Operands and Conditions.
// BuildOperand() will mainly be called recursively by the buildTree() method
// call.
type OperandBuilder interface {
	BuildOperand() (Operand, error)
}

// Name creates a NameBuilder, which implements the OperandBuilder interface.
//
// Example:
//
//     condition := Name("foo").Equal(Name("bar"))
func Name(name string) NameBuilder {
	return NameBuilder{
		name: name,
	}
}

// Value creates a ValueBuilder, which implements the OperandBuilder interface.
//
// Example:
//
//     condition := Name("foo").Equal(Value(10))
func Value(value interface{}) ValueBuilder {
	return ValueBuilder{
		value: value,
	}
}

// Size creates a SizeBuilder which implements the OperandBuilder interface.
// Size will only be called in a pattern in order to create ConditionBuilders.
//
// Example:
//
//     condition := Name("foo").Size().Equal(Value(10))
func (nb NameBuilder) Size() SizeBuilder {
	return SizeBuilder{
		nameBuilder: nb,
	}
}

// Key creates a KeyBuilder which implements the OperandBuilder interface.
// Key will only be called in a pattern in order to create KeyConditionBuilders.
//
// Example:
//
//     keyCondition := expression.Key("foo").Equal(expression.Value("bar"))
func Key(key string) KeyBuilder {
	return KeyBuilder{
		key: key,
	}
}

// Plus will create a SetValueBuilder to be used in as an argument to Set().
// The arguments can either be NameBuilders or ValueBuilders. Plus() only
// supports DynamoDB Number types, so the ValueBuilder must be a Number and the
// NameBuilder must specify an item attribute of type Number.
// More information: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.UpdateExpressions.html#Expressions.UpdateExpressions.SET.IncrementAndDecrement
//
// Example:
//
//     update, err := expression.Set(expression.Name("someName"), expression.Plus(expression.Value(5), expression.Value(10)))
func Plus(leftOperand, rightOperand OperandBuilder) SetValueBuilder {
	return SetValueBuilder{
		leftOperand:  leftOperand,
		rightOperand: rightOperand,
		mode:         plus,
	}
}

// Plus will create a SetValueBuilder. This will be the method call for
// NameBuilders
//
// Example:
//
//     // The following produces equivalent SetValueBuilders
//     SetValue := expression.Plus(expression.Name("someName"), expression.Value(5))
//     SetValue := expression.Name("someName").Plus(expression.Value(5))
func (nb NameBuilder) Plus(rightOperand OperandBuilder) SetValueBuilder {
	return Plus(nb, rightOperand)
}

// Plus will create a SetValueBuilder. This will be the method call for
// ValueBuilders
//
// Example:
//
//     // The following produces equivalent SetValueBuilders
//     SetValue := expression.Plus(expression.Value(10), expression.Value(5))
//     SetValue := expression.Value(10).Plus(expression.Value(5))
func (vb ValueBuilder) Plus(rightOperand OperandBuilder) SetValueBuilder {
	return Plus(vb, rightOperand)
}

// Minus will create a SetValueBuilder to be used in as an argument to Set().
// The arguments can either be NameBuilders or ValueBuilders. Minus() only
// supports DynamoDB Number types, so the ValueBuilder must be a Number and the
// NameBuilder must specify an item attribute of type Number.
// More information: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.UpdateExpressions.html#Expressions.UpdateExpressions.SET.IncrementAndDecrement
//
// Example:
//
//     update, err := expression.Set(expression.Name("someName"), expression.Minus(expression.Value(5), expression.Value(10)))
func Minus(leftOperand, rightOperand OperandBuilder) SetValueBuilder {
	return SetValueBuilder{
		leftOperand:  leftOperand,
		rightOperand: rightOperand,
		mode:         minus,
	}
}

// Minus will create a SetValueBuilder. This will be the method call for
// NameBuilders
//
// Example:
//
//     // The following produces equivalent SetValueBuilders
//     SetValue := expression.Minus(expression.Name("someName"), expression.Value(5))
//     SetValue := expression.Name("someName").Minus(expression.Value(5))
func (nb NameBuilder) Minus(rightOperand OperandBuilder) SetValueBuilder {
	return Minus(nb, rightOperand)
}

// Minus will create a SetValueBuilder. This will be the method call for
// ValueBuilders
//
// Example:
//
//     // The following produces equivalent SetValueBuilders
//     SetValue := expression.Minus(expression.Value(10), expression.Value(5))
//     SetValue := expression.Value(10).Minus(expression.Value(5))
func (vb ValueBuilder) Minus(rightOperand OperandBuilder) SetValueBuilder {
	return Minus(vb, rightOperand)
}

// ListAppend will create a SetValueBuilder to be used in as an argument to
// Set(). The arguments can either be NameBuilders or ValueBuilders.
// ListAppend() only supports DynamoDB List types, so the ValueBuilder must be a
// List and the NameBuilder must specify an item attribute of type List.
// More information: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.UpdateExpressions.html#Expressions.UpdateExpressions.SET.UpdatingListElements
//
// Example:
//
//     update, err := expression.Set(expression.Name("someName"), expression.ListAppend(expression.Name("nameToList"), expression.Value([]string{"some", "list"})))
func ListAppend(leftOperand, rightOperand OperandBuilder) SetValueBuilder {
	return SetValueBuilder{
		leftOperand:  leftOperand,
		rightOperand: rightOperand,
		mode:         listAppend,
	}
}

// ListAppend will create a SetValueBuilder. This will be the method call for
// NameBuilders
//
// Example:
//
//     // The following produces equivalent SetValueBuilders
//     SetValue := expression.ListAppend(expression.Name("nameToList"), expression.Value([]string{"some", "list"}))
//     SetValue := expression.Name("nameToList").ListAppend(expression.Value([]string{"some", "list"}))
func (nb NameBuilder) ListAppend(rightOperand OperandBuilder) SetValueBuilder {
	return ListAppend(nb, rightOperand)
}

// ListAppend will create a SetValueBuilder. This will be the method call for
// ValueBuilders
//
// Example:
//
//     // The following produces equivalent SetValueBuilders
//     SetValue := expression.ListAppend(expression.Value([]string{"a", "list"}), expression.Value([]string{"some", "list"}))
//     SetValue := expression.Value([]string{"a", "list"}).ListAppend(expression.Value([]string{"some", "list"}))
func (vb ValueBuilder) ListAppend(rightOperand OperandBuilder) SetValueBuilder {
	return ListAppend(vb, rightOperand)
}

// IfNotExists will create a SetValueBuilder to be used in as an argument to
// Set(). The first argument must be a NameBuilder representing the name where
// the new item attribute will be created. The second argument can either be a
// NameBuilder or a ValueBuilder. In the case that it is a NameBuilder, the
// value of the item attribute at the name specified will be the value of the
// new item attribute.
// More information: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.UpdateExpressions.html#Expressions.UpdateExpressions.SET.PreventingAttributeOverwrites
//
// Example:
//
//     update, err := expression.Set(expression.Name("someName"), expression.IfNotExists(expression.Name("someName"), expression.Value(5))
func IfNotExists(name NameBuilder, setValue OperandBuilder) SetValueBuilder {
	return SetValueBuilder{
		leftOperand:  name,
		rightOperand: setValue,
		mode:         ifNotExists,
	}
}

// IfNotExists will create a SetValueBuilder.
//
// Example:
//
//     // The following produces equivalent SetValueBuilders
//     SetValue := expression.IfNotExists(expression.Name("someName"), expression.Value(5))
//     SetValue := expression.Name("someName").IfNotExists(expression.Value(5))
func (nb NameBuilder) IfNotExists(rightOperand OperandBuilder) SetValueBuilder {
	return IfNotExists(nb, rightOperand)
}

// BuildOperand will create an Operand struct with an exprNode as a member,
// which is a generic representation of Operands and Conditions. BuildOperand()
// is mainly for the buildTree() method to call on, not for users to invoke.
// BuildOperand aliases all strings to avoid stepping over DynamoDB's reserved
// words.
// More information on reserved words at http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html
func (nb NameBuilder) BuildOperand() (Operand, error) {
	if nb.name == "" {
		return Operand{}, newUnsetParameterError("BuildOperand", "NameBuilder")
	}

	node := exprNode{
		names: []string{},
	}

	nameSplit := strings.Split(nb.name, ".")
	fmtNames := make([]string, 0, len(nameSplit))

	for _, word := range nameSplit {
		var substr string
		if word == "" {
			return Operand{}, newInvalidParameterError("BuildOperand", "NameBuilder")
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
			return Operand{}, newInvalidParameterError("BuildOperand", "NameBuilder")
		}

		// Create a string with special characters that can be substituted later: $p
		node.names = append(node.names, word)
		fmtNames = append(fmtNames, "$n"+substr)
	}
	node.fmtExpr = strings.Join(fmtNames, ".")
	return Operand{
		exprNode: node,
	}, nil
}

// BuildOperand will create an Operand struct with an exprNode as a member,
// which is a generic representation of Operands and Conditions. BuildOperand()
// is mainly for the BuildTree() method to call on, not for users to invoke.
func (vb ValueBuilder) BuildOperand() (Operand, error) {
	expr, err := dynamodbattribute.Marshal(vb.value)
	if err != nil {
		return Operand{}, newInvalidParameterError("BuildOperand", "ValueBuilder")
	}

	// Create a string with special characters that can be substituted later: $v
	operand := Operand{
		exprNode: exprNode{
			values:  []dynamodb.AttributeValue{*expr},
			fmtExpr: "$v",
		},
	}
	return operand, nil
}

// BuildOperand will create an Operand struct with an exprNode as a member,
// which is a generic representation of Operands and Conditions. BuildOperand()
// is mainly for the BuildTree() method to call on, not for users to invoke.
func (sb SizeBuilder) BuildOperand() (Operand, error) {
	operand, err := sb.nameBuilder.BuildOperand()
	operand.exprNode.fmtExpr = "size (" + operand.exprNode.fmtExpr + ")"

	return operand, err
}

// BuildOperand will create an Operand struct with an exprNode as a member,
// which is a generic representation of Operands and Conditions. BuildOperand()
// is mainly for the BuildTree() method to call on, not for users to invoke.
func (kb KeyBuilder) BuildOperand() (Operand, error) {
	if kb.key == "" {
		return Operand{}, newUnsetParameterError("BuildOperand", "KeyBuilder")
	}

	ret := Operand{
		exprNode: exprNode{
			names:   []string{kb.key},
			fmtExpr: "$n",
		},
	}

	return ret, nil
}

// BuildOperand will create an Operand struct with an exprNode as a member,
// which is a generic representation of Operands and Conditions. BuildOperand()
// is mainly for the BuildTree() method to call on, not for users to invoke. If
// the mode of SetValueBuilder is unset, ErrUnsetSetValue will be returned as
// the error.
func (svb SetValueBuilder) BuildOperand() (Operand, error) {
	if svb.mode == unsetValue {
		return Operand{}, newUnsetParameterError("BuildOperand", "SetValueBuilder")
	}

	left, err := svb.leftOperand.BuildOperand()
	if err != nil {
		return Operand{}, err
	}
	leftNode := left.exprNode

	right, err := svb.rightOperand.BuildOperand()
	if err != nil {
		return Operand{}, err
	}
	rightNode := right.exprNode

	node := exprNode{
		children: []exprNode{leftNode, rightNode},
	}

	switch svb.mode {
	case plus:
		node.fmtExpr = "$c + $c"
	case minus:
		node.fmtExpr = "$c - $c"
	case listAppend:
		node.fmtExpr = "list_append($c, $c)"
	case ifNotExists:
		node.fmtExpr = "if_not_exists($c, $c)"
	default:
		return Operand{}, fmt.Errorf("build operand error: unsupported mode: %v", svb.mode)
	}

	return Operand{
		exprNode: node,
	}, nil
}
