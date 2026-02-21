package bsdiff

import (
	"io"
)

type AdderReader struct {
	Buffer []byte
	Reader io.Reader

	offset int
}

var _ io.Reader = (*AdderReader)(nil)

func (ar *AdderReader) Read(p []byte) (int, error) {
	n, err := ar.Reader.Read(p)
	if err != nil {
		return n, err
	}

	b := ar.Buffer
	off := ar.offset

	for i := range n {
		p[i] += b[off+i]
	}
	ar.offset += n

	return n, nil
}
