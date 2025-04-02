package tasks

import "time"

// AddTask adds a new task to the task list with the given description and status.
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
