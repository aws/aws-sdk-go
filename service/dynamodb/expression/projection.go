package expression

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// ErrUnsetProjection is an error that is returned if BuildExpression is called
// on an empty ProjectionBuilder.
var ErrUnsetProjection = awserr.New("UnsetProjection", "buildProjection error: the argument ProjectionBuilder's name list is empty", nil)

// ProjectionBuilder will represent Projection Expressions in DynamoDB. It is
// composed of a list of NameBuilders. Users will be able to call the
// BuildExpression() method on a ProjectionBuilder to create an Expression which
// can then be used for operation inputs into DynamoDB.
// More Information at: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ProjectionExpressions.html
type ProjectionBuilder struct {
	names []NameBuilder
}

// NamesList will create a ProjectionBuilder with at least one NameBuilder as a
// child. The list of NameBuilders represent the item attribute that will be
// returned after the DynamoDB operation. The resulting ProjectionBuilder can be
// used to build other ProjectionBuilder or to create an Expression to be used
// in an operation input. This will be the function call.
//
// Example:
//
//     projection := expression.NamesList(expression.Name("foo"), expression.Name("bar"))
//
//     anotherProjection := expression.AddNames(projection, expression.Name("baz")) // Used in another projection
//     expression, err := projection.BuildExpression()                              // Used to make an Expression
func NamesList(nameBuilder NameBuilder, namesList ...NameBuilder) ProjectionBuilder {
	namesList = append([]NameBuilder{nameBuilder}, namesList...)
	return ProjectionBuilder{
		names: namesList,
	}
}

// NamesList will create a ProjectionBuilder. This will be the method call.
//
// Example:
//
//     // The following produces equivalent ProjectionBuilders:
//     projection := expression.NamesList(expression.Name("foo"), expression.Name("bar"))
//     projection := expression.Name("foo").NamesList(expression.Name("bar"))
func (nameBuilder NameBuilder) NamesList(namesList ...NameBuilder) ProjectionBuilder {
	return NamesList(nameBuilder, namesList...)
}

// AddNames will create a new ProjectionBuilder with a list of NameBuilders that
// is a combination of the list from the argument ProjectionBuilder and the
// argument NameBuilder list. The resulting ProjectionBuilder can be used to
// build other ProjectionBuilder or to create an Expression to be used in an
// operation input. This will be the function call.
//
// Example:
//
//     newProjection := expression.AddNames(oldProjection, expression.Name("foo"))
//
//     anotherProjection := expression.AddNames(newProjection, expression.Name("baz")) // Used in another projection
//     expression, err := newProjection.BuildExpression()                              // Used to make an Expression
func AddNames(projectionBuilder ProjectionBuilder, namesList ...NameBuilder) ProjectionBuilder {
	projectionBuilder.names = append(projectionBuilder.names, namesList...)
	return projectionBuilder
}

// AddNames will create a ProjectionBuilder. This will be the method call.
//
// Example:
//
//     // The following produces equivalent ProjectionBuilders:
//     newProjection := expression.AddNames(oldProjection, expression.Name("foo"))
//     newProjection := oldProjection.AddNames(expression.Name("foo"))
func (projectionBuilder ProjectionBuilder) AddNames(namesList ...NameBuilder) ProjectionBuilder {
	return AddNames(projectionBuilder, namesList...)
}

// BuildExpression will take an ProjectionBuilder as input and output an
// Expression which can be used in DynamoDB operational inputs (i.e.
// GetItemInput, QueryInput, etc) In the future, the Expression struct
// can be used in some injection method into the input structs.
//
// Example:
//
//     expr, err := someProjection.BuildExpression()
//
//     getItemInput := dynamodb.GetItemInput{
//       ProjectionExpression:      expr.ProjectionExpression(),
// 	     ExpressionAttributeNames:  expr.Names(),
//       ExpressionAttributeValues: expr.Values(),
//       Key: map[string]*dynamodb.AttributeValue{
//         "PartitionKey": &dynamodb.AttributeValue{
//           S: aws.String("SomeKey"),
//         },
//       },
//       TableName: aws.String("SomeTable"),
//     }
func (projectionBuilder ProjectionBuilder) BuildExpression() (Expression, error) {
	return Expression{
		expressionMap: map[string]TreeBuilder{
			"projection": projectionBuilder,
		},
	}, nil
}

// BuildTree will build a tree structure of ExprNodes based on the tree
// structure of the input ProjectionBuilder's child NameBuilders.
func (projectionBuilder ProjectionBuilder) BuildTree() (ExprNode, error) {
	if len(projectionBuilder.names) == 0 {
		return ExprNode{}, ErrUnsetProjection
	}

	childNodes, err := projectionBuilder.buildChildNodes()
	if err != nil {
		return ExprNode{}, err
	}
	ret := ExprNode{
		children: childNodes,
	}

	ret.fmtExpr = "$c" + strings.Repeat(", $c", len(projectionBuilder.names)-1)

	return ret, nil
}

// buildChildNodes will create the list of the child ExprNodes.
func (projectionBuilder ProjectionBuilder) buildChildNodes() ([]ExprNode, error) {
	childNodes := make([]ExprNode, 0, len(projectionBuilder.names))
	for _, name := range projectionBuilder.names {
		en, err := name.BuildOperand()
		if err != nil {
			return []ExprNode{}, err
		}
		childNodes = append(childNodes, en)
	}

	return childNodes, nil
}
