package main

import "time"

// UpdateTask updates the description and/or status of a task identified by its ID.
// It returns true if the task was found and updated, or false if the task was not found.
func (tl *TaskList) UpdateTask(id int, description string, status string) bool {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks[i].Description = description
			tl.Tasks[i].Status = status
			tl.Tasks[i].UpdatedAt = time.Now()
			tl.SaveTasks()
			return true
		}
	}
	return false
}
