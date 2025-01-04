package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var version = "0.1"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gotcher",
	Long:  `This is gotcher's current version`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("gotcher v%s\n", version)
	},
}
