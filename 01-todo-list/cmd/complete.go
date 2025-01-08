package cmd

import (
	"01-todo-list/db"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	rootCmd.AddCommand(completeCmd)
}

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		err = db.CompleteTask(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	},
}
