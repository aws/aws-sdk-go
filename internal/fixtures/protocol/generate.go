package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/awslabs/aws-sdk-go/internal/fixtures/helpers"
	"github.com/awslabs/aws-sdk-go/internal/model/api"
	"github.com/awslabs/aws-sdk-go/internal/util"
	"github.com/awslabs/aws-sdk-go/internal/util/utilassert"
)

type TestSuite struct {
	*api.API
	Description string
	Cases       []TestCase
	title       string
}

type TestCase struct {
	*TestSuite
	Given      *api.Operation
	Params     interface{}     `json:",omitempty"`
	Data       interface{}     `json:"result,omitempty"`
	InputTest  TestExpectation `json:"serialized"`
	OutputTest TestExpectation `json:"response"`
}

type TestExpectation struct {
	Body       string
	URI        string
	Headers    map[string]string
	StatusCode uint `json:"status_code"`
}

const preamble = `
var _ bytes.Buffer // always import bytes
var _ http.Request
var _ json.Marshaler
var _ time.Time
var _ xmlutil.XMLNode
var _ xml.Attr
var _ = ioutil.Discard
var _ = util.Trim("")
var _ = url.Values{}
var _ = io.EOF
`

var reStripSpace = regexp.MustCompile(`\s(\w)`)

var reImportRemoval = regexp.MustCompile(`(?s:import \((.+?)\))`)

func removeImports(code string) string {
	return reImportRemoval.ReplaceAllString(code, "")
}

var extraImports = []string{
	"bytes",
	"encoding/json",
	"encoding/xml",
	"io",
	"io/ioutil",
	"net/http",
	"testing",
	"time",
	"net/url",
	"github.com/awslabs/aws-sdk-go/internal/protocol/xml/xmlutil",
	"github.com/awslabs/aws-sdk-go/internal/util",
	"github.com/stretchr/testify/assert",
}

func addImports(code string) string {
	importNames := make([]string, len(extraImports))
	for i, n := range extraImports {
		importNames[i] = fmt.Sprintf("%q", n)
	}
	str := reImportRemoval.ReplaceAllString(code, "import (\n$1\n"+strings.Join(importNames, "\n")+")")
	return str
}

func (t *TestSuite) TestSuite() string {
	var buf bytes.Buffer

	t.title = reStripSpace.ReplaceAllStringFunc(t.Description, func(x string) string {
		return strings.ToUpper(x[1:])
	})
	t.title = regexp.MustCompile(`\W`).ReplaceAllString(t.title, "")

	for idx, c := range t.Cases {
		c.TestSuite = t
		buf.WriteString(c.TestCase(idx) + "\n")
	}
	return util.GoFmt(buf.String())
}

var tplInputTestCase = template.Must(template.New("inputcase").Parse(`
func Test{{ .OpName }}(t *testing.T) {
	svc := New{{ .TestCase.TestSuite.API.StructName }}(nil)
	svc.Endpoint = "https://test"

	input := {{ .ParamsString }}
	req, _ := svc.{{ .Given.ExportedName }}Request(input)
	r := req.HTTPRequest

	// build request
	{{ .TestCase.TestSuite.API.ProtocolPackage }}.Build(req)
	assert.NoError(t, req.Error)

	{{ if ne .TestCase.InputTest.Body "" }}// assert body
	assert.NotNil(t, r.Body)
	{{ .BodyAssertions }}{{ end }}

	{{ if ne .TestCase.InputTest.URI "" }}// assert URL
	assert.Equal(t, "https://test{{ .TestCase.InputTest.URI }}", r.URL.String()){{ end }}

	// assert headers
{{ range $k, $v := .TestCase.InputTest.Headers }}assert.Equal(t, "{{ $v }}", r.Header.Get("{{ $k }}"))
{{ end }}
}
`))

type tplInputTestCaseData struct {
	*TestCase
	OpName, ParamsString string
}

func (t tplInputTestCaseData) BodyAssertions() string {
	protocol, code := t.TestCase.TestSuite.API.Metadata.Protocol, ""
	switch protocol {
	case "rest-xml":
		code += "body := util.SortXML(r.Body)\n"
	default:
		code += "body, _ := ioutil.ReadAll(r.Body)\n"
	}

	code += "assert.Equal(t, util.Trim(`" + t.InputTest.Body + "`), util.Trim(string(body)))"
	return code
}

var tplOutputTestCase = template.Must(template.New("outputcase").Parse(`
func Test{{ .OpName }}(t *testing.T) {
	svc := New{{ .TestCase.TestSuite.API.StructName }}(nil)

	buf := bytes.NewReader([]byte({{ .Body }}))
	req, out := svc.{{ .Given.ExportedName }}Request(nil)
	req.HTTPResponse = &http.Response{StatusCode: 200, Body: ioutil.NopCloser(buf), Header: http.Header{}}

	// set headers
	{{ range $k, $v := .TestCase.OutputTest.Headers }}req.HTTPResponse.Header.Set("{{ $k }}", "{{ $v }}")
	{{ end }}

	// unmarshal response
	{{ .TestCase.TestSuite.API.ProtocolPackage }}.UnmarshalMeta(req)
	{{ .TestCase.TestSuite.API.ProtocolPackage }}.Unmarshal(req)
	assert.NoError(t, req.Error)

	// assert response
	assert.NotNil(t, out) // ensure out variable is used
	{{ .Assertions }}
}
`))

type tplOutputTestCaseData struct {
	*TestCase
	Body, OpName, Assertions string
}

func (i *TestCase) TestCase(idx int) string {
	var buf bytes.Buffer

	opName := i.API.StructName() + i.TestSuite.title + "Case" + strconv.Itoa(idx+1)

	if i.Params != nil { // input test
		// query test should sort body as form encoded values
		switch i.API.Metadata.Protocol {
		case "query", "ec2":
			m, _ := url.ParseQuery(i.InputTest.Body)
			i.InputTest.Body = m.Encode()
		case "rest-xml":
			i.InputTest.Body = util.SortXML(bytes.NewReader([]byte(i.InputTest.Body)))
		case "json", "rest-json":
			i.InputTest.Body = strings.Replace(i.InputTest.Body, " ", "", -1)
		}

		input := tplInputTestCaseData{
			TestCase:     i,
			OpName:       strings.ToUpper(opName[0:1]) + opName[1:],
			ParamsString: helpers.ParamsStructFromJSON(i.Params, i.Given.InputRef.Shape, false),
		}

		if err := tplInputTestCase.Execute(&buf, input); err != nil {
			panic(err)
		}
	} else {
		output := tplOutputTestCaseData{
			TestCase:   i,
			Body:       fmt.Sprintf("%q", i.OutputTest.Body),
			OpName:     strings.ToUpper(opName[0:1]) + opName[1:],
			Assertions: utilassert.GenerateAssertions(i.Data, i.Given.OutputRef.Shape, "out"),
		}

		if err := tplOutputTestCase.Execute(&buf, output); err != nil {
			panic(err)
		}
	}

	return util.GoFmt(buf.String())
}

func GenerateTestSuite(filename string) string {
	inout := "Input"
	if strings.Contains(filename, "output/") {
		inout = "Output"
	}

	var suites []TestSuite
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(f).Decode(&suites)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	buf.WriteString("package " + suites[0].ProtocolPackage() + "_test\n\n")

	var innerBuf bytes.Buffer
	innerBuf.WriteString("//\n// Tests begin here\n//\n\n\n")

	for i, suite := range suites {
		svcPrefix := inout + "Service" + strconv.Itoa(i+1)
		suite.API.Metadata.ServiceAbbreviation = svcPrefix + "ProtocolTest"
		suite.API.Operations = map[string]*api.Operation{}
		for idx, c := range suite.Cases {
			c.Given.ExportedName = svcPrefix + "TestCaseOperation" + strconv.Itoa(idx+1)
			suite.API.Operations[c.Given.ExportedName] = c.Given
		}

		suite.API.NoInflections = true // don't require inflections
		suite.API.NoInitMethods = true // don't generate init methods
		suite.API.Setup()
		suite.API.Metadata.EndpointPrefix = suite.API.PackageName()

		for n, s := range suite.API.Shapes {
			s.Rename(svcPrefix + "TestShape" + n)
		}

		svcCode := addImports(suite.API.ServiceGoCode())
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
	out := GenerateTestSuite(os.Args[1])
	if len(os.Args) == 3 {
		f, err := os.Create(os.Args[2])
		defer f.Close()
		if err != nil {
			panic(err)
		}
		f.WriteString(out + "\n")
	} else {
		fmt.Println(out)
	}
}
