package cmd

import (
	"github.com/spf13/cobra"

	"github.com/RAshkettle/td/task"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task.",
	Args:  cobra.ExactArgs(1),
	Long: `Add a new task to the list.  Note:  The task is not unique and can be added multiple times.
  This is intentional to allow for multiple instances of the same task.
  USAGE:  td add task <description>`,
	Run: func(cmd *cobra.Command, args []string) {
		tp := task.NewFileSystemPersistor()
		description := args[0]
		task.Add(description, tp)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
