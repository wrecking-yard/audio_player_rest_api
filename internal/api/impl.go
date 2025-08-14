package api

// https://github.com/oapi-codegen/oapi-codegen/blob/116a63daf90260d279c1083e99f6718f4a5731e8/examples/minimal-server/stdhttp/api/impl.go

import (
	"encoding/json"
	"net/http"

	"codeberg.org/filipmnowak/audio_player_rest_api/internal/db/sqlite"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct{
	DB sqlite.DB
}

func NewServer(db sqlite.DB) Server {
	return Server{DB: db}
}

func (s Server) AlbumGetbyID(w http.ResponseWriter, r *http.Request, uuid openapi_types.UUID) {
	resp := Object{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
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
		_uuid, _ := uuid.Parse(r.Id)
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
		_uuid, _ := uuid.Parse(r.Id)
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
		_uuid, _ := uuid.Parse(r.Id)
		resp = append(resp, Object{Id: &_uuid, NameOrTitle: &r.Title, Type: &song}) 
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
