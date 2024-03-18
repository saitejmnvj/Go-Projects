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






// Dont do this add writeJson in structure.go itself