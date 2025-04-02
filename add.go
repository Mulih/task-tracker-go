package main

import "time"

// Function to add a new task to the task list.
// It creates a new Task object with the provided description, sets its status to "todo",
func (tl *TaskList) AddTask(description string) {
	task := Task{
		ID:          tl.NextID,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tl.Tasks = append(tl.Tasks, task)
	tl.NextID++
	tl.SaveTasks()
}
