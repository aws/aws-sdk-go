// +build codegen

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"
	"strings"
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

func (c Request) EmptyShapeBuilder(ref *ShapeRef) string{
	var b ShapeValueBuilder
	return fmt.Sprintf("%s{\n%s\n}",
		b.GoType(ref, true),
		b.BuildShape(ref, map[string]interface{} {}, false),
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
	//a.AddImport("context")
	a.AddImport("testing")
	a.AddImport("net/http")
	a.AddImport("fmt")
	a.AddImport("time")
	a.AddImport("io/ioutil")
	a.AddImport("bytes")
	a.AddImport("net/url")
	a.AddImport("net/textproto")
	a.AddImport("strings")
	a.AddImport("encoding/base64")

	a.AddSDKImport("aws")
	a.AddSDKImport("awstesting")
	a.AddSDKImport("aws/session")
	a.AddSDKImport("aws/credentials")

	//a.AddSDKImport("aws/client")
	//a.AddSDKImport("private/protocol")
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

	ignoreImports := `

	`
	return a.importsGoCode() + ignoreImports + w.String()
}

//template map is defined in "eventstream.go"
var funcMap = template.FuncMap{"Map": templateMap,"Contains": strings.Contains}

//	defer env()//Might need to comment this out

var behaviorTestTmpl = template.Must(template.New(`behaviorTestTmpl`).Funcs(funcMap).Parse(`

{{define "StashCredentials"}}
	env := awstesting.StashEnv() //Stashes the current environment variables
	fmt.Println(env)
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

{{define "Assertions"}}
		//Assertions start here
		{{- range $k,$assertion :=$.testCase.Expect}}
			{{- range $assertionName,$assertionContext:=$assertion}}
				{{- if Contains $assertionName "request" }}
					{{- if eq $assertionName "requestBodyMatchesXml"}}
						if !{{$assertionName}}_assert(t , req, {{printf "%q" $assertionContext}}, {{ $.testCase.Request.EmptyShapeBuilder $.op.InputRef }} )
					{{- else}} {{if eq $assertionName "requestHeadersMatch"}}
						if !{{$assertionName}}_assert(t , req, {{printf "%#v" $assertionContext}})
					{{- else}} 
						if !{{$assertionName}}_assert(t , req, "{{$assertionContext}}") 
					{{- end}} {{- end}} {
						t.Error("expect no error, got{{printf "%s" $assertionName}} assertion failed")
						}
				{{- else}}
 						if !{{$assertionName}}_assert(t , response, {{printf "%#v" $assertionContext}}){
							t.Error("expect no error, got{{printf "%s" $assertionName}} assertion failed")
						}
				{{- end}}

			{{- end}}
			{{printf "\n"}}
		{{- end}}
{{end}}

{{define "ResponseBuild"}}
		{{- if eq $.testCase.Response.StatusCode 0}}
			response := &http.Response{StatusCode:200}
		{{- else }}
		  response := &http.Response{
							StatusCode:{{$.testCase.Response.StatusCode}},
						{{- if ne (len $.testCase.Response.Headers) 0}}
							Header: http.Header{
										{{- range $key,$val:=$.testCase.Response.Headers}}
											"{{$key}}":[]string{ "{{$val}}" },
										{{- end}}	
									},
						{{- end}}
						{{- if ne (len $.testCase.Response.BodyContent) 0}}
							Body: ioutil.NopCloser(bytes.NewBufferString({{printf "%q" $.testCase.Response.BodyContent}})),
						{{- end}}
					}
		{{- end}}	
		_ = response
		{{printf "\n"}}
{{end}}

{{define "RequestBuild"}}

		input := {{ $.testCase.Request.BuildInputShape $.op.InputRef }}
		req, resp := svc.{{$.testCase.Request.Operation}}Request(input)
		_ = resp

   		err := req.Send()
		if err != nil { 
			t.Errorf("expect no error, got %v", err)
		}
		{{printf "\n"}}
{{end}}

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

func requestMethodEquals_assert(t *testing.T, req *request.Request,val string) bool{
	if req.HTTPRequest.Method==val{
		return true
	}
	return false
}

func requestUrlMatches_assert(t *testing.T, req *request.Request,val string) bool{
	if req.HTTPRequest.URL.String()==val{
		return true
	}
	return false
}

func requestUrlPathMatches_assert(t *testing.T, req *request.Request,val string) bool{

	if req.HTTPRequest.URL.RequestURI()==val{
		return true
	}
	return false
}

func requestUrlQueryMatches_assert(t *testing.T, req *request.Request,val string) bool{
	u, err := url.Parse(val) // parsed val into a structure
	if err!=nil{
		t.Errorf("expect no error, got %v",err)
	}
	query_request := req.HTTPRequest.URL.Query() //parsed RawQuery of "req" to get the values inside
	query_val := u.Query() //parsed RawQuery of "val" to get the values inside
	
	if query_request.Encode() == query_val.Encode(){
		return true
	}
	return false
}

func requestHeadersMatch_assert(t *testing.T, req *request.Request,header map[string]interface{}) bool{
	for key, val_expect := range header{
		if val_req, ok := req.HTTPRequest.Header[textproto.CanonicalMIMEHeaderKey(key)]; ok {
			if val_req[0] != val_expect{
				return false
			}
		} else{
			return false
		}
	}
	return true
}

func requestBodyEqualsString_assert(t *testing.T, req *request.Request,val string) bool{

	var bytes_req_body []byte
	var err error
	if req.HTTPRequest.Body!= nil {
  		bytes_req_body, err  = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}
	}

	req.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bytes_req_body))
	string_req_body := string(bytes_req_body)

	if string_req_body == val{
		return true
	}
	return false
}

func requestBodyEqualsBytes_assert(t *testing.T, req *request.Request,val string) bool{

	var bytes_req_body []byte

	bytes_expect, err := base64.StdEncoding.DecodeString(val)

	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	if req.HTTPRequest.Body!= nil {
  		bytes_req_body, err = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}
	}

	req.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bytes_req_body))

	if bytes.Compare(bytes_req_body, bytes_expect) == 0 {
		return true
	}
	return false
}

func requestBodyMatchesXml_assert(t *testing.T, req *request.Request,val string,container interface{}) bool{
	var bytes_req_body []byte
	var err error
	if req.HTTPRequest.Body != nil {
		bytes_req_body, err = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}
	}
	req.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bytes_req_body))

	if !awstesting.AssertXML(t, val, util.Trim(string(bytes_req_body)),container ) {
		return false
	}

	return true
}

func requestBodyEqualsJson(t *testing.T, req *request.Request,val string) bool{
	
	var bytes_req_body []byte
	var err error
	if req.HTTPRequest.Body!= nil {
  		bytes_req_body, err  = ioutil.ReadAll(req.HTTPRequest.Body)
		if err != nil {
			t.Errorf("expect no error, got %v", err)
		}
	}

	req.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(bytes_req_body))
	string_req_body := string(bytes_req_body)

	if ! awstesting.AssertJSON(t, string_req_body, val){
		return false
	}

	return true
}

func requestIdEquals_assert(t *testing.T, req *request.Request,val string) bool{
	if req.RequestID == val{
		return true
	}
	return false
}

func responseDataEquals_assert(t *testing.T, req *request.Request,val map[string]interface{}) bool{
    if testing.Short() {
        t.Skip("skipping responseDataEquals assertion")
    }
	return true
}

func responseErrorIsKindOf_assert(t *testing.T, req *request.Request,val map[string]interface{}){
    if testing.Short() {
        t.Skip("skipping responseErrorIsKindOf assertion")
    }
}


func responseErrorMessageEquals_assert(t *testing.T, req *request.Request,val map[string]interface{}){
    if testing.Short() {
        t.Skip("skipping responseErrorMessageEquals assertion")
    }
}


func responseErrorDataEquals_assert(t *testing.T, req *request.Request,val map[string]interface{}){
    if testing.Short() {
        t.Skip("skipping responseErrorDataEquals assertion")
    }
}

func responseErrorRequestIdEquals_assert(t *testing.T, req *request.Request,val map[string]interface{}){
    if testing.Short() {
        t.Skip("skipping responseErrorRequestIdEquals assertion")
    }
}


{{- range $i, $testCase := $.Tests.Cases }}
	//{{printf "%s" $testCase.Description}}
	{{- $op := index $.API.Operations $testCase.Request.Operation }}
	func BehavTest_{{ printf "%02d" $i }}(t *testing.T) {

		{{template "StashCredentials" .}}

		{{- template "SessionSetup" Map "testCase" $testCase "Tests" $.Tests}}

		//Starts a new service using using sess
		svc := {{$.API.PackageName}}.New(sess)

		{{- template "RequestBuild" Map "testCase" $testCase "op" $op}}

		{{- template "ResponseBuild" Map "testCase" $testCase}}

		{{- template "Assertions" Map "testCase" $testCase "op" $op}}

	}
{{- end }}
`))
