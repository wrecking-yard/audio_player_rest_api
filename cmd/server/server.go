package server

import (
	"net"

	"codeberg.org/filipmnowak/audio_player_rest_api/internal/db/sqlite"
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
	cmd.Flags().IP("ip", net.IPv4(0, 0, 0, 0), "optional; IP to bind to")
	cmd.Flags().Uint32("port", 8080, "optional; port to bind to")
	cmd.Flags().String("db-path", "data/index.sqlite3", "optional; filesystem path to SQLite DB")
	return cmd
}

func Serve(cmd *cobra.Command) {
	ip, _ := cmd.Flags().GetIP("ip")
	port, _ := cmd.Flags().GetUint32("port")
	dbPath, _ := cmd.Flags().GetString("db-path")
	db := sqlite.NewDB(nil, dbPath, "")
	server.Serve(db, ip, port)
}
