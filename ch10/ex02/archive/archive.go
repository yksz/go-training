package archive

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type format struct {
	name   string
	offset int64
	magic  []byte
	open   func(string) (Archiver, error)
}

var formats []format

func RegisterFormat(name string, offset int64, magic []byte, open func(string) (Archiver, error)) {
	formats = append(formats, format{name, offset, magic, open})
}

type Archiver interface {
	Files() []string
}

func Open(name string) (a Archiver, kind string, err error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	for _, f := range formats {
		if _, err := file.Seek(f.offset, 0); err != nil {
			return nil, "", err
		}
		buf := make([]byte, len(f.magic))
		if _, err := io.ReadFull(file, buf); err != nil {
			return nil, "", err
		}
		if err == nil && match(f.magic, buf) {
			a, err := f.open(file.Name())
			if err != nil {
				return nil, "", err
			}
			return a, f.name, nil
		}
	}
	return nil, "", fmt.Errorf("unsupported format")
}

func match(magic []byte, b []byte) bool {
	return reflect.DeepEqual(magic, b)
}
