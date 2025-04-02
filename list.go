package main

// Lists the tasks and their details.
func (tl *TaskList) GetTasks() []Task {
	return tl.Tasks
}

// GetTaskByID retrieves a task by its ID from the task list.
// It returns a pointer to the task if found, or nil if not found.
func (tl *TaskList) GetTaskByID(id int) *Task {
	for _, task := range tl.Tasks {
		if task.ID == id {
			return &task
		}
	}
	return nil
}
