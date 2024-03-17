package parser

import (
	"encoding/json"
    "io/ioutil"
    "os"
	"log"
	"fmt"
)

type Structure struct {
	Data string `json:"Data"`
}

func ReadJson(filename string, append bool, ) {
	fileExists(filename)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	//var json_data map[string]interface{}
	data := []Structure{}
	json.Unmarshal(file, &data)
	fmt.Println(data)
}

func fileExists(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err){
		log.Fatal(err)
	} 
}

func WriteJson(data []Structure){
	dataBytes, err := json.Marshal(data)

}

// Dont do this add writeJson in structure.go itself