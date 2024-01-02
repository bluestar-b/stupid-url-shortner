package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

var db = make(map[string]URLInfo)
var mutex sync.Mutex
var jsonFilePath = "urls.json"

func loadDataFromJSONFile() {
	file, err := os.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	err = json.Unmarshal(file, &db)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data:", err)
	}
}

func saveDataToJSONFile() {
	mutex.Lock()
	defer mutex.Unlock()

	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON data:", err)
		return
	}

	err = os.WriteFile(jsonFilePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing to JSON file:", err)
	}
}
