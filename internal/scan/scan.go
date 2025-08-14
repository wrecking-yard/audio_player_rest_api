package scan

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"codeberg.org/filipmnowak/audio_player_rest_api/internal/db/sqlite"
	"codeberg.org/filipmnowak/audio_player_rest_api/internal/helpers"
)

var (
	MetaDiscoModes = map[string]func(AudioFile) (AudioFile, error){
		"path": getMetaFromPath,
		//"flac": getMetaFromFlac,
		//"dynamic": getMetaFromDynamicLookup,
	}
)

type Artist struct {
	Id   string `json:"uuid"`
	Name string
}

type Album struct {
	Id     string `json:"uuid"`
	Path   string
	Title  string
	Artist Artist
	Cover  []byte
}

func (albumA Album) SameAs(albumB Album) bool {
	return reflect.DeepEqual(albumA, albumB)
}

func (album Album) OneOf(albums []Album) bool {
	for _, a := range albums {
		if reflect.DeepEqual(album, a) {
			return true
		}
	}
	return false
}

type AudioFile struct {
	Id     string `json:"uuid"`
	Path   string
	Title  string
	Artist Artist
	Album  Album
	Cover  []byte
	Size   uint
	Type   string
}

type FSScanner struct {
	RootPath       string
	IndexFindings  bool
	ExcludePattern string
	IncludePattern string
	MetaFunc       func(AudioFile) (AudioFile, error)
	Result         []AudioFile
	Summary        []string
	DB             sqlite.DB
}

func (fss *FSScanner) Scan() []error {
	audioFiles := []AudioFile{}
	errors := []error{}
	entries, err := fswalk(fss.RootPath)
	if err != nil {
		return []error{err}
	}

	for _, e := range entries {
		audioFile, err := fss.MetaFunc(AudioFile{Path: e})
		if err != nil {
			errors = append(errors, fmt.Errorf(fmt.Sprintf("failed getting meta from path '%s': ", err.Error())))
			continue
		}
		audioFiles = append(audioFiles, audioFile)
	}

	if fss.IndexFindings {
		err = indexFindings(audioFiles, fss.DB)
		if err != nil {
			errors = append(errors, err)
		}
	}

	fss.Result = audioFiles
	return errors
}

func indexFindings(afs []AudioFile, db sqlite.DB) error {
	// compile list of all artists
	artists := []string{}
	for _, af := range afs {
		if !slices.Contains(artists, af.Artist.Name) {
			artists = append(artists, af.Artist.Name)
		}
	}

	// index artists
	input := []map[string]string{}
	for _, a := range artists {
		input = append(input, map[string]string{"name": a, "uuid": sqlite.UUID4()})
	}
	_, _ = db.TransactUpserts(input, "artists")

	// compile list of all albums
	albums := []Album{}
	for _, af := range afs {
		if !af.Album.OneOf(albums) {
			albums = append(albums, af.Album)
		}
	}

	// index albums
	input = []map[string]string{}
	for _, a := range albums {
		// get artist id
		artist_uuid, err := db.RunStatement(fmt.Sprintf("SELECT uuid FROM artists WHERE name = '%s';", a.Artist.Name), false, false, true)
		if err != nil {
			return err
		}
		input = append(input, map[string]string{"title": a.Title, "artist_uuid": artist_uuid, "uuid": sqlite.UUID4(), "path": a.Path})
	}
	_, _ = db.TransactUpserts(input, "albums")

	// index songs
	input = []map[string]string{}
	for _, af := range afs {
		// get artist id
		artist_uuid, err := db.RunStatement(fmt.Sprintf("SELECT uuid FROM artists WHERE name = '%s';", af.Artist.Name), false, false, true)
		// get album id
		album_uuid, err := db.RunStatement(fmt.Sprintf("SELECT uuid FROM albums WHERE title = '%s';", af.Album.Title), false, false, true)
		if err != nil {
			return err
		}
		input = append(input, map[string]string{"title": af.Title, "artist_uuid": strings.Trim(artist_uuid, "\r\n"), "album_uuid": strings.Trim(album_uuid, "\r\n"), "uuid": sqlite.UUID4(), "path": af.Path})
	}
	if len(afs) > 30 {
		splitInput, _ := helpers.SplitDBInput(input, 8)
		for _, chunk := range splitInput {
			_, _ = db.TransactUpserts(chunk, "songs")
		}
	} else {
		_, _ = db.TransactUpserts(input, "songs")
	}

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

	audioFile.Artist.Name = findings[1]
	audioFile.Album.Title = findings[2]
	audioFile.Album.Path = findings[1] + "/" + findings[2]
	audioFile.Album.Artist.Name = findings[1]
	audioFile.Title = findings[len(findings)-1]

	return audioFile, nil
}

func NewFSScanner(rootPath string, indexFindings bool, excludePattern, includePattern string, metaFunc func(AudioFile) (AudioFile, error), db sqlite.DB) FSScanner {
	fsScanner := FSScanner{
		RootPath:       rootPath,
		IndexFindings:  indexFindings,
		ExcludePattern: excludePattern,
		IncludePattern: includePattern,
		MetaFunc:       metaFunc,
		DB:             db,
	}

	return fsScanner
}
