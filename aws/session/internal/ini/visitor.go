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

// Sections is a map of Section structures that represent
// a configuration.
type Sections map[string]Section

// SharedConfigVisitor is used to visit statements and expressions
// and ensure that they are both of the correct format.
// In addition, upon visiting this will build sections and populate
// the Sections field which can be used to retrieve profile
// configuration.
type SharedConfigVisitor struct {
	scope    string
	Sections Sections
}

// NewSharedConfigVisitor return a SharedConfigVisitor
func NewSharedConfigVisitor() *SharedConfigVisitor {
	return &SharedConfigVisitor{
		Sections: Sections{},
	}
}

// GetSection will return section p. If section p does not exist,
// false will be returned in the second parameter.
func (t Sections) GetSection(p string) (Section, bool) {
	sections := map[string]Section(t)
	v, ok := sections[p]
	return v, ok
}

// Values represents a map of tokens.
type Values map[string]Token

// List will return a list of all sections that were successfully
// parsed.
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

// Section contains a name and values. This represent
// a sectioned entry in a configuration file.
type Section struct {
	Name   string
	Values Values
}

// ValueType will returned what type the union is set to. If
// k was not found, the NoneType will be returned.
func (t Section) ValueType(k string) UnionValueType {
	v, ok := t.Values[k].(literalToken)
	if !ok {
		return NoneType
	}
	return v.Value.Type
}

// Int returns an integer value at k
func (t Section) Int(k string) int64 {
	return t.Values[k].IntValue()
}

// Float64 returns a float value at k
func (t Section) Float64(k string) float64 {
	return t.Values[k].FloatValue()
}

// String returns the string value at k
func (t Section) String(k string) string {
	_, ok := t.Values[k]
	if !ok {
		return ""
	}
	return t.Values[k].StringValue()
}

// VisitExpr visits expressions...
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

// VisitStatement visits statements...
func (v *SharedConfigVisitor) VisitStatement(stmt AST) error {
	switch s := stmt.(type) {
	case CompletedSectionStatement:
		child, ok := s.V.(SectionStatement)
		if !ok {
			return awserr.New(ErrCodeParseError, fmt.Sprintf("unsupported child statement: %T", child), nil)
		}
		v.Sections[child.Name] = Section{}
		v.scope = child.Name
	default:
		return awserr.New(ErrCodeParseError, fmt.Sprintf("unsupported statement: %T", s), nil)
	}

	return nil
}
