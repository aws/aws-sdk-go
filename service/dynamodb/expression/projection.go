package expression

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// ErrUnsetProjection is an error that is returned if BuildExpression is called
// on an empty ProjectionBuilder.
var ErrUnsetProjection = awserr.New("UnsetProjection", "buildProjection error: the argument ProjectionBuilder's path list is empty", nil)

// ProjectionBuilder will represent Projection Expressions in DynamoDB. It is
// composed of a list of PathBuilders. Users will be able to call the
// BuildExpression() method on a ProjectionBuilder to create an Expression which
// can then be used for operation inputs into DynamoDB.
// More Information at: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ProjectionExpressions.html
type ProjectionBuilder struct {
	paths []PathBuilder
}

// Projection will create a ProjectionBuilder with at least one PathBuilder as a
// child. The list of PathBuilders represent the item attribute that will be
// returned after the DynamoDB operation. The resulting ProjectionBuilder can be
// used to build other ProjectionBuilder or to create an Expression to be used
// in an operation input. This will be the function call.
//
// Example:
//
//     projection := expression.Projection(expression.Path("foo"), expression.Path("bar"))
//
//     anotherProjection := expression.AddPaths(projection, expression.Path("baz")) // Used in another projection
//     expression, err := projection.BuildExpression()                              // Used to make an Expression
func Projection(p PathBuilder, pl ...PathBuilder) ProjectionBuilder {
	pl = append([]PathBuilder{p}, pl...)
	return ProjectionBuilder{
		paths: pl,
	}
}

// Projection will create a ProjectionBuilder. This will be the method call.
//
// Example:
//
//     // The following produces equivalent ProjectionBuilders:
//     projection := expression.Projection(expression.Path("foo"), expression.Path("bar"))
//     projection := expression.Path("foo").Projection(expression.Path("bar"))
func (p PathBuilder) Projection(pl ...PathBuilder) ProjectionBuilder {
	return Projection(p, pl...)
}

// AddPaths will create a new ProjectionBuilder with a list of PathBuilders that
// is a combination of the list from the argument ProjectionBuilder and the
// argument PathBuilder list. The resulting ProjectionBuilder can be used to
// build other ProjectionBuilder or to create an Expression to be used in an
// operation input. This will be the function call.
//
// Example:
//
//     newProjection := expression.AddPaths(oldProjection, expression.Path("foo"))
//
//     anotherProjection := expression.AddPaths(newProjection, expression.Path("baz")) // Used in another projection
//     expression, err := newProjection.BuildExpression()                              // Used to make an Expression
func AddPaths(proj ProjectionBuilder, pl ...PathBuilder) ProjectionBuilder {
	proj.paths = append(proj.paths, pl...)
	return proj
}

// AddPaths will create a ProjectionBuilder. This will be the method call.
//
// Example:
//
//     // The following produces equivalent ProjectionBuilders:
//     newProjection := expression.AddPaths(oldProjection, expression.Path("foo"))
//     newProjection := oldProjection.AddPaths(expression.Path("foo"))
func (proj ProjectionBuilder) AddPaths(pl ...PathBuilder) ProjectionBuilder {
	return AddPaths(proj, pl...)
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
//       ProjectionExpression:      aws.String(expr.Expression),
// 	     ExpressionAttributeNames:  expr.Names,
//       ExpressionAttributeValues: expr.Values,
//       Key: map[string]*dynamodb.AttributeValue{
//         "PartitionKey": &dynamodb.AttributeValue{
//           S: aws.String("SomeKey"),
//         },
//       },
//       TableName: aws.String("SomeTable"),
//     }
func (proj ProjectionBuilder) BuildExpression() (Expression, error) {
	en, err := proj.buildProjection()
	if err != nil {
		return Expression{}, err
	}
	return en.buildExprNodes(&aliasList{})
}

// buildProjection will build a tree structure of ExprNodes based on the tree
// structure of the input ProjectionBuilder's child PathBuilders.
func (proj ProjectionBuilder) buildProjection() (ExprNode, error) {
	if len(proj.paths) == 0 {
		return ExprNode{}, ErrUnsetProjection
	}

	childNodes, err := proj.buildChildNodes()
	if err != nil {
		return ExprNode{}, err
	}
	ret := ExprNode{
		children: childNodes,
	}

	ret.fmtExpr = "$c" + strings.Repeat(", $c", len(proj.paths)-1)

	return ret, nil
}

// buildChildNodes will create the list of the child ExprNodes.
func (proj ProjectionBuilder) buildChildNodes() ([]ExprNode, error) {
	childNodes := make([]ExprNode, 0, len(proj.paths))
	for _, path := range proj.paths {
		en, err := path.BuildOperand()
		if err != nil {
			return []ExprNode{}, err
		}
		childNodes = append(childNodes, en)
	}

	return childNodes, nil
}
