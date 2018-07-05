package ini

import (
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// Visitor is an interface used by walkers that will
// traverse an array of ASTs.
type Visitor interface {
	VisitExpr(AST) error
	VisitStatement(AST) error
}

type Sections map[string]Section

type SharedConfigVisitor struct {
	scope    string
	Sections Sections
}

func NewSharedConfigVisitor() *SharedConfigVisitor {
	return &SharedConfigVisitor{
		Sections: Sections{},
	}
}

func (t Sections) GetSection(p string) (Section, bool) {
	sections := map[string]Section(t)
	v, ok := sections[p]
	return v, ok
}

type Values map[string]iniToken

func (t Sections) List() []string {
	keys := make([]string, len(t))
	i := 0
	for k := range t {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	return keys
}

type Section struct {
	Name   string
	Values Values
}

func (t Section) Int(k string) int64 {
	return t.Values[k].IntValue()
}

func (t Section) Float64(k string) float64 {
	return t.Values[k].FloatValue()
}

func (t Section) String(k string) string {
	_, ok := t.Values[k]
	if !ok {
		return ""
	}
	return t.Values[k].StringValue()
}

func (v *SharedConfigVisitor) VisitExpr(expr AST) error {
	t := v.Sections[v.scope]
	if t.Values == nil {
		t.Values = Values{}
	}

	switch e := expr.(type) {
	case ExprStatement:
		switch opExpr := e.V.(type) {
		case EqualExpr:
			t.Values[opExpr.Key()] = opExpr.Right.(Expr).Root
		}
	default:
		return awserr.New(ErrCodeParseError, "unsupported expression", nil)
	}

	v.Sections[v.scope] = t
	return nil
}

func (v *SharedConfigVisitor) VisitStatement(stmt AST) error {
	switch s := stmt.(type) {
	case SectionStatement:
		v.Sections[s.Name] = Section{}
		v.scope = s.Name
	default:
		return awserr.New(ErrCodeParseError, fmt.Sprintf("unsupported statement: %T", s), nil)
	}

	return nil
}
