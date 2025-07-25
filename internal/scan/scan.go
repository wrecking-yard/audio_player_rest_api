package scan

import (
	"fmt"
	"os"
	"path/filepath"
)

type File struct {
	Path string
	Size uint
}

type FSScanner struct {
	RootPath       string
	ExcludePattern string
	IncludePattern string
	AddMeta        bool
	Result         []File
}

func (fss FSScanner) FSWalk() error {

	return fmt.Errorf("error: something something")
}

func fswalk(rootPath string) ([]string, error) {
	entries := []string{}
	err := filepath.Walk(rootPath, func(path string, f os.FileInfo, err error) error {
		if f.Mode().IsRegular() {
			entries = append(entries, path)
		}
		return nil
	})

	if err != nil {
		return []string{}, err
	}

	return entries, nil
}
