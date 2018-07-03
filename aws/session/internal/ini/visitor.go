package ini

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

type Visitor interface {
	VisitExpr(AST) error
	VisitStatement(AST) error
}

type Tables map[string]Table

type SharedConfigVisitor struct {
	scope  string
	Tables Tables
}

func NewSharedConfigVisitor() *SharedConfigVisitor {
	return &SharedConfigVisitor{
		Tables: Tables{},
	}
}

func (t Tables) GetSection(p string) (Table, bool) {
	tables := map[string]Table(t)
	v, ok := tables[p]
	return v, ok
}

type Values map[string]iniToken

type Table struct {
	Name   string
	Values Values
}

func (t Table) Int(k string) int {
	return t.Values[k].IntValue()
}

func (t Table) Float64(k string) float64 {
	return t.Values[k].FloatValue()
}

func (t Table) String(k string) string {
	_, ok := t.Values[k]
	if !ok {
		return ""
	}
	return t.Values[k].StringValue()
}

func (v *SharedConfigVisitor) VisitExpr(expr AST) error {
	t := v.Tables[v.scope]
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

	v.Tables[v.scope] = t
	return nil
}

func (v *SharedConfigVisitor) VisitStatement(stmt AST) error {
	switch s := stmt.(type) {
	case TableStatement:
		v.Tables[s.Name] = Table{}
		v.scope = s.Name
	default:
		return awserr.New(ErrCodeParseError, fmt.Sprintf("unsupported statement: %T", s), nil)
	}

	return nil
}
