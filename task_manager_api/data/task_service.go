package data

import (
	"errors"
	"time"

	"github.com/zaahidali/task_manager_api/models"
)

var tasks = []models.Task{
	{ID: 1, Title: "Task Manager Project", Description: "Add/View/Delete Tasks", DueDate: time.Now(), Status: "In Progress"},
	{ID: 2, Title: "Books Management Project", Description: "Add/View/Delete Books", DueDate: time.Now().AddDate(0, 0, -1), Status: "Completed"},
}
var nextID int = 3 // auto-increment ID

// Get all tasks
func GetAllTasks() []models.Task {
	return tasks
}

// Get a task by ID
func GetTaskByID(id int) (*models.Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

// Create a new task
func CreateTask(task models.Task) models.Task {
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	return task
}

// Update an existing task
func UpdateTask(id int, updated models.Task) (*models.Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			// Update fields
			if updated.Title != "" {
				tasks[i].Title = updated.Title
			}
			if updated.Description != "" {
				tasks[i].Description = updated.Description
			}
			if updated.Status != "" {
				tasks[i].Status = updated.Status
			}

			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

// Delete a task by ID
func DeleteTask(id int) error {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
