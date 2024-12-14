package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/RAshkettle/td/task"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task by its id.",
	Long: `Delete a task by its id.  This is to help clean up the task list or when you create a task incorrectly.
USAGE: task delete <id>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please enter a valid task id.")
			return
		}
		tp := task.NewFileSystemPersistor()

		task.Delete(id, tp)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
