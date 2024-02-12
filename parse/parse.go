package parse

import (
	"fmt"
	"strings"
)

func CommandParse(commandString string) {
	// Parsed command from promptUI
	splittedResult := strings.Split(commandString, " ")
	fmt.Println("splittedResult: ", splittedResult)

	//	handle all command here using switch statement
	switch splittedResult[0] {
	case pwd:
		fmt.Println("pwd command is called")
	case ls:
		fmt.Println("ls command is called")
	case cat:
		fmt.Println("cat command is called")
	case cp:
	case mv:
	case find:
	case grep:
	case cd:
	case mkdir:
	case rmdir:
	default:
		fmt.Println("command not recognized")
		return
	}
}