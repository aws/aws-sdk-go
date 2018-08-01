package ini

import (
	"reflect"
	"testing"
)

type mockAST struct {
	Value int
}

func (ast mockAST) Kind() ASTKind {
	return ASTKindNone
}

func TestStack(t *testing.T) {
	cases := []struct {
		asts     []AST
		expected []AST
	}{
		{
			asts: []AST{
				mockAST{0},
				mockAST{1},
				mockAST{2},
				mockAST{3},
				mockAST{4},
			},
			expected: []AST{
				mockAST{0},
				mockAST{1},
				mockAST{2},
				mockAST{3},
				mockAST{4},
			},
		},
	}

	for _, c := range cases {
		p := newParseStack(10, 10)
		for _, ast := range c.asts {
			p.Push(ast)
			p.MarkComplete(ast)
		}

		if e, a := len(c.expected), p.Len(); e != a {
			t.Errorf("expected the same legnth with %d, but received %d", e, a)
		}
		for i := len(c.expected) - 1; i >= 0; i-- {
			if e, a := c.expected[i], p.Pop(); e != a {
				t.Errorf("stack element %d invalid: expected %v, but received %v", i, e, a)
			}
		}

		if e, a := len(c.expected), p.index; e != a {
			t.Errorf("expected %d, but received %d", e, a)
		}

		if e, a := c.asts, p.list[:p.index]; !reflect.DeepEqual(e, a) {
			t.Errorf("expected %v, but received %v", e, a)
		}
	}
}
