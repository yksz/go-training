package main

import (
	"fmt"
	"io"
	"strings"
)

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitedReader{reader: r, maxBytes: n}
}

type limitedReader struct {
	reader   io.Reader
	maxBytes int64
}

func (r *limitedReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if r.maxBytes <= 0 {
		return 0, io.EOF
	}
	len := int64(len(p))
	if r.maxBytes < len {
		len = r.maxBytes
	}
	n, err := r.reader.Read(p[:len])
	if err != nil {
		return 0, nil
	}
	r.maxBytes -= int64(n)
	return n, nil
}

func main() {
	r := LimitReader(strings.NewReader("123456789"), 5)
	for {
		b := make([]byte, 3)
		_, err := r.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Println(string(b))
	}

	// Output:
	// 123
	// 45
}
