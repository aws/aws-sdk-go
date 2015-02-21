package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/awslabs/aws-sdk-go/internal/model/api"
	"github.com/awslabs/aws-sdk-go/internal/util"
)

type InputTestSuite struct {
	*api.API
	Description string
	Cases       []TestCase
}

type TestCase struct {
	Given      *api.Operation
	Params     interface{}     `json:",omitempty"`
	Data       interface{}     `json:"result,omitempty"`
	InputTest  TestExpectation `json:"serialized"`
	OutputTest TestExpectation `json:"response"`
}

type TestExpectation struct {
	Body      string
	URI       string
	Headers   map[string]string
	StatuCode uint `json:"status_code"`
}

const preamble = `
var reTrim = regexp.MustCompile("\\s")
func trim(s string) string {
	return reTrim.ReplaceAllString(s, "")
}
`

var reStripSpace = regexp.MustCompile(`\s(\w)`)

var reImportRemoval = regexp.MustCompile(`(?s:import \((.+?)\))`)

func removeImports(code string) string {
	return reImportRemoval.ReplaceAllString(code, "")
}

var extraImports = []string{
	"encoding/json",
	"io/ioutil",
	"regexp",
	"testing",
	"github.com/stretchr/testify/assert",
}

func addImports(code string) string {
	importNames := make([]string, len(extraImports))
	for i, n := range extraImports {
		importNames[i] = fmt.Sprintf("%q", n)
	}
	str := reImportRemoval.ReplaceAllString(code, "import (\n$1\n"+strings.Join(importNames, "\n")+")")
	return strings.Replace(str, `"github.com/awslabs/aws-sdk-go/aws/signer/v4"`, "", -1)
}

func removeSigner(code string) string {
	return strings.Replace(code, `service.Handlers.Sign.PushBack(v4.Sign)`, "", -1)
}

func (i *InputTestSuite) TestSuite() string {
	var buf bytes.Buffer

	prefix := reStripSpace.ReplaceAllStringFunc(i.Description, func(x string) string {
		return strings.ToUpper(x[1:])
	})

	for idx, c := range i.Cases {
		opName := prefix + "Case" + strconv.Itoa(idx+1)
		buf.WriteString(c.TestCase(i.API.StructName(), opName) + "\n")
	}
	return util.GoFmt(buf.String())
}

var tplInputTestCase = template.Must(template.New("inputcase").Parse(`
func Test{{ .OpName }}(t *testing.T) {
	svc := New{{ .ServiceName }}(nil)

	var input {{ .Given.InputRef.ShapeName }}
	json.Unmarshal([]byte({{ .ParamsString }}), &input)
	req := svc.{{ .Given.ExportedName }}Request(&input)
	req.Build()
	r := req.HTTPRequest

	// assert body
	body, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, trim(string(body)), trim({{ .Body }}))

	// assert URL
	assert.Equal(t, r.URL.Path, "{{ .TestCase.InputTest.URI }}")

	// assert headers
{{ range $k, $v := .TestCase.InputTest.Headers }}assert.Equal(t, r.Header.Get("{{ $k }}"), "{{ $v }}")
{{ end }}
}
`))

type tplInputTestCaseData struct {
	*TestCase
	ServiceName, Body, OpName, ParamsString string
}

var tplOutputTestCase = template.Must(template.New("outputcase").Parse(`
func Test{{ .Name }}(t *testing.T) {
	assert.Equal(t, req.HttpRequest.Body, {{ .Body }})

{{ range $k, $v := .TestCase.OutputTest.Headers }}
{{ end }}
}
`))

type tplOutputTestCaseData struct {
	*TestCase
	Body, ServiceName, OpName string
}

func (i *TestCase) TestCase(svcName, opName string) string {
	var buf bytes.Buffer

	if i.Params != nil { // input test
		pBuf, _ := json.Marshal(i.Params)
		input := tplInputTestCaseData{
			TestCase:     i,
			Body:         fmt.Sprintf("%q", i.InputTest.Body),
			OpName:       strings.ToUpper(opName[0:1]) + opName[1:],
			ServiceName:  svcName,
			ParamsString: fmt.Sprintf("%q", pBuf),
		}

		if err := tplInputTestCase.Execute(&buf, input); err != nil {
			panic(err)
		}
	} else {
		output := tplOutputTestCaseData{
			TestCase:    i,
			Body:        fmt.Sprintf("%q", i.OutputTest.Body),
			OpName:      strings.ToUpper(opName[0:1]) + opName[1:],
			ServiceName: svcName,
		}

		if err := tplOutputTestCase.Execute(&buf, output); err != nil {
			panic(err)
		}
	}

	return util.GoFmt(buf.String())
}

func GenerateInputTestSuite(filename string) string {
	var buf bytes.Buffer
	buf.WriteString("package protocol_test\n\n")

	var suites []InputTestSuite
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(f).Decode(&suites)
	if err != nil {
		panic(err)
	}

	var innerBuf bytes.Buffer
	for i, suite := range suites {
		svcPrefix := "Service" + strconv.Itoa(i+1)
		suite.API.Metadata.ServiceAbbreviation = svcPrefix + "ProtocolTest"
		suite.API.Operations = map[string]*api.Operation{}
		for idx, c := range suite.Cases {
			c.Given.ExportedName = svcPrefix + "TestCaseOperation" + strconv.Itoa(idx)
			suite.API.Operations[c.Given.ExportedName] = c.Given
		}

		suite.API.Setup()

		for n, s := range suite.API.Shapes {
			s.Rename(svcPrefix + "TestShape" + n)
		}

		svcCode := addImports(removeSigner(suite.API.ServiceGoCode()))
		if i == 0 {
			importMatch := reImportRemoval.FindStringSubmatch(svcCode)
			buf.WriteString(importMatch[0] + "\n\n")
			buf.WriteString(preamble + "\n\n")
		}
		svcCode = removeImports(svcCode)
		svcCode = strings.Replace(svcCode, "func New(", "func New"+suite.API.StructName()+"(", -1)

		buf.WriteString(svcCode + "\n\n")
		buf.WriteString(removeImports(suite.API.APIGoCode()) + "\n\n")
		innerBuf.WriteString(suite.TestSuite() + "\n")
	}

	return util.GoFmt(buf.String() + innerBuf.String())
}

func main() {
	fmt.Println(GenerateInputTestSuite(os.Args[1]))
}
