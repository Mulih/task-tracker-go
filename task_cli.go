package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Task represents a task with an ID, description, status, and timestamps.
// The ID is unique for each task, and the status can be "todo", "in progress", or "done".
// The CreatedAt and UpdatedAt fields store the timestamps of when the task was created and last updated.

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// The TaskList struct contains a slice of tasks and the next available ID for new tasks.
type TaskList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

const taskFile = "tasks.json"

// SaveTasks saves the current task list to a JSON file.
// It creates the file if it does not exist and overwrites it if it does.
// The file is created in the current working directory.
func (tl *TaskList) SaveTasks() error {
	file, err := os.Create(taskFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tl)
}

// LoadTasks loads tasks from the JSON file into the TaskList struct.
// If the file does not exist, it returns nil without error.
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

	// Check the command line arguments for the command to execute
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
	case "list":
		tasks := taskList.GetTasks()
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC3339), task.UpdatedAt.Format(time.RFC3339))
		}
	}


}
