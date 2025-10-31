package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/models"
	"library_management/services"
)

func RunLibrary() {
	reader := bufio.NewReader(os.Stdin)
	library := services.NewLibrary()

	for {
		fmt.Println("\n===== Library Management System =====")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove a book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List available books")
		fmt.Println("6. List borrowed books by member")
		fmt.Println("7. Add a member")
		fmt.Println("8. Exit")
		fmt.Println("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			addBook(library, reader)
		case 2:
			removeBook(library, reader)
		case 3:
			borrowBook(library, reader)
		case 4:
			returnBook(library, reader)
		case 5:
			listAvailableBooks(library)
		case 6:
			listBorrowedBooks(library, reader)
		case 7:
			addMember(library, reader)
		case 8:
			fmt.Println("Exiting the system. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func addBook(library *services.Library, reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	id := readInt(reader)

	fmt.Print("Enter Book Title: ")
	title := readString(reader)

	fmt.Print("Enter Book Author: ")
	author := readString(reader)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "available",
	}

	library.AddBook(book)
	fmt.Println("Book added successfully!")
}

func removeBook(library *services.Library, reader *bufio.Reader) {
	fmt.Print("Enter Book ID to remove: ")
	id := readInt(reader)

	err := library.RemoveBook(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Book removed successfully!")
}

func borrowBook(library *services.Library, reader *bufio.Reader) {
	fmt.Print("Enter Book ID to borrow: ")
	bookID := readInt(reader)

	fmt.Print("Enter Member ID: ")
	memberID := readInt(reader)

	err := library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func returnBook(library *services.Library, reader *bufio.Reader) {
	fmt.Print("Enter Book ID to return: ")
	bookID := readInt(reader)

	fmt.Print("Enter Member ID: ")
	memberID := readInt(reader)

	err := library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func listAvailableBooks(library *services.Library) {
	books := library.ListAvailableBooks()

	fmt.Println("\nAvailable Books:")
	for _, b := range books {
		fmt.Printf("Book ID: %d Title: %s Author: %s\n", b.ID, b.Title, b.Author)
	}
}

func listBorrowedBooks(library *services.Library, reader *bufio.Reader) {
	fmt.Print("Enter Member ID: ")
	memberID := readInt(reader)

	books := library.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books found for this member.")
		return
	}

	fmt.Println("\nBorrowed Books:")
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}

func addMember(library *services.Library, reader *bufio.Reader) {
	fmt.Print("Enter Member ID: ")
	id := readInt(reader)

	fmt.Print("Enter Member Name: ")
	name := readString(reader)

	member := models.Member{
		ID:   id,
		Name: name,
	}

	library.Members[id] = member
	fmt.Println("Member added successfully!")
}

func readInt(reader *bufio.Reader) int {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err == nil {
			return val
		}
		fmt.Print("Invalid number, try again: ")
	}
}

func readString(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
