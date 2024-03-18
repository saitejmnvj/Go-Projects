package logic

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"taskctl/taskctl/variables"
	"os"

	
)

type Task struct {
	Item   int `json:"Item"`
	Name   string `json:"Name"`
	Status string `json:"Status"`
	Date   string `json:"Date"`
}

// type Structure struct {
// 	Data string 
// }
// Struct will hold data of json file in unmarshalled format

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
	tasks := ReadJson(variables.FilePath)
	fmt.Println(tasks)
	
	//fmt.Println("This is task", &tasks)
	//Format the output using fmt.Printf("%-15s") -> Example

}

func AddTask(name string, status string, date string) {
	task := Task{}
	//Declaring task as empty instance of Task Struct, Other uage can be task:=Task{1,"hello","todo","10/03/2024"}
	var currentItemNumber int
	data := ReadJson(variables.FilePath)
	if len(data) == 0 {
		currentItemNumber = 1
	} else {
		item := data[len(data)-1]
		currentItemNumber = item.Item + 1
	}
	task.Item = currentItemNumber
	task.Name = name
	task.Status = status
	task.Date = date
	fmt.Println("This is len(data)", len(data))

	//fmt.Println("This is read data ", data)
	data = append(data, task)
	writeJson(data, variables.FilePath)
	fmt.Printf("Task added: \n Number:%d \n Name of the Task: %s \n Status of the Task %s \n Task to be completed by: %s", currentItemNumber, name, status, date)
	//fileExists(variables.FilePath)
}
//Append new task to exisitng tasks in json file

func writeJson(data []Task, filename string ){
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadJson(filename string) []Task{
	fileExists(filename)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	//var json_data map[string]interface{}
	data := []Task{}
	json.Unmarshal(file, &data)
	return data
}
//Read from file and pass the data back to getalltasks()



func fileExists(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err){
		log.Fatal(err)
	} 
}


