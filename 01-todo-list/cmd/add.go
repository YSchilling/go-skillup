package cmd

import (
	"01-todo-list/db"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := db.AddTask(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		fmt.Println("Added task: " + args[0])
	},
}
