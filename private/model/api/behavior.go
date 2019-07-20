// +build codegen

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

type BehaviorTestSuite struct {
	Defaults Defaults `json:"defaults"`
	Tests Tests `json:"tests"`
}

type Tests struct {
	Defaults Defaults `json:"defaults"`
	Cases []Case `json:"cases"`
}

type Defaults struct{
	Env map[string]string `json:"env"`
	Files interface{} `json:"files"`
	Config interface{} `json:"config"`
}

type Case struct{
	Description string `json:"description"`
	LocalConfig map[string]string `json:"localConfig"`
	Request Request `json:"request"`
	Response Response `json:"response"`
	Expect []map[string]interface{}   `json:"expect"`
}

type Response struct{
	StatusCode int `json:"statusCode"`
	BodyContent string `json:"bodyContent"`
	BodyType string `json: "bodyType"`
	Headers map[string]string `json:"headers"`

}

type Request struct{
	Operation string `json:"operation"`
	Input map[string]interface{} `json:"input"`
}

func (c Request) BuildInputShape(ref *ShapeRef) string {
	var b ShapeValueBuilder
	return fmt.Sprintf("&%s{\n%s\n}",
		b.GoType(ref, true),
		b.BuildShape(ref, c.Input, false),
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
	a.AddImport("context")
	a.AddImport("testing")
	a.AddImport("net/http")
	a.AddImport("time")
	a.AddImport("io/ioutil")
	a.AddImport("bytes")

	a.AddSDKImport("aws")
	a.AddSDKImport("awstesting")
	a.AddSDKImport("aws/client")
	a.AddSDKImport("private/protocol")
	a.AddSDKImport("private/util")
	a.AddSDKImport("aws/request")

	a.AddImport(a.ImportPath())

	behaviorTests := struct {
		API *API
		BehaviorTestSuite
	}{
		API:            a,
		BehaviorTestSuite: a.BehaviorTests,
	}

	if err := behaviorTestTmpl.Execute(w, behaviorTests); err != nil {
		panic(fmt.Sprintf("failed to create behavior tests, %v", err))
	}


	return a.importsGoCode() + w.String()
}

//template map is defined in "eventstream.go"
var funcMap = template.FuncMap{"Map": templateMap}

var behaviorTestTmpl = template.Must(template.New(`behaviorTestTmpl`).Funcs(funcMap).Parse(`

{{define "StashCredentials"}}
	env := awstesting.StashEnv() //Stashes the current environment variables
{{end}}

{{define "SessionSetup"}}
		//Starts a new session with credentials and region parsed from "defaults" in the Json file'
		sess := session.Must(session.NewSession(&aws.Config{
				 Region: aws.String( {{- if and (len $.testCase.LocalConfig) $.testCase.LocalConfig.AWS_REGION }} "{{$.testCase.LocalConfig.AWS_REGION}}" {{- else}} "{{$.Tests.Defaults.Env.AWS_REGION}}" {{- end}}),
				 Credentials: credentials.NewStaticCredentials(
								{{- if and (len $.testCase.LocalConfig) $.testCase.LocalConfig.AWS_ACCESS_KEY }}
									"{{$.testCase.LocalConfig.AWS_ACCESS_KEY}}",							
								{{- else}}
									"{{$.Tests.Defaults.Env.AWS_ACCESS_KEY}}",
								{{- end}}

								{{- if and (len $.testCase.LocalConfig) $.testCase.LocalConfig.AWS_SECRET_ACCESS_KEY }}
									"{{$.testCase.LocalConfig.AWS_SECRET_ACCESS_KEY}}",							
								{{- else}}
									"{{$.Tests.Defaults.Env.AWS_SECRET_ACCESS_KEY}}",
								{{- end}} ""),
			   }))
{{end}}

{{- range $i, $testCase := $.Tests.Cases }}
	//{{printf "%s" $testCase.Description}}
	{{- $op := index $.API.Operations $testCase.Request.Operation }}
	func BehavTest_{{ printf "%02d" $i }}(t *testing.T) {

		{{template "StashCredentials" .}}
		{{- template "SessionSetup" Map "testCase" $testCase "Tests" $.Tests}}

		//Starts a new service using using sess
		svc := {{$.API.PackageName}}.New(sess)

		input := {{ $testCase.Request.BuildInputShape $op.InputRef }}

		req, resp := svc.{{$testCase.Request.Operation}}Request(input)

   		err := req.Send()
		if err == nil { // resp is now filled
			fmt.Println(resp)
		}
		{{- if eq $testCase.Response.StatusCode 0}}
			response := Response{StatusCode:200}
		{{- else }}
		  response := &http.Response{
							StatusCode:{{$testCase.Response.StatusCode}},
						{{- if ne (len $testCase.Response.Headers) 0}}
							Header: http.Header{
										{{- range $key,$val:=$testCase.Response.Headers}}
											"{{$key}}":[]string{ "{{$val}}" },
										{{- end}}	
									},
						{{- end}}
						{{- if ne (len $testCase.Response.BodyContent) 0}}
							Body: ioutil.NopCloser(bytes.NewBufferString({{printf "%q" $testCase.Response.BodyContent}})),
						{{- end}}
					}

		{{- end}}	

	}
{{- end }}
`))
