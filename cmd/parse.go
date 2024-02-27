package cmd

import (
	"fmt"
	"os"
)

func CommandParse(command string) {
	// call the utility here
	args := TruncateExtraSpacesInCommand(command)

	//	handle all command here using switch statement
	switch args[0] {
	case pwd:
		CommandExecutor(args[0], "")
	case ls:
		if len(args) > 1 {
			CommandExecutor(args[0], args[1])
		} else {
			CommandExecutor(args[0], "")
		}
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
		err := fmt.Sprintf("gosh: command not recognized: %s", args[0])
		fmt.Println(err)
	}
	return
}