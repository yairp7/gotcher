package utils

import (
	"context"
	"os"
	"os/exec"
)

func ExecShell(ctx context.Context, cmd string) error {
	execCmd := exec.Command("bash", "-c", cmd)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	return execCmd.Run()
}
