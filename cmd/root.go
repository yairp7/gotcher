package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yairp7/gotcher/internal/utils"
)

var (
	logger   utils.Logger
	logLevel string

	rootCmd = &cobra.Command{
		Use:   "gotcher",
		Short: "A simple file watcher and responder",
		Long:  `gotcher is a tool to add actions to events in the filesystem.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Initialize logger
			stdLogger := utils.NewStdoutLogger()

			// Set log level
			switch logLevel {
			case "debug":
				stdLogger.SetLevel(utils.DEBUG)
			case "info":
				stdLogger.SetLevel(utils.INFO)
			case "warn":
				stdLogger.SetLevel(utils.WARN)
			case "error":
				stdLogger.SetLevel(utils.ERROR)
			default:
				stdLogger.SetLevel(utils.INFO)
			}

			logger = stdLogger
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Set the logging level (debug, info, warn, error)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error("Error: %v", err)
		ExitWithError(err)
	}
}
