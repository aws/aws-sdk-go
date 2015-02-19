package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolvedReferences(t *testing.T) {
	json := `
{
  "operations": {
  	"OperationName": {
  	  "input": {
  	    "shape": "ShapeName"
  	  }
  	}
  },
  "shapes": {
     "ShapeName": {
       "type": "structure",
       "members": {
         "memberName1": {
         	"shape": "OtherShape"
         },
	     "memberName2": {
	     	"shape": "OtherShape"
	     }
       }
     },
     "OtherShape": {
     	"type": "string"
     }
  }
}
`
	a := API{}
	a.AttachString(json)
	assert.Equal(t, len(a.Shapes["OtherShape"].refs), 2)
}
