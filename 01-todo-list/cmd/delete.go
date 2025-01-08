package cmd

import (
	"01-todo-list/db"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		err = db.DeleteTask(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	},
}
