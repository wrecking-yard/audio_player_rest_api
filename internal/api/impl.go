package api

// https://github.com/oapi-codegen/oapi-codegen/blob/116a63daf90260d279c1083e99f6718f4a5731e8/examples/minimal-server/stdhttp/api/impl.go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"codeberg.org/filipmnowak/audio_player_rest_api/internal/db/sqlite"
	"codeberg.org/filipmnowak/audio_player_rest_api/internal/scan"
	google_uuid "github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct {
	DB sqlite.DB
}

func NewServer(db sqlite.DB) Server {
	return Server{DB: db}
}

func (s Server) AlbumGetbyID(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	album := Album
	var artURL string
	_uuid := uuid.String()
	result, _ := GetAlbumByID(_uuid, s.DB)
	if result.SameAs(scan.Album{}) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(Error{})
	}

	_coverPath, _ := GetAlbumCoverPath(_uuid, s.DB)
	if _coverPath != "" {
		// TODO: auto-generate URL dynamically.
		// related: https://github.com/oapi-codegen/oapi-codegen/issues/163
		artURL = fmt.Sprintf("http://io:8080/api/album/%s/cover", result.Id)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Object{Id: &uuid, NameOrTitle: &result.Title, Type: &album, ArtURL: &artURL})
}

func (s Server) AlbumGetbyArtistID(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID, _ AlbumGetbyArtistIDParams) {
	resp := Object{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) ArtistGetbyID(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	resp := Object{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) SongGetbyID(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	resp := Object{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) SongsGetbyAlbumID(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID, _ SongsGetbyAlbumIDParams) {
	resp := Object{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) CoverGetbyAlbumID(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	_uuid := uuid.String()

	coverPath, _ := GetAlbumCoverPath(_uuid, s.DB)
	f, _ := os.Open(coverPath)
	cover, _ := io.ReadAll(f)
	w.WriteHeader(http.StatusOK)
	w.Write(cover)
}

func (s Server) SongsGetbyArtistID(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID, _ SongsGetbyArtistIDParams) {
	resp := Object{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) Enqueue(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	resp := Object{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) Play(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	resp := Object{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) PlayPause(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s Server) PlayUnpause(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s Server) QueueDeleteFromByRange(w http.ResponseWriter, r *http.Request, from int64, to int64) {
	resp := Object{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) QueueReorder(w http.ResponseWriter, r *http.Request, from int64, to int64) {
	resp := Object{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) SearchForArtists(w http.ResponseWriter, r *http.Request, search string, _ SearchForArtistsParams) {
	artist := Artist
	results, _ := SearchForArtists(search, s.DB)
	resp := []Object{}
	for _, r := range results {
		_uuid, _ := google_uuid.Parse(r.Id)
		resp = append(resp, Object{Id: &_uuid, NameOrTitle: &r.Name, Type: &artist})
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) SearchForAlbums(w http.ResponseWriter, r *http.Request, search string, _ SearchForAlbumsParams) {
	album := Album
	results, _ := SearchForAlbums(search, s.DB)
	resp := []Object{}
	for _, r := range results {
		_uuid, _ := google_uuid.Parse(r.Id)
		resp = append(resp, Object{Id: &_uuid, NameOrTitle: &r.Title, Type: &album})
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s Server) SearchForSongs(w http.ResponseWriter, r *http.Request, search string, params SearchForSongsParams) {
	song := Song
	results, _ := SearchForSongs(search, s.DB)
	resp := []Object{}
	for _, r := range results {
		_uuid, _ := google_uuid.Parse(r.Id)
		resp = append(resp, Object{Id: &_uuid, NameOrTitle: &r.Title, Type: &song})
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
