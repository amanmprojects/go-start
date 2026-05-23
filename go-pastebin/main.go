package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Paste struct {
	Content string
	ID      string
}

const filename = "pastes.json"

func randomID() string {
	id := uuid.New()
	return id.String()
}

func addPaste(paste Paste) {
	data, err := os.ReadFile(filename)
	if err != nil {
		data = []byte("[]")
	}
	var pastes []Paste
	json.Unmarshal(data, &pastes)

	pastes = append(pastes, paste)
	jsonData, err := json.Marshal(pastes)
	if err != nil {
		return
	}
	os.WriteFile(filename, jsonData, 0644)
}

func getPaste(id string) (Paste, bool) {
	data, err := os.ReadFile(filename)
	if err != nil {
		data = []byte("[]")
	}
	var pastes []Paste
	json.Unmarshal(data, &pastes)

	for _, paste := range pastes {
		if paste.ID == id {
			return paste, true
		}
	}
	return Paste{}, false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a paste content")
			return
		}
		content := os.Args[2]
		id := randomID()
		paste := Paste{Content: content, ID: id}
		addPaste(paste)
		fmt.Println("Paste added successfully")
		fmt.Println("ID:", id)
	case "get":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a paste ID")
			return
		}
		id := os.Args[2]
		paste, found := getPaste(id)
		if found {
			fmt.Println("Content:", paste.Content)
		} else {
			fmt.Println("Paste not found")
		}

	default:
		fmt.Println("Unknown command")
	}
}
