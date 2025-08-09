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

	cmd.Flags().String("db-path", "data/index.sqlite3", "optional; filesystem path to SQLite DB")

	return cmd
}

func Serve(cmd *cobra.Command) {
	dbPath, _ := cmd.Flags().GetString("db-path")
	db_ro = sqlite.NewDB(nil, dbPath, "")
	server.Serve()
}
