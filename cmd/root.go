package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

//var developer string

var RootCmd = &cobra.Command{
	Use:   "root [sub]",
	Short: "Gosh is a command line tool.",
	Long:  "Gosh is a command line tool that mimics the behaviour of a shell like zsh or bash in a minimalistic way\n\r" +
			"By minimalistic, we mean supporting only a limited number of commands",
	Args: cobra.MinimumNArgs(1),
	Aliases: []string{"man", "manual"},
}

func Execute() {
	RootCmd.AddCommand(pwdCmd, lsCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println("error in root command: ", err)
		os.Exit(1)
	}
}
