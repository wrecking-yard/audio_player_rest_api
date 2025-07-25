package scan

import (
	"github.com/spf13/cobra"
)

func NewScanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "scan file system for music files",
		Run: func(cmd *cobra.Command, args []string) {
			//targetDir, _ := cmd.Flags().GetString("target-dir")
			//indexFindings, _ := cmd.Flags().GetBool("index-findings?")
			// something something
		},
	}

	cmd.PersistentFlags().String("target-dir", "", "directory to scan for music files")
	cmd.PersistentFlags().Bool("index-findings?", true, "switching toggling indexing")
	return cmd
}
