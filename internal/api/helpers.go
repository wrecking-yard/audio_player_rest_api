package api

import (
	"encoding/json"
	"fmt"

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
