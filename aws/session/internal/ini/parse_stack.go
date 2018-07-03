package ini

import (
	"bytes"
	"fmt"
)

type ParseStack struct {
	container []AST
	list      []AST
}

func (s *ParseStack) Pop() AST {
	temp := s.container[0]
	s.container = s.container[1:]
	return temp
}

func (s *ParseStack) Push(ast AST) {
	s.container = append(s.container, ast)
}

func (s *ParseStack) Epsilon(ast AST) {
	s.Push(Start{})
	s.list = append(s.list, ast)
}

func (s *ParseStack) Len() int {
	return len(s.container)
}

func (s ParseStack) String() string {
	buf := bytes.Buffer{}
	for i, node := range s.list {
		buf.WriteString(fmt.Sprintf("%d: %v\n", i+1, node))
	}

	return buf.String()
}
