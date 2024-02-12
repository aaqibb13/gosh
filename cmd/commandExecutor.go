package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func CommandExecutor(command, switchStatement string) {
	v, _ := os.Getwd()
	print("working dir: ", v)
	cmd := exec.Command(command)
	if len(switchStatement) > 0 {
		cmd = exec.Command(command, switchStatement)
	}

	byteOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	//print("out here: ", byteOutput)
	output := string(byteOutput[:])
	fmt.Println("\noutput: \n", output)
}
