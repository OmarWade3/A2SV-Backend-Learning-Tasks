# Library Management System — Documentation

## 1. Introduction

The **Library Management System** is a console-based application built using the Go programming language.
It demonstrates clean code organization using packages, modular design, struct-based models, and separation of concerns through controllers and services.

The system allows users to manage library books and members through a simple text-based menu.
Core operations include adding books, borrowing books, returning books, and viewing book availability.

---

## 2. Features / Console Interactions

The system uses a simple command-line interface. When the program runs, the user is presented with a menu offering the following options:

### **1. Add a New Book**

* Prompts the user to enter a **Book ID**, **Title**, and **Author**.
* Creates a new book and stores it with a default status of `"Available"`.

### **2. Remove an Existing Book**

* Asks for the **Book ID** to remove.
* Validates book existence.
* Removes the book from the library.

### **3. Borrow a Book**

* Requests a **Book ID** and a **Member ID**.
* Ensures:

  * The book exists
  * The member exists
  * The book is available
* Marks the book as `"Borrowed"`.
* Adds the book to the member’s `BorrowedBooks` list.

### **4. Return a Book**

* Requires the user to enter:

  * Book ID
  * Member ID
* Confirms:

  * The member exists
  * The book exists
  * The member has borrowed the book
* Marks the book as `"Available"` again.
* Removes the book from the member’s borrowed list.

### **5. List All Available Books**

* Displays every book whose `Status == "Available"`.
* Shows the book’s ID, title, and author.

### **6. List All Borrowed Books by a Member**

* Asks for a **Member ID**.
* Displays all books currently borrowed by that member.

### **7. Add a Member**

* Creates a new library member by entering:

  * Member ID
  * Member Name

### **8. Exit**

* Terminates the application.

---

## 3. Folder Structure Explanation

The project follows a clean, modular structure:

```
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   ├── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod
```

### **main.go**

* Entry point of the application.
* Calls the library controller to start the console interface.

---

### **controllers/**

Contains user interaction and console-input logic.

#### `library_controller.go`

* Displays the menu.
* Reads user input.
* Delegates logic to the service layer.
* Ensures separation between UI and business logic.

---

### **models/**

Defines the data models used by the system.

#### `book.go`

Contains the `Book` struct

#### `member.go`

Contains the `Member` struct

---

### **services/**

Implements business logic and data handling.

#### `library_service.go`

Responsible for managing:

* Books map
* Members map

Implements key functions:

* `AddBook(book Book)`
* `RemoveBook(id int)`
* `BorrowBook(bookID, memberID int)`
* `ReturnBook(bookID, memberID int)`
* `ListAvailableBooks() []Book`
* `ListBorrowedBooks(memberID int) []Book`

This ensures controllers do not manipulate data directly.

---

### **docs/**

Contains documentation files for the project.

#### `documentation.md`

This file — explains:

* Program purpose
* How to interact with it
* Folder structure
* Features and design decisions

---

## 4. How to Run the Program

Follow these steps to run the application:

### **1. Navigate to the project folder**

```sh
cd library_management
```

### **2. Run the program**

```sh
go run main.go
```

### **3. Use the menu**

Choose options (1–8) to interact with the system.

---

## 5. Design Considerations & Architecture

### **Separation of Concerns**

The system is broken into:

* **Models** → define data
* **Services** → handle business logic
* **Controllers** → manage UI and input/output

This structure makes the application:

* Easier to maintain
* More readable
* Easier to extend

### **Maps for Data Storage**

Books and Members are stored using:

* `map[int]Book`
* `map[int]Member`

This gives:

* Fast lookup by ID
* Easy addition and removal
* Simple simulation of a real database

### **Console UI**

The controller uses:

* `bufio.Reader`
* `fmt.Println`
* `strconv.Atoi`

This makes the application simple and intuitive for console interaction.

---

## 6. Conclusion

This Library Management System showcases:

* Go modular design
* Basic CRUD operations
* Struct-based modeling
* Console interaction
* Proper layering using controllers and services