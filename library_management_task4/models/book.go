package models

type Book struct {
	ID, ReservedBy        int
	Title, Author, Status string
}
