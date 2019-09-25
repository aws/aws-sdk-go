// +build codegen

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"
)

type BehaviorTestSuite struct {
	BehaviorTest BehaviorTest `json:"tests"`
}

type BehaviorTest struct {
	Defaults Defaults `json:"defaults"`
	Cases    []Case   `json:"cases"`
}

type Defaults struct {
	Env    map[string]string `json:"env"`
	Files  interface{}       `json:"files"`
	Config map[string]string `json:"config"`
}

type Case struct {
	Description string                   `json:"description"`
	LocalEnv    map[string]string        `json:"env"`
	LocalConfig map[string]string        `json:"config"`
	Request     Request                  `json:"request"`
	Response    Response                 `json:"response"`
	Expect      []map[string]interface{} `json:"expect"`
}

type Response struct {
	StatusCode  int               `json:"statusCode"`
	BodyContent interface{}       `json:"bodyContent"`
	BodyType    string            `json:"bodyType"`
	Headers     map[string]string `json:"headers"`
}

type Request struct {
	Operation string                 `json:"operation"`
	Input     map[string]interface{} `json:"input"`
}

func (c Request) BuildInputShape(ref *ShapeRef) string {
	b := NewShapeValueBuilder()
	b.Base64BlobValues = true
	return fmt.Sprintf("&%s{\n%s\n}",
		b.GoType(ref, true),
		b.BuildShape(ref, c.Input, false),
	)
}

//Outputs the string to define an empty shape
func (c Request) EmptyShapeBuilder(ref *ShapeRef) string {
	b := NewShapeValueBuilder()
	return fmt.Sprintf("%s{}", b.GoType(ref, true))
}

func (c Case) BuildOutputShape(ref *ShapeRef) string {
	b := NewShapeValueBuilder()
	b.Base64BlobValues = true
	index := 0
	for i, expectVal := range c.Expect {
		if _, ok := expectVal["responseDataEquals"]; ok {
			index = i
		}
	}
	return fmt.Sprintf("&%s{\n%s\n}",
		b.GoType(ref, true),
		b.BuildShape(ref, c.Expect[index]["responseDataEquals"].(map[string]interface{}), false),
	)
}

// AttachBehaviorTests attaches the Behavior test cases to the API model.
func (a *API) AttachBehaviorTests(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("failed to open behavior tests %s, err: %v", filename, err))
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&a.BehaviorTests); err != nil {
		panic(fmt.Sprintf("failed to decode behavior tests %s, err: %v", filename, err))
	}

}

// APIBehaviorTestsGoCode returns the Go Code string for the Behavior tests.
func (a *API) APIBehaviorTestsGoCode() string {
	w := bytes.NewBuffer(nil)
	a.resetImports()

	a.AddImport("testing")
	a.AddImport("net/http")
	a.AddImport("time")
	a.AddImport("io/ioutil")
	a.AddImport("bytes")
	a.AddImport("strings")
	a.AddImport("encoding/json")

	a.AddSDKImport("aws")
	a.AddSDKImport("awstesting")
	a.AddSDKImport("aws/session")
	a.AddSDKImport("aws/credentials")
	a.AddSDKImport("aws/corehandlers")
	a.AddSDKImport("aws/request")
	a.AddSDKImport("private/protocol")
	a.AddSDKImport("internal/sdktesting")

	a.AddImport(a.ImportPath())

	behaviorTests := struct {
		API *API
		BehaviorTestSuite
	}{
		API:               a,
		BehaviorTestSuite: a.BehaviorTests,
	}

	if err := behaviorTestTmpl.Execute(w, behaviorTests); err != nil {
		panic(fmt.Sprintf("failed to create behavior tests, %v", err))
	}

	ignoreImports := `
	var _ *time.Time
	var _ = protocol.ParseTime
	var _ = strings.NewReader
	var _ = json.Marshal
	`

	return a.importsGoCode() + ignoreImports + w.String()
}

// BuildHeaderMapCompare returns a string which defines a map[string]interface{}
// variable inline
func BuildHeaderMapCompare(input interface{}) string {
	if reflect.ValueOf(input).Kind() != reflect.Map {
		panic("requestHeadersMatch field doesn't contain a map")
	}
	newInput := input.(map[string]interface{})
	output := fmt.Sprint("map[string]awstesting.CompareValue {\n")
	for key, val := range newInput {
		switch key {
		case "Header-Json-Value":
			output += fmt.Sprintf("%#v: awstesting.JSONValueCompareWith(t, %#v),\n", key, val)
		default:
			output += fmt.Sprintf("%#v: awstesting.DefaultCompareWith(t, %#v),\n", key, val)
		}
	}
	output += "}"
	return output
}

//Generates assertions
func (c Case) GenerateAssertions(op *Operation) string {
	var val string = "//Assertions start here"
	for _, assertion := range c.Expect {
		for assertionName, assertionContext := range assertion {
			val += fmt.Sprintf("\nawstesting.Assert")

			switch assertionName {
			case "requestMethodEquals":
				val += fmt.Sprintf("RequestMethodEquals(t, %q, req.HTTPRequest.Method)", assertionContext)
			case "requestUrlMatches":
				val += fmt.Sprintf("RequestURLMatches(t, %q, req.HTTPRequest.URL.String())", assertionContext)
			case "requestUrlPathMatches":
				val += fmt.Sprintf("RequestURLPathMatches(t, %q, req.HTTPRequest.URL.EscapedPath())", assertionContext)
			case "requestUrlQueryMatches":
				val += fmt.Sprintf("RequestURLQueryMatches(t, %q, req.HTTPRequest.URL)", assertionContext)
			case "requestHeadersMatch":
				val += fmt.Sprintf("RequestHeadersMatch(t, %s, req.HTTPRequest.Header)", BuildHeaderMapCompare(assertionContext))
			case "requestBodyEqualsBytes":
				val += fmt.Sprintf("RequestBodyEqualsBytes(t, %q, req)", assertionContext)
			case "requestBodyEqualsJson":
				val += fmt.Sprintf("RequestBodyEqualsJSON(t, %#v, req)", assertionContext)
			case "requestBodyMatchesXml":
				val += fmt.Sprintf("RequestBodyMatchesXML(t, %q, req, %v)", assertionContext, c.Request.EmptyShapeBuilder(&op.InputRef))
			case "requestBodyEqualsString":
				val += fmt.Sprintf("RequestBodyEqualsString(t, %q, req)", assertionContext)
			case "requestIdEquals":
				val += fmt.Sprintf("RequestIDEquals(t, %q, req.RequestID)", assertionContext)
			case "responseDataEquals":
				val += fmt.Sprintf("ResponseValueEquals(t, %v, resp)", c.BuildOutputShape(&op.OutputRef))
			case "responseErrorDataEquals":
				val += fmt.Sprintf("ResponseErrorDataEquals(t, %#v, err)", assertionContext)
			case "responseErrorIsKindOf":
				val += fmt.Sprintf("ResponseErrorIsKindOf(t, %q, err)", assertionContext)
			case "responseErrorMessageEquals":
				val += fmt.Sprintf("ResponseErrorMessageEquals(t, %q, err)", assertionContext)
			case "responseErrorRequestIdEquals":
				val += fmt.Sprintf("ResponseErrorRequestIDEquals(t, %q, err)", assertionContext)
			default:
				panic("Invalid assertion key")
			}
		}
	}
	return val
}

// Returns a value to set Credentials
func (c Case) ConfigurationString(T BehaviorTest) string {
	region := T.Defaults.Env["AWS_REGION"]
	accessKeyId := T.Defaults.Env["AWS_ACCESS_KEY"]
	secretAccessKey := T.Defaults.Env["AWS_SECRET_ACCESS_KEY"]
	endpointUrl := ""
	if len(c.LocalEnv) > 0 {
		for key, val := range c.LocalEnv {
			switch key {
			case "AWS_REGION":
				region = val
			case "AWS_ACCESS_KEY":
				accessKeyId = val
			case "AWS_SECRET_ACCESS_KEY":
				secretAccessKey = val
			case "EndPointURL":
				endpointUrl = val
			}
		}
	}
	if len(c.LocalConfig) > 0 {
		for key, val := range c.LocalConfig {
			switch key {
			case "region":
				region = val
			case "accessKeyId":
				accessKeyId = val
			case "secretAccessKey":
				secretAccessKey = val
			case "endpointUrl":
				endpointUrl = val
			}
		}
	}
	const cfgFmt = `
					Region: aws.String(%#v),
					Credentials: credentials.NewStaticCredentials(%#v, %#v, %#v),
					`
	return fmt.Sprintf(cfgFmt, region, accessKeyId, secretAccessKey, endpointUrl)
}

// ErrorAssertExists returns true if there is an assertion in the expect map
// which requires error as an argument
func (c Case) ErrorAssertExists() bool {
	for _, assertion := range c.Expect {
		for assertionName := range assertion {
			if strings.Contains(assertionName, "Error") {
				return true
			}
		}
	}
	return false
}

var funcMap = template.FuncMap{
	"Map": templateMap,
}

var behaviorTestTmpl = template.Must(template.New(`behaviorTestTmpl`).Funcs(funcMap).Parse(`

{{define "StashCredentials"}}
restoreEnv := sdktesting.StashEnv() //Stashes the current environment
defer restoreEnv()
{{end}}

{{define "SessionSetup"}}
// Starts a new session with credentials and region parsed from "defaults" in the Json file'
sess := session.Must(session.NewSession(&aws.Config{ {{$.testCase.ConfigurationString $.BehaviorTest}} }))
{{end}}

{{define "ResponseBuild"}}
	r.HTTPResponse = &http.Response{
				{{- if eq $.testCase.Response.StatusCode 0}}
					StatusCode: 200,
				{{- else}}
					StatusCode: {{$.testCase.Response.StatusCode}},
				{{- end}}
					Header: http.Header{
								{{- range $key,$val:=$.testCase.Response.Headers}}
									"{{$key}}":[]string{ "{{$val}}" },
								{{- end}}	
							},
				{{- if eq ($.testCase.Response.BodyType) "xml"}}
					Body: ioutil.NopCloser(bytes.NewBufferString({{printf "%q" $.testCase.Response.BodyContent}})),
				{{- else if eq ($.testCase.Response.BodyType) "json"}}
					Body: ioutil.NopCloser(bytes.NewBuffer(
							func() []byte {
								v, err := json.Marshal({{printf "%#v" $.testCase.Response.BodyContent}})
								if err != nil {
									panic(err)
								}
								return v
							}())),
				{{- else}}
					Body: ioutil.NopCloser(&bytes.Buffer{}),
				{{- end}}
				}
{{end}}

{{define "RequestBuild"}}
	input := {{ $.testCase.Request.BuildInputShape $.op.InputRef }}

	//Build request
	req, resp := svc.{{$.testCase.Request.Operation}}Request(input)
	_ = resp

	MockHTTPResponseHandler := request.NamedHandler{Name: "core.SendHandler", Fn: func (r *request.Request){ 
		{{ template "ResponseBuild" Map "testCase" $.testCase -}}	
	}}
	req.Handlers.Send.Swap( corehandlers.SendHandler.Name, MockHTTPResponseHandler )

	err := req.Send()
	if err {{- if $.testCase.ErrorAssertExists}}=={{- else}}!={{- end}} nil {
		t.Fatal(err)
	}

	{{printf "\n"}}
{{end}}

{{- range $i, $testCase := $.BehaviorTest.Cases }}
	// {{printf "%s" $testCase.Description}}
	{{- $op := index $.API.Operations $testCase.Request.Operation }}
	func TestBehavior_{{ printf "%03d" $i }}(t *testing.T) {

		{{template "StashCredentials" .}}

		{{ template "SessionSetup" Map "testCase" $testCase "BehaviorTest" $.BehaviorTest}}

		//Starts a new service using using sess
		svc := {{$.API.PackageName}}.New(sess)

		{{ template "RequestBuild" Map "testCase" $testCase "op" $op}}
		
		{{$testCase.GenerateAssertions $op}}

	}
{{- end }}
`))
