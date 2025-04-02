package main

func (tl *TaskList) GetTasks() []Task {
	return tl.Tasks
}
func (tl *TaskList) GetTaskByID(id int) *Task {
	for _, task := range tl.Tasks {
		if task.ID == id {
			return &task
		}
	}
	return nil
}
