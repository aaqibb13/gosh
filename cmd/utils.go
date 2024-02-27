package cmd

import (
	"regexp"
	"strings"
)
// TruncateExtraSpacesInCommand removes extra spaces from the commands so they work as intended
// function takes command (string) as an input parameter
// returns []string which is the main command and the subcommand

func TruncateExtraSpacesInCommand(command string) []string {
	singleSpacePattern := regexp.MustCompile(`\s+`)
	formattedCommand := singleSpacePattern.ReplaceAllString(command, " ")
	splittedResult := strings.Split(formattedCommand, " ")
	return splittedResult
}
