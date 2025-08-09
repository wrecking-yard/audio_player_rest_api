package server

import (
	"codeberg.org/filipmnowak/audio_player_rest_api/internal/api"
	"embed"
	"net/http"
)

//go:embed swagger-ui
var swaggerUI embed.FS

func Serve() {

	fs := http.FileServer(http.FS(swaggerUI))

	mux := http.NewServeMux()
	mux.Handle("/swagger-ui/", fs)

	// https://github.com/oapi-codegen/oapi-codegen/blob/116a63daf90260d279c1083e99f6718f4a5731e8/examples/minimal-server/stdhttp/main.go
	apiServer := api.NewServer()
	h := api.HandlerFromMuxWithBaseURL(apiServer, mux, "/api")

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	s.ListenAndServe()
}
