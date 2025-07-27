package scan

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type AudioFile struct {
	Path   string
	Title  string
	Artist string
	Album  string
	Cover  []byte
	Size   uint
	Type   string
}

type FSScanner struct {
	RootPath       string
	IndexFindings  bool
	ExcludePattern string
	IncludePattern string
	AddMeta        bool
	MetaFunc       func(AudioFile) (AudioFile, error)

	Result  []AudioFile
	Summary []string
}

func (fss *FSScanner) Scan() []error {
	audioFiles := []AudioFile{}
	errors := []error{}
	entries, err := fswalk(fss.RootPath)
	if err != nil {
		return []error{err}
	}

	if fss.AddMeta {
		for _, e := range entries {
			audioFile, err := fss.MetaFunc(AudioFile{Path: e})
			if err != nil {
				errors = append(errors, fmt.Errorf(fmt.Sprintf("failed getting meta from path '%s': ", err.Error())))
				continue
			}
			audioFiles = append(audioFiles, audioFile)
		}
	}

	fss.Result = audioFiles
	return errors
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

func NewDefaultFSScanner() FSScanner {
	fsScanner := FSScanner{
		RootPath:       "Music/",
		IndexFindings:  true,
		ExcludePattern: "",
		IncludePattern: "*.flac",
		AddMeta:        true,
		MetaFunc:       getMetaFromPath,
	}
	return fsScanner
}

func getMetaFromPath(af AudioFile) (AudioFile, error) {
	audioFile := af

	// expected path structure:
	// <root dir>/<artist>/<album>/<song>
	// expected song structure:
	// <optional index><optional index sepratator: " - "><song title>.<file type suffix>.
	// example:
	// Music/ZUNTATA/Groove Coaster (Original Soundtrack)/40 - Got more raves？.flac

	if audioFile.Path == "" {
		return AudioFile{}, fmt.Errorf("audioFile.Path can't be empty!")
	}

	matched, err := regexp.MatchString("[^/]+/[^/]+/[^/]+/[^/]+", audioFile.Path)

	if err != nil {
		return AudioFile{}, err
	}

	if !matched {
		return AudioFile{}, fmt.Errorf("invalid path structure, expected: <root dir>/<artist>/<album>/<song>")
	}

	re := regexp.MustCompile(`/?[^/]+/([^/]+)/([^/]+)/([0-9]+ ?- ?)?(.+)\.(?i).+$`)
	findings := re.FindStringSubmatch(audioFile.Path)

	audioFile.Artist = findings[1]
	audioFile.Album = findings[2]
	audioFile.Title = findings[len(findings)-1]

	return audioFile, nil
}

func NewFSScanner(rootPath string, indexFindings bool, excludePattern, includePattern string, addMeta bool, metaFunc func(AudioFile) (AudioFile, error)) FSScanner {
	fsScanner := FSScanner{
		RootPath:       rootPath,
		IndexFindings:  indexFindings,
		ExcludePattern: excludePattern,
		IncludePattern: includePattern,
		AddMeta:        addMeta,
		MetaFunc:       metaFunc,
	}

	return fsScanner
}
