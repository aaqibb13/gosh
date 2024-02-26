package ui

import (
	"bufio"
	"fmt"
	"gosh/main/parse"
	"os"
	"strings"
)

func ExecuteCommand() string {
	for {
		// Keeping the command running
		reader := bufio.NewReader(os.Stdin)
		const colorGreen = "\033[0;32m"
		const colorNone = "\033[0m"

		_, _ = fmt.Fprintf(os.Stdin, "%s gosh> %s", colorGreen, colorNone)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return err.Error()
		}

		// Remove trailing newline character
		input = strings.TrimSpace(input)

		// Process the command
		//fmt.Println("You entered:", input)

		parse.CommandParse(input)
	}
}
