package cmd

import (
	"taskctl/taskctl/logic"

	"github.com/spf13/cobra"
)

var task string
var status string
var date string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "taskctl add will a ta into the catalogue",
	Run: func(cmd *cobra.Command, args []string) {
		task, err := cmd.Flags().GetString("task")
		if err != nil {
			panic(err)
		}
		status, err := cmd.Flags().GetString("status")
		if err != nil {
			panic(err)
		}
		date, err := cmd.Flags().GetString("date")
		if err != nil {
			panic(err)
		}
		logic.AddTask(task, status, date)

		//sum := 5
		//logic.GetAllTasks()

		//sum:=logic.getAllTasks()
		//fmt.Println("List is this big", sum)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVar(&task, "task", "", "Description of the task to be added")
	addCmd.Flags().StringVar(&status, "status", "", "Add Status of the task like New, Started , Completed")
	addCmd.Flags().StringVar(&date, "date", "", "Add Date by which this task has to be completed")
}
