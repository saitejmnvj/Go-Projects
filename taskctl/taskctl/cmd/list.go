package cmd

import (
	"taskctl/taskctl/logic"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "taskctl list will list all the tasks that are processed",
	Run: func(cmd *cobra.Command, args []string) {
		//sum := 5
		logic.GetAllTasks()

		//sum:=logic.getAllTasks()
		//fmt.Println("List is this big", sum)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
