package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolvedReferences(t *testing.T) {
	json := `{
		"operations": {
			"OperationName": {
				"input": { "shape": "ShapeName" }
			}
		},
		"shapes": {
			"ShapeName": {
				"type": "structure",
				"members": {
					"memberName1": { "shape": "OtherShape" },
					"memberName2": { "shape": "OtherShape" }
				}
			},
			"OtherShape": { "type": "string" }
		}
	}`
	a := API{}
	a.AttachString(json)
	assert.Equal(t, len(a.Shapes["OtherShape"].refs), 2)
}

func TestRenamedShapes(t *testing.T) {
	json := `{
		"operations": {
			"OperationName": {
				"input": { "shape": "ShapeRequest" },
				"output": { "shape": "ShapeResult" }
			}
		},
		"shapes": {
			"ShapeRequest": {
				"type": "structure",
				"members": {
					"memberName1": { "shape": "ShapeVpnIcmp" },
					"memberName2": { "shape": "ShapeVpnIcmp" }
				}
			},
			"ShapeVpnIcmp": { "type": "string" },
			"ShapeResult": {
				"type": "structure",
				"members": {
					"memberName1": { "shape": "ShapeVpnIcmp" }
				}
			}
		}
	}`
	a := API{}
	a.AttachString(json)
	assert.Nil(t, a.Shapes["ShapeRequest"])
	assert.NotNil(t, a.Shapes["ShapeInput"])
	assert.Nil(t, a.Shapes["ShapeInput"].MemberRefs["memberName1"])
	assert.NotNil(t, a.Shapes["ShapeInput"].MemberRefs["MemberName1"])
	assert.Nil(t, a.Shapes["ShapeInput"].MemberRefs["memberName2"])
	assert.NotNil(t, a.Shapes["ShapeInput"].MemberRefs["MemberName2"])

	assert.Nil(t, a.Shapes["ShapeResult"])
	assert.NotNil(t, a.Shapes["ShapeOutput"])

	assert.Nil(t, a.Shapes["ShapeVpnIcmp"])
	assert.NotNil(t, a.Shapes["ShapeVPNICMP"])
}
