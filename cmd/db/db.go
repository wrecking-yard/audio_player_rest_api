package db

import (
	"fmt"
	"os"

	"codeberg.org/filipmnowak/audio_player_rest_api/internal/db/sqlite"
	"github.com/spf13/cobra"
)

func NewDBCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db",
		Short: "DB-related commands",
	}
	cmd.AddCommand(NewDBInitCmd())
	return cmd
}

func NewDBInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initalize database",
		Run: func(cmd *cobra.Command, _ []string) {
			DBInit(cmd)
		},
	}

	cmd.Flags().String("db-path", "data/index.sqlite3", "optional; filesystem path to SQLite DB")
	return cmd
}

func DBInit(cmd *cobra.Command) {
	dbPath, _ := cmd.Flags().GetString("db-path")
	db := sqlite.NewDB(nil, dbPath, "")
	if !db.Init() {
		for _, e := range db.InitErrors {
			fmt.Println(e)
		}
		os.Exit(1)
	}
}
