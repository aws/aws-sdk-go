// +build codegen

package api

import (
	"reflect"
	"strconv"
	"testing"
)

func TestResolvedReferences(t *testing.T) {
	json := `{
		"operations": {
			"OperationName": {
				"input": { "shape": "TestName" }
			}
		},
		"shapes": {
			"TestName": {
				"type": "structure",
				"members": {
					"memberName1": { "shape": "OtherTest" },
					"memberName2": { "shape": "OtherTest" }
				}
			},
			"OtherTest": { "type": "string" }
		}
	}`
	a := API{}
	a.AttachString(json)
	if len(a.Shapes["OtherTest"].refs) != 2 {
		t.Errorf("Expected %d, but received %d", 2, len(a.Shapes["OtherTest"].refs))
	}
}

func TestTrimModelServiceVersions(t *testing.T) {
	cases := []struct {
		Paths  []string
		Expect []string
	}{
		{
			Paths: []string{
				"foo/baz/2018-01-02",
				"foo/baz/2019-01-02",
				"foo/baz/2017-01-02",
				"foo/bar/2019-01-02",
				"foo/bar/2013-04-02",
				"foo/bar/2019-01-03",
			},
			Expect: []string{
				"foo/bar/2019-01-03",
				"foo/baz/2019-01-02",
			},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := trimModelServiceVersions(c.Paths)
			if e, a := c.Expect, result; !reflect.DeepEqual(e, a) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}
