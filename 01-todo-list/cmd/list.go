package cmd

import (
	"01-todo-list/db"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

var showAll bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks (per default all unfinished)",
	Run: func(cmd *cobra.Command, args []string) {
		// get tasks
		tasks, err := db.GetTasks()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		var processedTasks []db.Task
		// filter based on flag
		if !showAll {
			for i := range tasks {
				if !tasks[i].Done {
					processedTasks = append(processedTasks, tasks[i])
				}
			}
		} else {
			processedTasks = tasks
		}

		// print to stdout
		writer := tabwriter.NewWriter(os.Stdout, 5, 0, 1, ' ', 0)
		fmt.Fprintln(writer, "Id\tTitle\tDone")
		for i := range processedTasks {
			fmt.Fprintf(writer, "%v\t%v\t%v\n", processedTasks[i].Id, processedTasks[i].Title, processedTasks[i].Done)
		}
		writer.Flush()
	},
}
