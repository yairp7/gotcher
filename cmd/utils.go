package cmd

import (
	"os"
)

func ExitWithError(err error) {
	logger.Error("%v", err)
	os.Exit(1)
}
