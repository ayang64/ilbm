package iff

import (
	"bufio"
	"io"
)

type IFF struct{}

type Decoder struct {
	r io.ByteScanner
}

func NewDecoder(r io.Reader) *Decoder {
	byteScanner := func(r io.Reader) io.ByteScanner {
		if bs, ok := r.(io.ByteScanner); ok {
			return bs
		}
		return bufio.NewReader(r)
	}

	rc := Decoder{
		r: byteScanner(r),
	}

	return &rc
}
