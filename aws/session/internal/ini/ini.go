package ini

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
)

// OpenFile takes a path to a given file, and will open  and parse
// that file.
func OpenFile(path string) (Sections, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, awserr.New(ErrCodeUnableToReadFile, "unable to open file", err)
	}
	defer f.Close()

	return Parse(f)
}

// Parse will parse the given file using the shared config
// visitor.
func Parse(f io.Reader) (Sections, error) {
	tree, err := ParseAST(f)
	if err != nil {
		return nil, err
	}

	v := NewSharedConfigVisitor()
	if err = Walk(tree, v); err != nil {
		return nil, err
	}

	return v.Sections, nil
}
