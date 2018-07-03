package ini

import (
	"os"
)

func OpenFile(path string) (Tables, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ParseFile(f)
}

func ParseFile(f *os.File) (Tables, error) {
	tree, err := Parse(f)
	if err != nil {
		return nil, err
	}

	v := NewSharedConfigVisitor()
	if err = Walk(tree, v); err != nil {
		return nil, err
	}

	return v.Tables, nil
}
