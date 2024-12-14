package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/RAshkettle/td/task"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all Tasks",
	Long:  `Lists all Tasks in the task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		tp := task.NewFileSystemPersistor()

		task.List(tp)
		// Display this in a table using bubbletea
		fmt.Println("ID\tDescription\tStatus")
		for _, t := range task.Tasks {
			fmt.Printf(
				"%d\t%s\t%s\n",
				t.FriendlyId,
				t.Description,
				t.Status.String(),
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
