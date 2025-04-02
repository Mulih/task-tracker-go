package tasks

// DeleteTask removes a task from the task list by its ID.
func (tl *TaskList) DeleteTask(id int) bool {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			tl.SaveTasks()
			return true
		}
	}
	return false
}
