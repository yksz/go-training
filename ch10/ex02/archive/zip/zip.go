package zip

import (
	"archive/zip"

	"../../archive"
)

func init() {
	magic := []byte{'P', 'K'}
	archive.RegisterFormat("zip", 0, magic, open)
}

func open(name string) (archive.Archiver, error) {
	rc, err := zip.OpenReader(name)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	var files []string
	for _, file := range rc.File {
		files = append(files, file.Name)
	}
	return &zipArchiver{files}, nil
}

type zipArchiver struct {
	files []string
}

func (a *zipArchiver) Files() []string {
	return a.files
}
