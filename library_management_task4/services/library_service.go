package services

import (
	"errors"
	"library_management/models"
	"slices"
	"sync"
)

type Library struct {
	Mu              sync.Mutex
	Books           map[int]models.Book
	Members         map[int]models.Member
	ReservationChan chan ReservationRequest
}

type ReservationRequest struct {
	BookID   int
	MemberID int
}

func NewLibrary(ch chan ReservationRequest) *Library {
	lib := Library{Books: make(map[int]models.Book),
		Members:         make(map[int]models.Member),
		ReservationChan: ch}
	return &lib
}

func (lib *Library) ReserveBook(bookID int, memberID int) error {
	lib.Mu.Lock()
	defer lib.Mu.Unlock()

	book, exist := lib.Books[bookID]
	if !exist {
		return errors.New("no book with this ID")
	}
	if book.Status != "available" {
		return errors.New("this book is not available")
	}

	if _, exist := lib.Members[memberID]; !exist {
		return errors.New(("member not found"))
	}

	book.Status = "reserved"
	lib.Books[bookID] = book
	lib.ReservationChan <- ReservationRequest{bookID, memberID}

	return nil
}

func (lib *Library) AddBook(book models.Book) {
	lib.Books[book.ID] = book
}

func (lib *Library) RemoveBook(id int) error {
	_, exist := lib.Books[id]
	if !exist {
		return errors.New("no book with this ID")
	}
	delete(lib.Books, id)
	return nil
}

func (lib *Library) BorrowBook(bookID int, memberID int) error {
	person, exist := lib.Members[memberID]
	if !exist {
		return errors.New(("member not found"))
	}

	book, exist := lib.Books[bookID]
	if !exist {
		return errors.New(("book not found"))
	}

	book.Status = "borrowed"
	person.BorrowedBooks = append(person.BorrowedBooks, book)

	lib.Books[bookID] = book
	lib.Members[memberID] = person

	return nil
}

func (lib *Library) ReturnBook(bookID int, memberID int) error {
	person, exist := lib.Members[memberID]
	if !exist {
		return errors.New(("member not found"))
	}

	book, exist := lib.Books[bookID]
	if !exist {
		return errors.New(("book not found"))
	}
	if book.Status != "borrowed" {
		return errors.New(("book is not borrowed"))
	}

	book.Status = "available"
	for idx, bk := range person.BorrowedBooks {
		if book.ID == bk.ID {
			person.BorrowedBooks = slices.Delete(person.BorrowedBooks, idx, idx+1)
			lib.Books[bookID] = book
			lib.Members[memberID] = person
			return nil
		}
	}

	return errors.New("member did not borrow this book")
}

func (lib *Library) ListAvailableBooks() []models.Book {
	available := []models.Book{}
	for _, book := range lib.Books {
		if book.Status == "available" {
			available = append(available, book)
		}
	}
	return available
}

func (lib *Library) ListBorrowedBooks(memberID int) []models.Book {
	person := lib.Members[memberID]
	return person.BorrowedBooks
}
