package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/RAshkettle/td/task"
)

// closeCmd represents the close command
var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Mark a task as completed.",
	Long: `Mark a task as completed.  This will close the task.
  USAGE: task close <id>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please provide a valid task id")
			return
		}
		tp := task.NewFileSystemPersistor()

		task.Close(id, tp)
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
}
