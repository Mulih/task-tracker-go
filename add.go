package main

import "time"

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
