package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"text/template"
)

type Waiter struct {
	Name          string
	Delay         int
	MaxAttempts   int
	OperationName string `json:"operation"`
	Operation     *Operation
	Acceptors     []WaitAcceptor
}

type WaitAcceptor struct {
	Expected interface{}
	Matcher  string
	State    string
	Argument string
}

func (a *API) WaitersGoCode() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "import (\n\t%q\n)",
		"github.com/aws/aws-sdk-go/internal/waiter")

	for _, w := range a.Waiters {
		buf.WriteString(w.GoCode())
	}
	return buf.String()
}

// used for unmarshaling from the waiter JSON file
type waiterDefinitions struct {
	*API
	Waiters map[string]Waiter
}

func (a *API) AttachWaiters(filename string) {
	p := waiterDefinitions{API: a}

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(f).Decode(&p)
	if err != nil {
		panic(err)
	}

	p.setup()
}

func (p *waiterDefinitions) setup() {
	p.API.Waiters = []Waiter{}
	i, keys := 0, make([]string, len(p.Waiters))
	for k, _ := range p.Waiters {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for _, n := range keys {
		e := p.Waiters[n]
		n = p.ExportableName(n)
		e.Name = n
		e.OperationName = p.ExportableName(e.OperationName)
		e.Operation = p.API.Operations[e.OperationName]
		if e.Operation == nil {
			panic("unknown operation " + e.OperationName + " for waiter " + n)
		}
		p.API.Waiters = append(p.API.Waiters, e)
	}
}

func (a *WaitAcceptor) ExpectedString() string {
	switch a.Expected.(type) {
	case string:
		return fmt.Sprintf("%q", a.Expected)
	default:
		return fmt.Sprintf("%v", a.Expected)
	}
}

func (w *Waiter) GoCode() string {
	var buf bytes.Buffer
	if err := tplWaiter.Execute(&buf, w); err != nil {
		panic(err)
	}

	return buf.String()
}

var tplWaiter = template.Must(template.New("waiter").Parse(`
var waiter{{ .Name }} *waiter.Config

func (c *{{ .Operation.API.StructName }}) WaitUntil{{ .Name }}(input {{ .Operation.InputRef.GoType }}) error {
	if waiter{{ .Name }} == nil {
		waiter{{ .Name }} = &waiter.Config{
			Operation:   "{{ .OperationName }}",
			Delay:       {{ .Delay }},
			MaxAttempts: {{ .MaxAttempts }},
			Acceptors: []waiter.WaitAcceptor{
				{{ range $_, $a := .Acceptors }}waiter.WaitAcceptor{
					State:    "{{ .State }}",
					Matcher:  "{{ .Matcher }}",
					Argument: "{{ .Argument }}",
					Expected: {{ .ExpectedString }},
				},
				{{ end }}
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiter{{ .Name }},
	}
	return w.Wait()
}
`))
