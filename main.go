package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"task-tracker/tasks" // Import the tasks package
)

func main() {
	// Load tasks from the file at the start
	taskList := &tasks.TaskList{} // Use tasks.TaskList
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
		status := os.Args[3]
		taskList.AddTask(description, status) // Use tasks.TaskList methods
		fmt.Printf("Task added: %s\n", description)
	case "list":
		tasks := taskList.GetTasks() // Use tasks.TaskList methods
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC3339), task.UpdatedAt.Format(time.RFC3339))
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Please provide a task ID and status.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		description := os.Args[3]
		status := os.Args[4]
		if taskList.UpdateTask(id, description, status) { // Use tasks.TaskList methods
			fmt.Printf("Task %d updated: %s, Status: %s\n", id, description, status)
		} else {
			fmt.Println("Task not found.")
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID to delete.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		if taskList.DeleteTask(id) { // Use tasks.TaskList methods
			fmt.Printf("Task %d deleted.\n", id)
		} else {
			fmt.Println("Task not found.")
		}
	case "get":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID.")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		task := taskList.GetTaskByID(id) // Use tasks.TaskList methods
		if task != nil {
			fmt.Printf("Task found: ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC3339), task.UpdatedAt.Format(time.RFC3339))
		} else {
			fmt.Println("Task not found.")
		}

	}

}
