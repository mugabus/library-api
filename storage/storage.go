package storage

import (
	"errors"
	"library-api/models"
)

// in memory storage

var Books = make(map[int]models.Book)
var Members = make(map[int]models.Member)
var NextBookID = 1
var NextMemberID = 1

func AddBook(book models.Book) models.Book {
	book.BookID = NextBookID
	Books[NextBookID] = book
	NextBookID++
	return book
}
func AddMember(member models.Member) models.Member {
	member.MemberID = NextMemberID
	Members[NextMemberID] = member
	NextMemberID++
	return member
}

// Borrow book
func borrowBook(bookID int) error {
	book, exists := Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Borrowed {
		return errors.New("book already borrowed")
	}
	book.Borrowed = true
	Books[bookID] = book
	return nil
}
func ReturnBook(bookID int) error {
	book, exists := Books[bookID]
	if !exists {
		return errors.New("book not foound")
	}
	if !book.Borrowed {
		return errors.New("book was not borrowed")
	}
	book.Borrowed = false
	Books[bookID] = book
	return nil
}
