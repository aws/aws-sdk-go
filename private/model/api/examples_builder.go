// +build codegen

package api

type examplesBuilder interface {
	BuildShape(*ShapeRef, map[string]interface{}, bool) string
	BuildList(string, string, *ShapeRef, []interface{}) string
	BuildComplex(string, string, *ShapeRef, *Shape, map[string]interface{}) string
	GoType(*ShapeRef, bool) string
	Imports(*API) string
}

type defaultExamplesBuilder struct {
	ShapeValueBuilder
}

func NewExamplesBuilder() defaultExamplesBuilder {
	b := defaultExamplesBuilder{
		ShapeValueBuilder: NewShapeValueBuilder(),
	}
	b.ParseTimeString = parseExampleTimeString
	return b
}

func (builder defaultExamplesBuilder) Imports(a *API) string {
	return `"fmt"
	"strings"
	"time"

	"` + SDKImportRoot + `/aws"
	"` + SDKImportRoot + `/aws/awserr"
	"` + SDKImportRoot + `/aws/session"
	"` + SDKImportRoot + `/private/protocol"
	"` + a.ImportPath() + `"
	`
}
