package main

import "time"

// AddTask adds a new task to the task list with the given description.
// It initializes the task with a unique ID, sets its status to "todo",
// and records the current time as both CreatedAt and UpdatedAt timestamps.
func (tl *TaskList) AddTask(description string, status string) {
	task := Task{
		ID:          tl.NextID,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tl.Tasks = append(tl.Tasks, task)
	tl.NextID++
	tl.SaveTasks()
}
