package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data := make(map[string]string)
	var name, address string

	fmt.Print("Name: ")
	fmt.Scan(&name)

	fmt.Print("Address: ")
	fmt.Scan(&address)

	data["name"] = name
	data["address"] = address

	jsonData, err := json.Marshal(data)
	if err == nil {
		os.Stdout.Write(jsonData)
	}
}
