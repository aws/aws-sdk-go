package ini

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	xID, _, _ := newLitToken([]rune("x = 1234"))
	s3ID, _, _ := newLitToken([]rune("s3 = 1234"))

	regionID, _, _ := newLitToken([]rune("region"))
	regionLit, _, _ := newLitToken([]rune(`"us-west-2"`))
	regionNoQuotesLit, _, _ := newLitToken([]rune("us-west-2"))

	credentialID, _, _ := newLitToken([]rune("credential_source"))
	ec2MetadataLit, _, _ := newLitToken([]rune("Ec2InstanceMetadata"))

	outputID, _, _ := newLitToken([]rune("output"))
	outputLit, _, _ := newLitToken([]rune("json"))

	equalOp, _, _ := newOpToken([]rune("= 1234"))
	equalColonOp, _, _ := newOpToken([]rune(": 1234"))
	numLit, _, _ := newLitToken([]rune("1234"))
	defaultID, _, _ := newLitToken([]rune("default"))
	assumeID, _, _ := newLitToken([]rune("assumerole"))

	cases := []struct {
		r             io.Reader
		expectedStack []AST
		expectedError bool
	}{
		{
			r: bytes.NewBuffer([]byte(`;foo`)),
			expectedStack: []AST{
				newCommentStatement(commentToken{comment: ";foo"}),
			},
		},
		{
			r:             bytes.NewBuffer([]byte(`0==0`)),
			expectedError: true,
		},
		{
			r:             bytes.NewBuffer([]byte(`0=:0`)),
			expectedError: true,
		},
		{
			r:             bytes.NewBuffer([]byte(`0:=0`)),
			expectedError: true,
		},
		{
			r:             bytes.NewBuffer([]byte(`0::0`)),
			expectedError: true,
		},
		{
			r: bytes.NewBuffer([]byte(`[ default ]x`)),
			expectedStack: []AST{
				newCompletedSectionStatement(
					newSectionStatement(defaultID),
				),
				newExpression(xID),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`# foo`)),
			expectedStack: []AST{
				newCommentStatement(commentToken{comment: "# foo"}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`// foo`)),
			expectedStack: []AST{
				newCommentStatement(commentToken{comment: "// foo"}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`;foo
			//bar
					# baz
					`)),
			expectedStack: []AST{
				newCommentStatement(commentToken{comment: ";foo"}),
				newCommentStatement(commentToken{comment: "//bar"}),
				newCommentStatement(commentToken{comment: "# baz"}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`x = 1234`)),
			expectedStack: []AST{
				newExprStatement(EqualExpr{
					Left:  newExpression(xID),
					Root:  equalOp,
					Right: newExpression(numLit),
				}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`x=1234`)),
			expectedStack: []AST{
				newExprStatement(EqualExpr{
					Left:  newExpression(xID),
					Root:  equalOp,
					Right: newExpression(numLit),
				}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`x : 1234`)),
			expectedStack: []AST{
				newExprStatement(EqualExpr{
					Left:  newExpression(xID),
					Root:  equalColonOp,
					Right: newExpression(numLit),
				}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`x:1234`)),
			expectedStack: []AST{
				newExprStatement(EqualExpr{
					Left:  newExpression(xID),
					Root:  equalColonOp,
					Right: newExpression(numLit),
				}),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`[ default ]`)),
			expectedStack: []AST{
				newCompletedSectionStatement(
					newSectionStatement(defaultID),
				),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`[default]`)),
			expectedStack: []AST{
				newCompletedSectionStatement(
					newSectionStatement(defaultID),
				),
			},
		},
		{
			r: bytes.NewBuffer([]byte(`[default]
							region="us-west-2"`)),
			expectedStack: []AST{
				newCompletedSectionStatement(
					newSectionStatement(defaultID),
				),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionID),
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
				newCompletedSectionStatement(
					newSectionStatement(defaultID),
				),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionID),
					Root:  equalOp,
					Right: newExpression(regionNoQuotesLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(credentialID),
					Root:  equalOp,
					Right: newExpression(ec2MetadataLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(outputID),
					Root:  equalOp,
					Right: newExpression(outputLit),
				}),
				newCompletedSectionStatement(
					newSectionStatement(assumeID),
				),
				newExprStatement(EqualExpr{
					Left:  newExpression(outputID),
					Root:  equalOp,
					Right: newExpression(outputLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionID),
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
				newCompletedSectionStatement(
					newSectionStatement(defaultID),
				),
				newSkipStatement(newEqualExpr(newExpression(s3ID), equalOp)),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionID),
					Root:  equalOp,
					Right: newExpression(regionNoQuotesLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(credentialID),
					Root:  equalOp,
					Right: newExpression(ec2MetadataLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(outputID),
					Root:  equalOp,
					Right: newExpression(outputLit),
				}),
				newCompletedSectionStatement(
					newSectionStatement(assumeID),
				),
				newExprStatement(EqualExpr{
					Left:  newExpression(outputID),
					Root:  equalOp,
					Right: newExpression(outputLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionID),
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
				newCompletedSectionStatement(
					newSectionStatement(defaultID),
				),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionID),
					Root:  equalOp,
					Right: newExpression(regionNoQuotesLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(credentialID),
					Root:  equalOp,
					Right: newExpression(ec2MetadataLit),
				}),
				newSkipStatement(newEqualExpr(newExpression(s3ID), equalOp)),
				newExprStatement(EqualExpr{
					Left:  newExpression(outputID),
					Root:  equalOp,
					Right: newExpression(outputLit),
				}),
				newCompletedSectionStatement(
					newSectionStatement(assumeID),
				),
				newExprStatement(EqualExpr{
					Left:  newExpression(outputID),
					Root:  equalOp,
					Right: newExpression(outputLit),
				}),
				newExprStatement(EqualExpr{
					Left:  newExpression(regionID),
					Root:  equalOp,
					Right: newExpression(regionNoQuotesLit),
				}),
			},
		},
	}

	for i, c := range cases {
		fmt.Println("running", i)
		stack, err := ParseAST(c.r)

		if e, a := c.expectedError, err != nil; e != a {
			t.Errorf("%d: expected %t, but received %t with error %v", i, e, a, err)
		}

		if e, a := c.expectedStack, stack; !reflect.DeepEqual(e, a) {
			t.Errorf("%d: expected %v, but received %v", i, e, a)
		}
	}
}
