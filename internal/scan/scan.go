package scan

import (
	"os"
	"path/filepath"
)

type AudioFile struct {
	Path string
	Size uint
	Type string
}

type FSScanner struct {
	RootPath       string
	IndexFindings  bool
	ExcludePattern string
	IncludePattern string
	AddMeta        bool
	Result         []string
}

func (fss *FSScanner) FSWalk() error {
	entries, err := fswalk(fss.RootPath)
	if err != nil {
		return err
	}
	fss.Result = entries
	return nil
}

func fswalk(root string) ([]string, error) {
	entries := []string{}

	dirEntry, _ := os.ReadDir(root)

	for _, e := range dirEntry {
		if e.Type().IsRegular() {
			entries = append(entries, filepath.Join(root, e.Name()))
		}

		if e.Type().IsDir() {
			_entries, _ := fswalk(filepath.Join(root, e.Name()))
			entries = append(entries, _entries...)
		}
	}

	return entries, nil
}
