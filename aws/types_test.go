package aws

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWriteAtBuffer(t *testing.T) {
	b := &WriteAtBuffer{}

	n, err := b.WriteAt([]byte{1}, 0)
	assert.NoError(t, err)
	assert.Equal(t, 1, n)

	n, err = b.WriteAt([]byte{2}, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, n)

	n, err = b.WriteAt([]byte{3}, 2)
	assert.NoError(t, err)
	assert.Equal(t, 1, n)

	n, err = b.WriteAt([]byte{1,1,1}, 5)
	assert.NoError(t, err)
	assert.Equal(t, 3, n)

	assert.Equal(t, []byte{1,2,3,0,0,1,1,1}, b.Bytes())
}
