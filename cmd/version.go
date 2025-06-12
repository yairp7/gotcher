package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yairp7/gotcher/internal/version"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information of gotcher",
	Long:  `Display detailed version information including the version number, git commit hash, and build date.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.GetVersionInfo())
	},
}
