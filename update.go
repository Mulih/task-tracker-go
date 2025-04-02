package main

import "time"

func (tl *TaskList) UpdateTask(id int, description string, status string) bool {
	task := tl.GetTaskByID(id)
	if task == nil {
		return false
	}

	if description != "" {
		task.Description = description
	}
	if status != "" {
		task.Status = status
	}
	task.UpdatedAt = time.Now()
	tl.SaveTasks()
	return true
}
