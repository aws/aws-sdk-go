package s3crypto

import ()

const bufSize = 2048

type buffer struct {
	data []byte
}

// appendToBuffer takes a byte array that we will add to the end of the reader's
// buffer and will change the size of the reader as fit
// May have issues with padding if pLen isn't blocksize or more
func (buf *buffer) appendToBuffer(data *[]byte, b []byte) int {
	buf.data = append(buf.data, b...)

	max := len(b)
	dLen := len(*data)
	if max > dLen {
		max = dLen
	}
	copy(*data, buf.data[:max])

	// shift bytes
	buf.data = buf.data[max:]

	return max
}

func (buf *buffer) drainBody(data *[]byte, blockSize int) int {
	cutoff, dLen := len(buf.data)-blockSize, len(*data)
	if cutoff > dLen {
		cutoff = dLen
	}

	copy(*data, buf.data[:cutoff])
	buf.data = buf.data[cutoff:]
	return cutoff
}

// Called when nothing has been read or io.EOF was returned. Meaning we are at the end
// of the reads
func (buf *buffer) finalize(lastBlock bool, dst *[]byte, data *[]byte, blockSize int) int {
	*dst = append(buf.data, (*dst)...)
	cLen, dLen := len(*dst), len(*data)
	if cLen <= dLen {
		copy(*data, *dst)
		return cLen
	}

	copy(*data, (*dst)[:dLen])
	*dst = (*dst)[dLen:]
	cLen = len(*dst)
	buf.data = (*data)
	return dLen
}
