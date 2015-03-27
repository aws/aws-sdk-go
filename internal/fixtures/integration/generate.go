package main

//go:generate go run generate.go service/*.json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/awslabs/aws-sdk-go/internal/fixtures/helpers"
	"github.com/awslabs/aws-sdk-go/internal/model/api"
	"github.com/awslabs/aws-sdk-go/internal/util"
)

type TestSuite struct {
	API         *api.API
	PackageName string
	APIVersion  string `json:"api_version"`
	Cases       []*TestCase
}

type TestCase struct {
	API         *api.API
	Description string
	Operation   string
	Input       interface{}
	Assertions  []*TestAssertion
}

type TestAssertion struct {
	Case      *TestCase
	Assertion string
	Context   string
	Path      string
	Expected  interface{}
}

var tplTestSuite = template.Must(template.New("testsuite").Parse(`
// +build integration

package {{ .API.PackageName }}_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/util/utilassert"
	"github.com/awslabs/aws-sdk-go/service/{{ .API.PackageName }}"
	"github.com/stretchr/testify/assert"
)

var (
	_ = assert.Equal
	_ = utilassert.Match
)

{{ range $_, $t := .Cases }}{{ $t.GoCode }}{{ end }}
`))

var tplTestCase = template.Must(template.New("testcase").Parse(`
func Test{{ .TestName }}(t *testing.T) {
	client := {{ .API.PackageName }}.New(nil)
	resp, e := client.{{ .API.ExportableName .Operation }}({{ .InputCode }})
	err := aws.Error(e)
	_, _, _ = resp, e, err // avoid unused warnings

	{{ range $_, $a := .Assertions }}{{ $a.GoCode }}{{ end }}
}
`))

func (t *TestSuite) setup() {
	_, d, _, _ := runtime.Caller(1)
	file := filepath.Join(path.Dir(d), "..", "..", "..", "apis",
		t.PackageName, t.APIVersion+".normal.json")

	t.API = &api.API{}
	t.API.Attach(file)

	for _, c := range t.Cases {
		c.API = t.API
		for _, a := range c.Assertions {
			a.Case = c
		}
	}
}

func (t *TestSuite) write() {
	_, d, _, _ := runtime.Caller(1)
	file := filepath.Join(path.Dir(d), "..", "..", "..", "service",
		t.PackageName, "integration_test.go")

	var buf bytes.Buffer
	if err := tplTestSuite.Execute(&buf, t); err != nil {
		panic(err)
	}

	b := []byte(util.GoFmt(buf.String()))
	ioutil.WriteFile(file, b, 0644)
}

func (t *TestCase) TestName() string {
	out := ""
	for _, v := range strings.Split(t.Description, " ") {
		out += util.Capitalize(v)
	}
	return out
}

func (t *TestCase) GoCode() string {
	var buf bytes.Buffer
	if err := tplTestCase.Execute(&buf, t); err != nil {
		panic(err)
	}
	return util.GoFmt(buf.String())
}

func (t *TestCase) InputCode() string {
	op := t.API.Operations[t.API.ExportableName(t.Operation)]
	if op.InputRef.Shape == nil {
		return ""
	}
	return helpers.ParamsStructFromJSON(t.Input, op.InputRef.Shape, true)
}

func (t *TestAssertion) GoCode() string {
	call, actual, expected := "", "", fmt.Sprintf("%#v", t.Expected)

	if expected == "<nil>" {
		expected = "nil"
	}

	switch t.Context {
	case "error":
		actual = "err"
	case "data":
		actual = "resp"
	default:
		panic("unsupported assertion context " + t.Context)
	}

	if t.Path != "" {
		actual += "." + util.Capitalize(t.Path)
	}

	switch t.Assertion {
	case "typeof":
		return "" // do nothing for typeof checks
	case "equal":
		if actual == "err" && expected == "nil" {
			call = "assert.NoError"
		} else {
			call = "assert.Equal"
		}
	case "notequal":
		call = "assert.NotEqual"
	case "contains":
		call = "utilassert.Match"
	default:
		panic("unsupported assertion type " + t.Assertion)
	}

	return fmt.Sprintf("%s(t, %s, %s)\n", call, expected, actual)
}

func GenerateIntegrationSuite(testFile string) {
	pkgName := strings.Replace(filepath.Base(testFile), ".json", "", -1)
	suite := &TestSuite{PackageName: pkgName}

	if file, err := os.Open(testFile); err == nil {
		defer file.Close()
		if err = json.NewDecoder(file).Decode(&suite); err != nil {
			panic(err)
		}

		suite.setup()
		suite.write()
	} else {
		panic(err)
	}
}

func main() {
	files := []string{}
	for _, arg := range os.Args[1:] {
		paths, _ := filepath.Glob(arg)
		files = append(files, paths...)
	}

	for _, file := range files {
		GenerateIntegrationSuite(file)
	}
}
