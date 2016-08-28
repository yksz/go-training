// Package bzip provides a writer that uses bzip2 compression (bzip.org).
package bzip

import (
	"bytes"
	"io"
	"os/exec"
)

type writer struct {
	w io.Writer // underlying output stream
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) io.WriteCloser {
	return &writer{w: out}
}

func (w *writer) Write(data []byte) (int, error) {
	cmd := exec.Command("bzip2", "-cz")
	cmd.Stdin = bytes.NewBuffer(data)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return 0, err
	}
	if _, err := w.w.Write(out.Bytes()); err != nil {
		return 0, err
	}
	return len(data), nil
}

// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer.
func (w *writer) Close() error {
	return nil
}
