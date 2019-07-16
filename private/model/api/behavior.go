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
	Expect []map[string]interface{}   `json:"expect"`
}

type Request struct{
	Operation string `json:"operation"`
	Input map[string]interface{} `json:"input"`
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
	a.AddImport("time")
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

var funcMap = template.FuncMap{"Map": templateMap}

var behaviorTestTmpl = template.Must(template.New(`behaviorTestTmpl`).Funcs(funcMap).Parse(`

{{define "StashCredentials"}}
	env := awstesting.StashEnv() //Stashes the current environment variables
{{end}}

{{define "SessionSetup"}}
		{{- if len $.testCase.LocalConfig }}
			access_key="{{$.testCase.LocalConfig.AWS_ACCESS_KEY}}"
			secret_access_key="{{$.testCase.LocalConfig.AWS_SECRET_ACCESS_KEY}}"
			aws_region="{{$.testCase.LocalConfig.AWS_REGION}}"
		{{- else}}
			access_key:="{{$.Tests.Defaults.Env.AWS_ACCESS_KEY}}"
			secret_access_key:="{{$.Tests.Defaults.Env.AWS_SECRET_ACCESS_KEY}}"
			aws_region:="{{$.Tests.Defaults.Env.AWS_REGION}}"
		{{- end}}

		//Starts a new session with credentials and region parsed from "defaults" in the Json file'
		sess := session.Must(session.NewSession(&aws.Config{
				 Region: aws.String(aws_region),
				 Credentials: credentials.NewStaticCredentials(access_key, secret_access_key, ""),
			   }))
{{end}}

{{- range $i, $testCase := $.Tests.Cases }}
	//{{printf "%s" $testCase.Description}}
	func BehavTest_{{ printf "%02d" $i }}(t *testing.T) {

		{{template "StashCredentials" .}}
		{{- template "SessionSetup" Map "testCase" $testCase "Tests" $.Tests}}
		
		//Starts a new service using using sess
		svc := {{$.API.PackageName}}.New(sess)

		req, _ := svc.BehavTestRequestGenerator_{{printf "%02d" $i }}("")
		r := req.HTTPRequest
	
		// build request
		req.Build()

		fmt.Println("Write behavior tests here")
	}
{{- end }}
`))
