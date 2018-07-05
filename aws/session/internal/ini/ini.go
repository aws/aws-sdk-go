package ini

import (
	"os"
)

// OpenFile takes a path to a given file, and will open  and parse
// that file.
func OpenFile(path string) (Sections, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ParseFile(f)
}

// ParseFile will parse the given file using the shared config
// visitor.
func ParseFile(f *os.File) (Sections, error) {
	tree, err := Parse(f)
	if err != nil {
		return nil, err
	}

	v := NewSharedConfigVisitor()
	if err = Walk(tree, v); err != nil {
		return nil, err
	}

	return v.Sections, nil
}
