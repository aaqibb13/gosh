package cmd

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func Run() {

	for {
		runtimeEnv := runtime.GOOS
		if runtimeEnv != "darwin" {
			fmt.Printf("command line tool works on %s only", runtimeEnv)
			break
		}

		// Keeping the command running
		reader := bufio.NewReader(os.Stdin)
		const colorGreen = "\033[0;32m"
		const colorNone = "\033[0m"

		_, _ = fmt.Fprintf(os.Stdin, "%s gosh> %s", colorGreen, colorNone)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Remove trailing newline character
		input = strings.TrimSpace(input)
		CommandParse(input)
	}
}
