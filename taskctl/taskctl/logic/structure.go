package logic

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"taskctl/taskctl/variables"
	"taskctl/taskctl/parser"

	
)

type Task struct {
	Item   int
	Name   string
	Status string
	Date   string
}

var tasks []Task

//tasks is an array variable which will hold Struct Task as individual element of array , Basically array of Struct

func GetAllTasks() {

	// if len(tasks) == 0 {
	// 	fmt.Println("No Tasks to List")
	// } else {
	// 	fmt.Println("Item Name Status")
	// 	for _, task := range tasks {
	// 		fmt.Println(task.Item, task.Name, task.Status)
	// 	}
	// }
	parser.ReadJson(variables.FilePath)
	
	//fmt.Println("This is task", &tasks)
	//Format the output using fmt.Printf("%-15s") -> Example

}

func AddTask(name string, status string, date string) {
	task := Task{}
	//Declaring task as empty instance of Task Struct, Other uage can be task:=Task{1,"hello","todo","10/03/2024"}
	var currentItemNumber int

	if len(tasks) == 0 {
		currentItemNumber = 1
	} else {
		item := tasks[len(tasks)-1]
		currentItemNumber = item.Item + 1
	}

	task.Item = currentItemNumber
	//task.Item = id
	task.Name = name
	task.Status = status
	task.Date = date
	tasks = append(tasks, task)
	fmt.Printf("Task added: \n Number:%d \n Name of the Task: %s \n Status of the Task %s \n Task to be completed by: %s", currentItemNumber, name, status, date)
	writeJson(task)
}

func writeJson(value Task){
	//fileExists(variables.FilePath)
	content, err := json.Marshal(value)
	//converting struct(value) into json format using Marshal and assinging it to content
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(variables.FilePath,content,0644)
	if err != nil{
		log.Fatal(err)
	}
}


