package ini

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	xId, _, _ := newLitToken([]byte("x = 1234"))

	regionId, _, _ := newLitToken([]byte("region"))
	regionLit, _, _ := newLitToken([]byte(`"us-west-2"`))
	regionNoQuotesLit, _, _ := newLitToken([]byte("us-west-2"))

	credentialId, _, _ := newLitToken([]byte("credential_source"))
	ec2MetadataLit, _, _ := newLitToken([]byte("Ec2InstanceMetadata"))

	outputId, _, _ := newLitToken([]byte("output"))
	outputLit, _, _ := newLitToken([]byte("json"))

	equalOp, _, _ := newOpToken([]byte("= 1234"))
	numLit, _, _ := newLitToken([]byte("1234"))
	defaultId, _, _ := newLitToken([]byte("default"))
	assumeId, _, _ := newLitToken([]byte("assumerole"))

	cases := []struct {
		r             io.Reader
		expectedStack []AST
		expectedError bool
	}{
		{
			r: bytes.NewBuffer([]byte(`;foo`)),
			expectedStack: []AST{
				newCommentStatement(CommentToken{comment: ";foo"}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`# foo`)),
			expectedStack: []AST{
				newCommentStatement(CommentToken{comment: "# foo"}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`// foo`)),
			expectedStack: []AST{
				newCommentStatement(CommentToken{comment: "// foo"}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`;foo
	//bar
			# baz
			`)),
			expectedStack: []AST{
				newCommentStatement(CommentToken{comment: ";foo"}),
				newCommentStatement(CommentToken{comment: "//bar"}),
				newCommentStatement(CommentToken{comment: "# baz"}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`x = 1234`)),
			expectedStack: []AST{
				newExprStatement(EqualExpr{
					Left:  newExpression(xId),
					Root:  equalOp,
					Right: newExpression(numLit),
				}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`x=1234`)),
			expectedStack: []AST{
				newExprStatement(EqualExpr{
					Left:  newExpression(xId),
					Root:  equalOp,
					Right: newExpression(numLit),
				}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`[ default ]`)),
			expectedStack: []AST{
				newTableStatement(defaultId),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`[default]`)),
			expectedStack: []AST{
				newTableStatement(defaultId),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`[default]
					region="us-west-2"`)),
			expectedStack: []AST{
				newTableStatement(defaultId),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionId),
					Root:  equalOp,
					Right: newExpression(regionLit),
				}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`[default]
region = us-west-2
credential_source = Ec2InstanceMetadata
output = json

[assumerole]
output = json
region = us-west-2
		`)),
			expectedStack: []AST{
				newTableStatement(defaultId),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionId),
					Root:  equalOp,
					Right: newExpression(regionNoQuotesLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(credentialId),
					Root:  equalOp,
					Right: newExpression(ec2MetadataLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(outputId),
					Root:  equalOp,
					Right: newExpression(outputLit),
				}),
				newTableStatement(assumeId),
				newExprStatement(EqualExpr{
					Left:  newExpression(outputId),
					Root:  equalOp,
					Right: newExpression(outputLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionId),
					Root:  equalOp,
					Right: newExpression(regionNoQuotesLit),
				}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`[default]
s3 =
	foo=bar
	bar=baz
region = us-west-2
credential_source = Ec2InstanceMetadata
output = json

[assumerole]
output = json
region = us-west-2
		`)),
			expectedStack: []AST{
				newTableStatement(defaultId),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionId),
					Root:  equalOp,
					Right: newExpression(regionNoQuotesLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(credentialId),
					Root:  equalOp,
					Right: newExpression(ec2MetadataLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(outputId),
					Root:  equalOp,
					Right: newExpression(outputLit),
				}),
				newTableStatement(assumeId),
				newExprStatement(EqualExpr{
					Left:  newExpression(outputId),
					Root:  equalOp,
					Right: newExpression(outputLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionId),
					Root:  equalOp,
					Right: newExpression(regionNoQuotesLit),
				}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`[default]
region = us-west-2
credential_source = Ec2InstanceMetadata
s3 =
	foo=bar
	bar=baz
output = json

[assumerole]
output = json
region = us-west-2
		`)),
			expectedStack: []AST{
				newTableStatement(defaultId),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionId),
					Root:  equalOp,
					Right: newExpression(regionNoQuotesLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(credentialId),
					Root:  equalOp,
					Right: newExpression(ec2MetadataLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(outputId),
					Root:  equalOp,
					Right: newExpression(outputLit),
				}),
				newTableStatement(assumeId),
				newExprStatement(EqualExpr{
					Left:  newExpression(outputId),
					Root:  equalOp,
					Right: newExpression(outputLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionId),
					Root:  equalOp,
					Right: newExpression(regionNoQuotesLit),
				}),
			},
		},
	}

	for i, c := range cases {
		stack, err := Parse(c.r)

		if e, a := c.expectedError, err != nil; e != a {
			t.Errorf("%d: expected %t, but received %t with error %v", i+1, e, a, err)
		}

		if e, a := c.expectedStack, stack; !reflect.DeepEqual(e, a) {
			t.Errorf("%d: expected %v, but received %v", i+1, e, a)
		}
	}
}
