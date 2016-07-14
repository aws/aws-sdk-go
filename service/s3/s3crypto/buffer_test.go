package s3crypto

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendToBufferLessSize(t *testing.T) {
	b := buffer{}
	data := make([]byte, 128)
	// extra is going to be copied into data
	extra := bytes.Repeat([]byte{'a'}, 128)
	size := b.appendToBuffer(&data, extra)

	assert.Equal(t, 128, size)
	assert.Equal(t, extra, data)
}

func TestAppendToBufferSameSize(t *testing.T) {
	b := buffer{}
	data := make([]byte, bufSize)
	// extra is going to be copied into data
	extra := bytes.Repeat([]byte{'a'}, bufSize)
	size := b.appendToBuffer(&data, extra)

	assert.Equal(t, bufSize, size)
	assert.Equal(t, extra, data)
}

func TestAppendToBufferGreaterSize(t *testing.T) {
	b := buffer{}
	data := make([]byte, bufSize*2)
	appendData := bytes.Repeat([]byte{'a'}, bufSize*2)
	size := b.appendToBuffer(&data, appendData)

	assert.Equal(t, bufSize*2, size)
	assert.Equal(t, appendData[:bufSize*2], data[:size])
	assert.Equal(t, 0, len(b.data))
}
