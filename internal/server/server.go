package server

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed swagger-ui
var swaggerUI embed.FS

func Serve() {

	fs := http.FileServer(http.FS(swaggerUI))

	mux := http.NewServeMux()
	mux.Handle("/swagger-ui/", fs)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
