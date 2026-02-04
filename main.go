package main

import "fmt"

func main() {
	task1 := Task{
		ID:     "1",
		Title:  "Learn Go basics",
		Status: "OPEN",
	}

	task2 := Task{
		ID:     "2",
		Title:  "Build backend project",
		Status: "IN_PROGRESS",
	}

	AddTask(task1)
	AddTask(task2)

	fmt.Println("Tasks in system:")
	for _, t := range GetAllTasks() {
		fmt.Println(t)
	}
}
