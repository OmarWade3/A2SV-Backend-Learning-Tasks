package data

import (
	"fmt"

	"github.com/zaahidali/task_manager_api/models"

	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//	var tasks = []models.Task{
//		{ID: 1, Title: "Task Manager Project", Description: "Add/View/Delete Tasks", DueDate: time.Now(), Status: "In Progress"},
//		{ID: 2, Title: "Books Management Project", Description: "Add/View/Delete Books", DueDate: time.Now().AddDate(0, 0, -1), Status: "Completed"},
//	}
var nextID int = 1 // auto-increment ID
var taskCollection *mongo.Collection

func InitDB() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}
	taskCollection = client.Database("Task_Manager").Collection("tasks")
}

// Get all tasks
func GetAllTasks() ([]models.Task, error) {
	allOptions := options.Find()
	var tasks []models.Task
	curr, err := taskCollection.Find(context.TODO(), bson.D{{}}, allOptions)
	if err != nil {
		return nil, err
	}

	for curr.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem models.Task
		err := curr.Decode(&elem)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, elem)
	}

	if err := curr.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

// Get a task by ID
func GetTaskByID(id int) (*models.Task, error) {
	var task models.Task

	filter := bson.D{{Key: "id", Value: id}}

	err := taskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// Create a new task
func CreateTask(task models.Task) (*models.Task, error) {
	task.ID = nextID
	nextID++
	_, err := taskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// Update an existing task
func UpdateTask(id int, updated models.Task) (*models.Task, error) {
	filter := bson.D{{Key: "id", Value: id}}

	updateFields := bson.D{}

	if updated.Title != "" {
		updateFields = append(updateFields, bson.E{Key: "title", Value: updated.Title})
	}

	if updated.Description != "" {
		updateFields = append(updateFields, bson.E{Key: "description", Value: updated.Description})
	}

	if updated.Status != "" {
		updateFields = append(updateFields, bson.E{Key: "status", Value: updated.Status})
	}

	// If no fields to update, return error:
	if len(updateFields) == 0 {
		return nil, fmt.Errorf("no fields provided for update")
	}

	update := bson.D{
		{Key: "$set", Value: updateFields},
	}

	_, err := taskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	// Return the updated task
	var task models.Task
	err = taskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// Delete a task by ID
func DeleteTask(id int) error {
	filter := bson.D{{Key: "id", Value: id}}
	_, err := taskCollection.DeleteOne(context.TODO(), filter)
	return err
}
