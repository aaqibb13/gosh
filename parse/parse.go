package parse

import (
	"fmt"
	"gosh/main/cmd"
	"os"
	"strings"
)

func CommandParse(commandString string) {
	// Parsed command from promptUI
	splittedResult := strings.Split(commandString, " ")
	//fmt.Println("splittedResult: ", splittedResult)

	//	handle all command here using switch statement
	switch splittedResult[0] {
	case pwd:
		fmt.Println("pwd command is called")
	case ls:
		fmt.Println("ls command is called")
		cmd.Execute()
		//if len(splittedResult) > 1 {
		//	cmdResponse := cmd.CommandExecutor(splittedResult[0], splittedResult[1])
		//	return cmdResponse
		//} else {
		//	cmdResponse := cmd.CommandExecutor(splittedResult[0], "")
		//	return cmdResponse
		//}
	case cat:
		fmt.Println("cat command is called")
	case cp:
	case mv:
	case find:
	case grep:
	case cd:
	case mkdir:
	case rmdir:
	case exit:
		os.Exit(0)
	default:
		err := fmt.Sprintf("gosh: command not recognized: %s", splittedResult[0])
		fmt.Println(err)
	}
	return
}