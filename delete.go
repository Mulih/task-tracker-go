package main

func (tl *TaskList) DeleteTask(id int) bool {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:1], tl.Tasks[i+1:]...)
			tl.SaveTasks()
			return true
		}
	}
	return false
}
