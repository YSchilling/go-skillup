package cmd

import (
	"01-todo-list/config"
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Tasks is a simple cli todo app",
}

func init() {
	headers := []string{"id", "title", "done"}

	// check existence of db file
	_, err := os.Stat(config.DBPath)
	if err == nil {
		return
	}

	// if other error print err
	if !os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// creat the file
	file, err := os.Create(config.DBPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(headers); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
