package main2

import (
	"fmt"
)

// Need to give a tag if a task is not by that date, when the tasks are fetched
//But this can be a improvement feature
//If a new task is added get the number of the task do a +1 and add that a item number

//Have to handle errors , why?
//Application can't break if there's an error , So error should be handled well
//Error messages should be proper

type Task struct {
	Item   int
	Name   string
	Status string
	Date   string
}

var tasks []Task

func main() {

}
func cli() {

}

func addTask(name string, status string, date string) {
	task := Task{}
	item := tasks[len(tasks)-1]
	currentItemNumber := item.Item + 1
	task.Item = currentItemNumber
	//task.Item = id
	task.Name = name
	task.Status = status
	task.Date = date
	tasks = append(tasks, task)
	fmt.Printf("Task added: \n Number:%d\n  Name of the Task: %s \n Status of the Task %s \n Task to be completed by: %s", currentItemNumber, name, status, date)
}

func deleteTask(tasknumber int) {

	//If  a task is deleted
	//All the active tasks Number should be decremented by 1
	//If a task has to be delete by its Itemnumber can't you just delete by the item number? -> Because that's
	//index of array anyways
	tasks = append(tasks[:tasknumber], tasks[tasknumber+1:]...)
	task := Task{}
	for i := tasknumber - 1; i < len(tasks); i++ {
		tasks(i).Item = task.Item - 1
	}
}
func getAllTasks() {

	if len(tasks) == 0 {
		fmt.Println("No Tasks to List")
	} else {
		fmt.Println("Item Name Status")
		for _, task := range tasks {
			fmt.Println(task.Item, task.Name, task.Status)
		}
	}
	//Format the output using fmt.Printf("%-15s") -> Example

}
func getTaskByStatus() {

}
func getTasksbyString() {

}
func updateTask() {

}
