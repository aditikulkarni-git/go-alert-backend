package store 

var taskStore = make(map[string]Task)

func AddTask(task Task) {
	taskStore[task.ID] = task
}

func GetAllTasks() []Task {
	tasks := []Task{}
	for _, task := range taskStore {
		tasks = append(tasks, task)
	}
	return tasks
}
