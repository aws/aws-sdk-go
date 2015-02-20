package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/awslabs/aws-sdk-go/internal/model/api"
)

type InputTestSuite struct {
	api.API
	Description string
	Cases       []InputTestCase
}

type InputTestCase struct {
	Given      api.Operation
	Params     interface{}
	Serialized InputTestExpectation
}

type InputTestExpectation struct {
	Body    string
	URI     string
	Headers map[string]string
}

func (i *InputTestSuite) GoCode() (code string) {
	i.API.Setup()
	prefix := reStripSpace.ReplaceAllStringFunc(i.Description, func(x string) string {
		return strings.ToUpper(x)
	})
	for _, c := range i.Cases {
		c.GoCode(prefix)
	}
	//code = fmt.Sprintf("%#v\n", suite)
	return
}

func (i *InputTestCase) GoCode(name string) (code string) {
	code = "func Test" + strings.ToUpper(name[0:1]) + name[1:] + "(t *testing.T) {\n"
	code += "if req.HttpRequest.Body == " + fmt.Sprintf("%q", i.Serialized.Body) + " {\n"
	code += "}"
	code += "}\n"
	return
}

var reStripSpace = regexp.MustCompile(`\s(\w)`)

func GenerateInputTestSuite(filename string) {
	var suites []InputTestSuite
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	json.NewDecoder(f).Decode(&suites)

	for _, suite := range suites {
		fmt.Println(suite.GoCode())
	}
}

func main() {
	GenerateInputTestSuite(os.Args[1])
}
