package store

import "go-alert-backend/internal/model"

var taskStore = make(map[string]model.Task)

func AddTask(task model.Task) {
	taskStore[task.ID] = task
}

func GetAllTasks() []model.Task {
	tasks := []model.Task{}
	for _, task := range taskStore {
		tasks = append(tasks, task)
	}
	return tasks
}
