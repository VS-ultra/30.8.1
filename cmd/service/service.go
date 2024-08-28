package main

import (
	"fmt"
	"log"

	"30.8.1/pkg/storage"
)

func main() {
	databaseUrl := fmt.Sprintf("postgres://postgres:newpassword@localhost:5432/tasks?connect_timeout=10&sslmode=prefer")

	store, err := storage.New(databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer store.Close()

	// Я не уверен нужена ли эта часть кода тк ее не было в ТЗ для ее работы нужена быть заполнена БД
	newTask := storage.Task{
		AuthorID:   1,
		AssignedID: 2,
		Title:      "New Task",
		Content:    "This is a new task",
	}
	taskID, err := store.NewTask(newTask)
	if err != nil {
		log.Fatalf("Unable to create new task: %v\n", err)
	}
	fmt.Printf("New task created with ID: %d\n", taskID)
	tasks, err := store.AllTasks()
	if err != nil {
		log.Fatalf("Unable to get all tasks: %v\n", err)
	}
	fmt.Println("All tasks:", tasks)

	updatedTask := storage.Task{
		ID:      taskID,
		Title:   "Updated Task",
		Content: "This is an updated task",
	}
	err = store.UpdateTask(updatedTask)
	if err != nil {
		log.Fatalf("Unable to update task: %v\n", err)
	}
	fmt.Println("Task updated")

	err = store.DeleteTask(taskID)
	if err != nil {
		log.Fatalf("Unable to delete task: %v\n", err)
	}
	fmt.Println("Task deleted")
}
