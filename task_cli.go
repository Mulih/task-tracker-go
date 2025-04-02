package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

const taskFile = "ttasks.json"

func (tl *TaskList) SaveTasks() error {
	file, err := os.Create(taskFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tl)
}

func (tl *TaskList) LoadTasks() error {
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(tl)
}

func main() {
	// Load tasks from the file at the start
	taskList := &TaskList{}
	if err := taskList.LoadTasks(); err != nil {
		panic(err)
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run task_cli.go <command>")
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run task_cli.go add <task_description>")
			return
		}
		description := os.Args[2]
		taskList.AddTask(description)
		fmt.Printf("Task added: %s\n", description)
	}
}
