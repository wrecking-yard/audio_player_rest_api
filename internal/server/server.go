package server

import (
	"embed"
	"fmt"
	"net"
	"net/http"

	"codeberg.org/filipmnowak/audio_player_rest_api/internal/api"
	"codeberg.org/filipmnowak/audio_player_rest_api/internal/db/sqlite"
)

//go:embed swagger-ui
var swaggerUI embed.FS

func Serve(db sqlite.DB, ip net.IP, port uint32) {

	fs := http.FileServer(http.FS(swaggerUI))

	mux := http.NewServeMux()
	mux.Handle("/swagger-ui/", fs)

	// https://github.com/oapi-codegen/oapi-codegen/blob/116a63daf90260d279c1083e99f6718f4a5731e8/examples/minimal-server/stdhttp/main.go
	apiServer := api.NewServer(db)
	h := api.HandlerFromMuxWithBaseURL(apiServer, mux, "/api")

	s := &http.Server{
		Handler: h,
		Addr:    ip.String() + ":" + fmt.Sprintf("%v", port),
	}

	s.ListenAndServe()
}
