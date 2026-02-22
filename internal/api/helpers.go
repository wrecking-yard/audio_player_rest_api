package api

import (
	"encoding/json"
	"fmt"
	"strings"

	"codeberg.org/filipmnowak/audio_player_rest_api/internal/db/sqlite"
	"codeberg.org/filipmnowak/audio_player_rest_api/internal/scan"
)

func SearchForArtists(s string, db sqlite.DB) ([]scan.Artist, error) {
	var as []scan.Artist

	query := `SELECT uuid, name FROM artists WHERE name LIKE '%%%s%%';`
	query = fmt.Sprintf(query, s)
	result, _ := db.RunStatement(query, false, false, false)

	err := json.Unmarshal([]byte(result), &as)
	if err != nil {
		return []scan.Artist{}, err
	}

	return as, nil
}

func SearchForAlbums(s string, db sqlite.DB) ([]scan.Album, error) {
	var as []scan.Album

	query := `SELECT uuid, title FROM albums WHERE title LIKE '%%%s%%';`
	query = fmt.Sprintf(query, s)
	result, _ := db.RunStatement(query, false, false, false)

	err := json.Unmarshal([]byte(result), &as)
	if err != nil {
		return []scan.Album{}, err
	}

	return as, nil
}

func SearchForSongs(s string, db sqlite.DB) ([]scan.AudioFile, error) {
	var afs []scan.AudioFile

	query := `SELECT uuid, title FROM songs WHERE songs.title LIKE '%%%s%%';`
	query = fmt.Sprintf(query, s)
	result, _ := db.RunStatement(query, false, false, false)

	err := json.Unmarshal([]byte(result), &afs)
	if err != nil {
		return []scan.AudioFile{}, err
	}

	return afs, nil
}

func GetAlbumByID(albumUUID string, db sqlite.DB) (scan.Album, error) {
	var album scan.Album

	query := `SELECT title, uuid FROM albums WHERE uuid = '%s';`
	query = fmt.Sprintf(query, albumUUID)
	result, _ := db.RunStatement(query, false, false, false)

	err := json.Unmarshal([]byte(result), &album)
	if err != nil {
		return scan.Album{}, err
	}

	return album, nil
}

func GetAlbumCoverPath(albumUUID string, db sqlite.DB) (string, error) {
	// so, cover is a special kind of song, apparently.
	query := `SELECT path FROM songs WHERE songs.album_uuid = '%s' AND title = 'cover' LIMIT 1;`
	query = fmt.Sprintf(query, albumUUID)
	result, err := db.RunStatement(query, false, false, true)

	if err != nil {
		return "", err
	}

	return strings.TrimRight(result, "\r\n"), nil
}
