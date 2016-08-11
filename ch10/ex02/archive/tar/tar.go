package tar

import (
	"archive/tar"
	"io"
	"os"

	"../../archive"
)

func init() {
	magic := []byte{'u', 's', 't', 'a', 'r', 0, '0', '0'}
	archive.RegisterFormat("tar", 0x101, magic, open)
}

func open(name string) (archive.Archiver, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var files []string
	r := tar.NewReader(file)
	for {
		header, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		files = append(files, header.Name)
	}
	return &tarArchiver{files}, nil
}

type tarArchiver struct {
	files []string
}

func (a *tarArchiver) Files() []string {
	return a.files
}
