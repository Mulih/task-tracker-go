package tasks

import (
	"encoding/json"
	"os"
	"time"
)

// Task represents a task with an ID, description, status, and timestamps.
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TaskList contains a slice of tasks and the next available ID for new tasks.
type TaskList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

const taskFile = "tasks.json"

// SaveTasks saves the current task list to a JSON file.
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
