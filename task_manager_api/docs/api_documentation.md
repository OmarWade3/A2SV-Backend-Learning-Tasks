# Task Management API Documentation

## Overview

The Task Management API is a RESTful service built using Go and the Gin framework. It provides basic CRUD operations for managing tasks stored in an in-memory database.

Base URL:

```
http://localhost:8080
```

---

# Endpoints

## 1. Get All Tasks

Retrieve a list of all available tasks.

**Endpoint**

```
GET /tasks
```

**Response (200 OK)**

```json
[
  {
    "id": 1,
    "title": "Buy groceries",
    "description": "Milk, Bread, Eggs",
    "due_date": "2025-01-16",
    "status": "pending"
  }
]
```

---

## 2. Get Task by ID

Retrieve a single task by providing its unique ID.

**Endpoint**

```
GET /tasks/:id
```

**Example**

```
GET /tasks/1
```

**Response (200 OK)**

```json
{
  "id": 1,
  "title": "Buy groceries",
  "description": "Milk, Bread, Eggs",
  "due_date": "2025-01-16",
  "status": "pending"
}
```

**Response (404 Not Found)**

```json
{
  "error": "task not found"
}
```

---

## 3. Create Task

Create a new task by sending a JSON request body.

**Endpoint**

```
POST /tasks
```

**Request Body**

```json
{
  "title": "Clean the room",
  "description": "Organize desk and vacuum floor",
  "due_date": "2025-01-20",
  "status": "in progress"
}
```

**Response (201 Created)**

```json
{
  "id": 2,
  "title": "Clean the room",
  "description": "Organize desk and vacuum floor",
  "due_date": "2025-01-20",
  "status": "in progress"
}
```

**Response (400 Bad Request)**

```json
{
  "error": "invalid request body"
}
```

---

## 4. Update Task

Update an existing task using its ID.

**Endpoint**

```
PUT /tasks/:id
```

**Example**

```
PUT /tasks/1
```

**Request Body**

```json
{
  "title": "Buy groceries and supplies",
  "description": "Milk, Bread, Eggs, Soap",
  "due_date": "2025-01-18",
  "status": "pending"
}
```

**Response (200 OK)**

```json
{
  "id": 1,
  "title": "Buy groceries and supplies",
  "description": "Milk, Bread, Eggs, Soap",
  "due_date": "2025-01-18",
  "status": "pending"
}
```

**Response (404 Not Found)**

```json
{
  "error": "task not found"
}
```

**Response (400 Bad Request)**

```json
{
  "error": "invalid request body"
}
```

---

## 5. Delete Task

Delete a task by ID.

**Endpoint**

```
DELETE /tasks/:id
```

**Example**

```
DELETE /tasks/1
```

**Response (200 OK)**

```json
{
  "message": "task deleted successfully"
}
```

**Response (404 Not Found)**

```json
{
  "error": "task not found"
}
```

---

# Data Model

## Task Object

```json
{
  "id": 1,
  "title": "string",
  "description": "string",
  "due_date": "string",
  "status": "string"
}
```

Field descriptions:

* `id`: Auto-increment integer assigned by the system
* `title`: Title of the task
* `description`: Additional details about the task
* `due_date`: Deadline or target completion date
* `status`: Task state such as "pending", "in progress", or "done"

---

# Example cURL Commands

## Get all tasks

```
curl -X GET http://localhost:8080/tasks
```

## Get a task by ID

```
curl -X GET http://localhost:8080/tasks/1
```

## Create a task

```
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title": "New Task", "description": "Details", "due_date": "2025-01-20", "status": "pending"}'
```

## Update a task

```
curl -X PUT http://localhost:8080/tasks/1 \
-H "Content-Type: application/json" \
-d '{"title": "Updated", "description": "Updated desc", "due_date": "2025-01-30", "status": "done"}'
```

## Delete a task

```
curl -X DELETE http://localhost:8080/tasks/1
```

---

# Notes

* All data is stored in memory and will reset when the server restarts.
* Input validation is minimal since the goal of this exercise is to practice API structure, routing, and controllers.