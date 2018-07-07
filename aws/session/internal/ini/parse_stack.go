package ini

import (
	"bytes"
	"fmt"
)

// ParseStack is a stack that contains a container, the stack portion,
// and the list which is the list of ASTs that have been successfully
// parsed.
type ParseStack struct {
	container []AST
	list      []AST
}

// Pop will return and truncate the last container element.
func (s *ParseStack) Pop() AST {
	temp := s.container[s.Len()-1]
	s.container = s.container[:s.Len()-1]
	return temp
}

// Push will add the new AST to the container
func (s *ParseStack) Push(ast AST) {
	s.container = append(s.container, ast)
}

// Epsilon will append the AST to the list of completed statements
func (s *ParseStack) Epsilon(ast AST) {
	s.list = append(s.list, ast)
}

// Len will return the length of the container
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
