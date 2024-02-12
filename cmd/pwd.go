package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "Print the working directory",
	Long:  `This command prints the working directory of this shell`,
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pwd command is triggered: ", cmd.CalledAs())
	},
}
