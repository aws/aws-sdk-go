package expression

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// ErrUnsetProjection is an error that is returned if buildTree() is called on an
// empty ProjectionBuilder.
var ErrUnsetProjection = awserr.New("UnsetProjection", "buildProjection error: the argument ProjectionBuilder's name list is empty", nil)

// ProjectionBuilder will represent Projection Expressions in DynamoDB. It is
// composed of a list of NameBuilders. ProjectionBuilders will be a building
// block of Builders.
// More Information at: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ProjectionExpressions.html
type ProjectionBuilder struct {
	names []NameBuilder
}

// NamesList will create a ProjectionBuilder with at least one NameBuilder as a
// child. The list of NameBuilders represent the item attribute that will be
// returned after the DynamoDB operation. The resulting ProjectionBuilder can be
// used to build other ProjectionBuilder or to create an Builder to be used in
// an operation input. This will be the function call.
//
// Example:
//
//     projection := expression.NamesList(expression.Name("foo"), expression.Name("bar"))
//
//     anotherProjection := expression.AddNames(projection, expression.Name("baz")) // Used in another projection
//     builder := WithProjection(newProjection)                                     // Used to make an Builder
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
func (nb NameBuilder) NamesList(namesList ...NameBuilder) ProjectionBuilder {
	return NamesList(nb, namesList...)
}

// AddNames will create a new ProjectionBuilder with a list of NameBuilders that
// is a combination of the list from the argument ProjectionBuilder and the
// argument NameBuilder list. The resulting ProjectionBuilder can be used to
// build other ProjectionBuilder or to create an Builder to be used in an
// operation input. This will be the function call.
//
// Example:
//
//     newProjection := expression.AddNames(oldProjection, expression.Name("foo"))
//
//     anotherProjection := expression.AddNames(newProjection, expression.Name("baz")) // Used in another projection
//     builder := WithProjection(newProjection)                                        // Used to make an Builder
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
func (pb ProjectionBuilder) AddNames(namesList ...NameBuilder) ProjectionBuilder {
	return AddNames(pb, namesList...)
}

// buildTree will build a tree structure of exprNodes based on the tree
// structure of the input ProjectionBuilder's child NameBuilders. buildTree()
// satisfies the treeBuilder interface so ProjectionBuilder can be a part of
// Builder and Expression struct.
func (pb ProjectionBuilder) buildTree() (exprNode, error) {
	if len(pb.names) == 0 {
		return exprNode{}, ErrUnsetProjection
	}

	childNodes, err := pb.buildChildNodes()
	if err != nil {
		return exprNode{}, err
	}
	ret := exprNode{
		children: childNodes,
	}

	ret.fmtExpr = "$c" + strings.Repeat(", $c", len(pb.names)-1)

	return ret, nil
}

// buildChildNodes will create the list of the child exprNodes.
func (pb ProjectionBuilder) buildChildNodes() ([]exprNode, error) {
	childNodes := make([]exprNode, 0, len(pb.names))
	for _, name := range pb.names {
		operand, err := name.BuildOperand()
		if err != nil {
			return []exprNode{}, err
		}
		childNodes = append(childNodes, operand.exprNode)
	}

	return childNodes, nil
}
