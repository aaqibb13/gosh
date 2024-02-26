package cmd

import (
	"fmt"
	"os/exec"
)

func CommandExecutor(command, switchStatement string) string {
	//v, _ := os.Getwd()
	//print("working dir: ", v)
	cmd := exec.Command(command)
	if len(switchStatement) > 0 {
		cmd = exec.Command(command, switchStatement)
	}

	byteOutput, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", err)
		return ""
	}

	output := string(byteOutput[:])
	fmt.Println("\noutput: \n", output)
	return output
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("error: ", err)
	//}
}
