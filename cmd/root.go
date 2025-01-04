package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gotcher",
		Short: "A simple file watcher and responder",
		Long:  `gotcher is a tool to add actions to events in the filesystem.`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ExitWithError(err)
	}
}
