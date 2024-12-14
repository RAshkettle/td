package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/RAshkettle/td/task"
)

// reopenCmd represents the reopen command
var reopenCmd = &cobra.Command{
	Use:   "reopen",
	Short: "Reopen a completed tasks.",
	Long: `Reopen a task that has been marked as completed.
  The task will be moved from the completed list to the active list.
USAGE: task reopen <id>`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please provide a valid task id")
			return
		}
		tp := task.NewFileSystemPersistor()

		task.ReOpen(id, tp)
	},
}

func init() {
	rootCmd.AddCommand(reopenCmd)
}
