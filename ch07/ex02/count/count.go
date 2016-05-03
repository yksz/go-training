package count

import "io"

type countingWriter struct {
	writer io.Writer
	count  *int64
}

func (w *countingWriter) Write(p []byte) (int, error) {
	n, err := w.writer.Write(p)
	if err != nil {
		return 0, err
	}
	*w.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c int64
	cw := &countingWriter{writer: w, count: &c}
	return cw, &c
}
