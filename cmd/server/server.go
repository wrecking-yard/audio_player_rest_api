package server

import (
	"codeberg.org/filipmnowak/audio_player_rest_api/internal/server"
	"github.com/spf13/cobra"
)

func NewServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "HTTP server-related functionalities",
		Run: func(cmd *cobra.Command, _ []string) {
			Serve(cmd)
		},
	}

	return cmd
}

func Serve(cmd *cobra.Command) {
	server.Serve()
}
