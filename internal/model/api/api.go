package api

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/awslabs/aws-sdk-go/internal/util"
)

type API struct {
	Metadata   Metadata
	Operations map[string]*Operation
	Shapes     map[string]*Shape

	// Disables inflection checks. Only use this when generating tests
	NoInflections bool

	// Set to true to avoid removing unused shapes
	NoRemoveUnusedShapes bool

	// Set to true to ignore service/request init methods (for testing)
	NoInitMethods bool

	initialized       bool
	imports           map[string]bool
	name              string
	unrecognizedNames map[string]string
}

type Metadata struct {
	APIVersion          string
	EndpointPrefix      string
	ServiceAbbreviation string
	ServiceFullName     string
	SignatureVersion    string
	JSONVersion         string
	TargetPrefix        string
	Protocol            string
}

func (a *API) PackageName() string {
	return strings.ToLower(a.StructName())
}

var nameRegex = regexp.MustCompile(`^Amazon|AWS\s*|\(.*|\s+|\W+`)

// StructName returns the service struct name for a given service
func (a *API) StructName() string {
	if a.name == "" {
		name := a.Metadata.ServiceAbbreviation
		if name == "" {
			name = a.Metadata.ServiceFullName
		}

		name = nameRegex.ReplaceAllString(name, "")
		switch name {
		case "ElasticLoadBalancing":
			a.name = "ELB"
		case "Config":
			a.name = "ConfigService"
		default:
			a.name = name
		}
	}
	return a.name
}

func (a *API) UseInitMethods() bool {
	return !a.NoInitMethods
}

func (a *API) NiceName() string {
	if a.Metadata.ServiceAbbreviation != "" {
		return a.Metadata.ServiceAbbreviation
	}
	return a.Metadata.ServiceFullName
}

func (a *API) ProtocolPackage() string {
	switch a.Metadata.Protocol {
	case "json":
		return "jsonrpc"
	case "ec2":
		return "ec2query"
	default:
		return strings.Replace(a.Metadata.Protocol, "-", "", -1)
	}
}

func (a *API) OperationNames() []string {
	i, names := 0, make([]string, len(a.Operations))
	for n, _ := range a.Operations {
		names[i] = n
		i++
	}
	sort.Strings(names)
	return names
}

func (a *API) OperationList() []*Operation {
	list := make([]*Operation, len(a.Operations))
	for i, n := range a.OperationNames() {
		list[i] = a.Operations[n]
	}
	return list
}

func (a *API) ShapeNames() []string {
	i, names := 0, make([]string, len(a.Shapes))
	for n, _ := range a.Shapes {
		names[i] = n
		i++
	}
	sort.Strings(names)
	return names
}

func (a *API) ShapeList() []*Shape {
	list := make([]*Shape, len(a.Shapes))
	for i, n := range a.ShapeNames() {
		list[i] = a.Shapes[n]
	}
	return list
}

func (a *API) resetImports() {
	a.imports = map[string]bool{
		"github.com/awslabs/aws-sdk-go/aws": true,
	}
}

func (a *API) importsGoCode() string {
	if len(a.imports) == 0 {
		return ""
	}

	corePkgs, extPkgs := []string{}, []string{}
	for i, _ := range a.imports {
		if strings.Contains(i, ".") {
			extPkgs = append(extPkgs, i)
		} else {
			corePkgs = append(corePkgs, i)
		}
	}
	sort.Strings(corePkgs)
	sort.Strings(extPkgs)

	code := "import (\n"
	for _, i := range corePkgs {
		code += fmt.Sprintf("\t%q\n", i)
	}
	code += "\n"
	for _, i := range extPkgs {
		code += fmt.Sprintf("\t%q\n", i)
	}
	code += ")\n\n"
	return code
}

var tplAPI = template.Must(template.New("api").Parse(`
var oprw sync.Mutex

{{ range $_, $o := .OperationList }}
{{ $o.GoCode }}

{{ end }}

{{ range $_, $s := .ShapeList }}
{{ if eq $s.Type "structure" }}{{ $s.GoCode }}{{ end }}

{{ end }}
`))

func (a *API) APIGoCode() string {
	a.resetImports()
	a.imports["sync"] = true
	var buf bytes.Buffer
	err := tplAPI.Execute(&buf, a)
	if err != nil {
		panic(err)
	}

	code := a.importsGoCode() + strings.TrimSpace(buf.String())
	return util.GoFmt(code)
}

var tplService = template.Must(template.New("service").Parse(`
// {{ .StructName }} is a client for {{ .NiceName }}.
type {{ .StructName }} struct {
	*aws.Service
}

{{ if .UseInitMethods }}// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)
{{ end }}

// New returns a new {{ .StructName }} client.
func New(config *aws.Config) *{{ .StructName }} {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "{{ .Metadata.EndpointPrefix }}",
		APIVersion:   "{{ .Metadata.APIVersion }}",
{{ if eq .Metadata.Protocol "json" }}JSONVersion:  "{{ .Metadata.JSONVersion }}",
		TargetPrefix: "{{ .Metadata.TargetPrefix }}",
{{ end }}
  }
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack({{ .ProtocolPackage }}.Build)
	service.Handlers.Unmarshal.PushBack({{ .ProtocolPackage }}.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack({{ .ProtocolPackage }}.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack({{ .ProtocolPackage }}.UnmarshalError)

	{{ if .UseInitMethods }}// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}
	{{ end  }}

	return &{{ .StructName }}{service}
}

// newRequest creates a new request for a {{ .StructName }} operation and runs any
// custom request initialization.
func (c *{{ .StructName }}) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	{{ if .UseInitMethods }}// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}
	{{ end }}

	return req
}
`))

func (a *API) ServiceGoCode() string {
	a.resetImports()
	a.imports["github.com/awslabs/aws-sdk-go/internal/signer/v4"] = true
	a.imports["github.com/awslabs/aws-sdk-go/internal/protocol/"+a.ProtocolPackage()] = true

	var buf bytes.Buffer
	err := tplService.Execute(&buf, a)
	if err != nil {
		panic(err)
	}

	code := a.importsGoCode() + buf.String()
	return util.GoFmt(code)
}

func (a *API) ExampleGoCode() string {
	exs := []string{}
	for _, o := range a.OperationList() {
		exs = append(exs, o.Example())
	}

	code := fmt.Sprintf("import (\n%q\n%q\n%q\n\n%q\n%q\n%q\n)\n\n"+
		"var _ time.Duration\nvar _ bytes.Buffer\n\n%s",
		"bytes",
		"fmt",
		"time",
		"github.com/awslabs/aws-sdk-go/aws",
		"github.com/awslabs/aws-sdk-go/aws/awsutil",
		"github.com/awslabs/aws-sdk-go/service/"+a.PackageName(),
		strings.Join(exs, "\n\n"),
	)
	return util.GoFmt(code)
}
