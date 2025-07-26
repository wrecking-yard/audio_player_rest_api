package cmd

import (
	"codeberg.org/filipmnowak/audio_player_rest_api/cmd/db"
	"codeberg.org/filipmnowak/audio_player_rest_api/cmd/scan"
	"codeberg.org/filipmnowak/audio_player_rest_api/cmd/server"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "apra",
	Short: "Audio Player REST API",
}

func init() {
	rootCmd.AddCommand(db.NewDBCmd())
	rootCmd.AddCommand(scan.NewScanCmd())
	rootCmd.AddCommand(server.NewServerCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
