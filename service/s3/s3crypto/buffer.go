package s3crypto

import ()

const bufSize = 2048

type buffer struct {
	data [bufSize]byte
	size int
}

// appendToBuffer takes a byte array that we will add to the end of the reader's
// buffer and will change the size of the reader as fit
// May have issues with padding if pLen isn't blocksize or more
func (buf *buffer) appendToBuffer(data *[]byte, b []byte) int {
	pLen := len(b)
	rem := bufSize - buf.size
	if bLen := len(b); rem > bLen {
		rem = bLen
	}

	for i := 0; i < rem; i++ {
		buf.data[buf.size+i] = b[i]
	}
	buf.size += rem
	b = b[rem:]
	copy(*data, buf.data[:buf.size])
	tmpSize := buf.size

	buf.size = pLen - rem
	for i := 0; i < buf.size; i++ {
		buf.data[i] = b[i]
	}
	return tmpSize
}

func (buf *buffer) drainBody(data *[]byte, blockSize int) int {
	cutoff, dLen := buf.size-blockSize, len(*data)
	if cutoff > dLen {
		cutoff = dLen
	}

	copy(*data, buf.data[:cutoff])
	for i := 0; cutoff+i < buf.size; i++ {
		buf.data[i] = buf.data[cutoff+i]
	}
	buf.size -= cutoff
	return cutoff
}

// Called when nothing has been read or io.EOF was returned. Meaning we are at the end
// of the reads
func (buf *buffer) finalize(lastBlock bool, dst *[]byte, data *[]byte, blockSize int) int {
	*dst = append(buf.data[:buf.size], (*dst)...)
	cLen, dLen := len(*dst), len(*data)
	if cLen <= dLen {
		copy(*data, *dst)
		buf.size = 0
		return cLen
	}

	copy(*data, (*dst)[:dLen])
	*dst = (*dst)[dLen:]
	cLen = len(*dst)
	for i := 0; i < cLen; i++ {
		buf.data[i] = (*dst)[i]
	}
	buf.size = cLen
	return dLen
}
