package scan

import (
	"fmt"
	"os"
	"slices"

	"codeberg.org/filipmnowak/audio_player_rest_api/internal/db/sqlite"
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
	cmd.Flags().String("exclude-pattern", "", "exclude directories or files by pattern")
	cmd.Flags().String("include-pattern", "", "override exclude pattern")
	cmd.Flags().Bool("index-findings", true, "indexing on/off switch")
	cmd.Flags().String("meta-discovery-mode", "path", "supported mechanisms: path, flac, dynamic")
	cmd.Flags().String("db-path", "data/index.sqlite3", "optional; filesystem path to SQLite DB")
	return cmd
}

func Scan(cmd *cobra.Command) {
	targetDir, _ := cmd.Flags().GetString("target-dir")
	indexFindings, _ := cmd.Flags().GetBool("index-findings")
	excludePattern, _ := cmd.Flags().GetString("exclude-pattern")
	includePattern, _ := cmd.Flags().GetString("include-pattern")
	metaDiscoMode, _ := cmd.Flags().GetString("meta-discovery-mode")
	dbPath, _ := cmd.Flags().GetString("db-path")

	if targetDir == "" {
		fmt.Printf("target-dir flag can't set to an empty string!\n")
		os.Exit(1)
	}

	if !slices.Contains([]string{"path", "flac", "dynamic"}, metaDiscoMode) {
		fmt.Printf("meta-discovery-mode needs to be one of: path, flac, dynamic!\n")
		os.Exit(1)
	}

	db := sqlite.NewDB(nil, dbPath, "")
	db.Init()
	if len(db.InitErrors) > 0 {
		for _, err := range db.InitErrors {
			fmt.Printf(err.Error.Error())
		}
		os.Exit(1)
	}

	fss := scan.NewFSScanner(targetDir, indexFindings, excludePattern, includePattern, scan.MetaDiscoModes[metaDiscoMode], db)
	err := fss.Scan()
	if len(err) > 0 {
		for _, e := range err {
			fmt.Printf(e.Error())
		}
		os.Exit(1)
	}
}
