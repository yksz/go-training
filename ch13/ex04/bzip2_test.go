package bzip

import (
	"bytes"
	"reflect"
	"testing"

	"gopl.io/ch13/bzip"
)

func TestBzip2(t *testing.T) {
	data := []byte("hello")

	var want []byte
	{
		var out bytes.Buffer
		w := bzip.NewWriter(&out)
		w.Write(data)
		w.Close()
		want = out.Bytes()
	}

	var got []byte
	{
		var out bytes.Buffer
		w := NewWriter(&out)
		w.Write(data)
		w.Close()
		got = out.Bytes()
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %d bytes, want: %d byets", len(got), len(want))
	}
}
