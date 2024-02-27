package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func CommandExecutor(command, switchStatement string) {
	cmd := exec.Command(command)
	if len(switchStatement) > 0 {
		cmd = exec.Command(command, switchStatement)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
}
