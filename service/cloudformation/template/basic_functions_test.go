package template_test

import (
	"testing"

	. "github.com/aws/aws-sdk-go/service/cloudformation/template"
)

func TestFnBase64(t *testing.T) {
	base64 := FnBase64("some string to encode")
	assertMarshalsTo(t, base64, `{ "Fn::Base64" : "some string to encode" }`)
}

func TestFnBase64WithObjects(t *testing.T) {
	base64 := FnBase64(Ref("some-logical-name"))
	assertMarshalsTo(t, base64, `{ "Fn::Base64" : { "Ref": "some-logical-name" } }`)
}

func TestFnFindInMap(t *testing.T) {
	mapped := FnFindInMap("RegionMap", Ref("AWS::Region"), "32")
	assertMarshalsTo(t, mapped, `{ "Fn::FindInMap" : [ "RegionMap", { "Ref" : "AWS::Region" }, "32"]}`)
}

func TestFnGetAtt(t *testing.T) {
	att := FnGetAtt("MyLB", Ref("something"))
	assertMarshalsTo(t, att, `{ "Fn::GetAtt": [ "MyLB" , { "Ref": "something" } ] }`)
}

func TestFnGetAZs(t *testing.T) {
	azs := FnGetAZs(Ref("AWS::Region"))
	assertMarshalsTo(t, azs, `{ "Fn::GetAZs": { "Ref": "AWS::Region" } }`)
}

func TestFnGetAZsEmpty(t *testing.T) {
	azs := FnGetAZs("")
	assertMarshalsTo(t, azs, `{ "Fn::GetAZs": "" }`)
}

func TestFnGetAZsNil(t *testing.T) {
	azs := FnGetAZs(nil)
	assertMarshalsTo(t, azs, `{ "Fn::GetAZs": "" }`)
}

func TestFnJoin(t *testing.T) {
	joined := FnJoin("\n", "some", "list", Ref("AWS::StackId"))
	assertMarshalsTo(t, joined, `{ "Fn::Join" : ["\n", [ "some", "list", { "Ref" : "AWS::StackId" } ] ] }`)
}

func TestFnSelect(t *testing.T) {
	selected := FnSelect(Ref("MyIndex"), FnGetAZs(""))
	assertMarshalsTo(t, selected, `{ "Fn::Select" : [ { "Ref": "MyIndex" }, { "Fn::GetAZs": "" } ] }`)
}

func TestFnSelectWithLiteralIndex(t *testing.T) {
	selected := FnSelect(2, FnGetAZs(""))
	assertMarshalsTo(t, selected, `{ "Fn::Select" : [ "2", { "Fn::GetAZs": "" } ] }`)
}

func TestRefs(t *testing.T) {
	ref := Ref("AWS::StackId")
	assertMarshalsTo(t, ref, `{"Ref":"AWS::StackId" }`)
}
