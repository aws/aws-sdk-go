package dynamodbattribute

import (
	"reflect"
	"testing"
)

func TestTagParse(t *testing.T) {
	cases := []struct {
		in             reflect.StructTag
		json, yaml, av bool
		expect         tag
	}{
		{`json:""`, true, false, false, tag{}},
		{`json:"name"`, true, false, false, tag{Name: "name"}},
		{`json:"name,omitempty"`, true, false, false, tag{Name: "name", OmitEmpty: true}},
		{`json:"-"`, true, false, false, tag{Ignore: true}},
		{`json:",omitempty"`, true, false, false, tag{OmitEmpty: true}},
		{`json:",string"`, true, false, false, tag{AsString: true}},
		{`yaml:""`, false, true, false, tag{}},
		{`yaml:"name"`, false, true, false, tag{Name: "name"}},
		{`yaml:"name,omitempty"`, false, true, false, tag{Name: "name", OmitEmpty: true}},
		{`yaml:"-"`, false, true, false, tag{Ignore: true}},
		{`yaml:",omitempty"`, false, true, false, tag{OmitEmpty: true}},
		{`yaml:",string"`, false, true, false, tag{AsString: true}},
		{`dynamodbav:""`, false, false, true, tag{}},
		{`dynamodbav:","`, false, false, true, tag{}},
		{`dynamodbav:"name"`, false, false, true, tag{Name: "name"}},
		{`dynamodbav:"name"`, false, false, true, tag{Name: "name"}},
		{`dynamodbav:"-"`, false, false, true, tag{Ignore: true}},
		{`dynamodbav:",omitempty"`, false, false, true, tag{OmitEmpty: true}},
		{`dynamodbav:",omitemptyelem"`, false, false, true, tag{OmitEmptyElem: true}},
		{`dynamodbav:",string"`, false, false, true, tag{AsString: true}},
		{`dynamodbav:",binaryset"`, false, false, true, tag{AsBinSet: true}},
		{`dynamodbav:",numberset"`, false, false, true, tag{AsNumSet: true}},
		{`dynamodbav:",stringset"`, false, false, true, tag{AsStrSet: true}},
		{`dynamodbav:",stringset,omitemptyelem"`, false, false, true, tag{AsStrSet: true, OmitEmptyElem: true}},
		{`dynamodbav:"name,stringset,omitemptyelem"`, false, false, true, tag{Name: "name", AsStrSet: true, OmitEmptyElem: true}},
	}

	for i, c := range cases {
		actual := tag{}
		if c.json {
			actual.parseStructTag("json", c.in)
		}
		if c.yaml {
			actual.parseStructTag("yaml", c.in)
		}
		if c.av {
			actual.parseAVTag(c.in)
		}
		if e, a := c.expect, actual; !reflect.DeepEqual(e, a) {
			t.Errorf("case %d, expect %v, got %v", i, e, a)
		}
	}
}
