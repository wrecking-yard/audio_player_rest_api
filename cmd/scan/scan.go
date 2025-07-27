package scan

import (
	"fmt"
	"os"

	"codeberg.org/filipmnowak/audio_player_rest_api/internal/scan"
	"github.com/spf13/cobra"
)

func NewScanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "scan file system for music files",
		Run: func(cmd *cobra.Command, _ []string) {
			Scan(cmd)
		},
	}

	cmd.PersistentFlags().String("target-dir", "", "directory to scan for music files")
	cmd.MarkPersistentFlagRequired("target-dir")
	cmd.PersistentFlags().Bool("index-findings", true, "switching toggling indexing")
	return cmd
}

func Scan(cmd *cobra.Command) {
	targetDir, _ := cmd.Flags().GetString("target-dir")
	//indexFindings, _ := cmd.Flags().GetBool("index-findings")

	if targetDir == "" {
		fmt.Printf("target-dir flag can't set to an empty string!\n")
		os.Exit(1)
	}

	fss := scan.NewDefaultFSScanner()
	err := fss.Scan()
	if len(err) > 0 {
		for _, e := range err {
			fmt.Printf(e.Error())
		}
		os.Exit(1)
	}

	for _, f := range fss.Result {
		fmt.Printf("Path: %s\n", f.Path)
		fmt.Printf("Title: %s\n", f.Title)
		fmt.Printf("Artist: %s\n", f.Artist)
		fmt.Printf("Album: %s\n", f.Album)
		fmt.Printf("\n\n")
	}
}
