package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all the files",
	Long:  `This command prints the working directory of this shell`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ls command is triggered: ", cmd.CalledAs())
		fmt.Println("ls command is triggered: ", args)
		//CommandExecutor()
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
