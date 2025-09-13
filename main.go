package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Book struct
type Book struct {
	BookID   int    `json:"book_id"`
	Title    string `json:"title"`
	Author   string `json: "author"`
	Borrowed bool   `json: "borrowed"`
}

// Member struct
type Member struct {
	MemberID int    `json: "member_id"`
	Name     string `json: "name"`
}

// in memory storage
var (
	nextBID = 1
	books   = make(map[int]Book)
	nextMID = 1
	members = make(map[int]Member)
)

func createBook(title, author string) Book {
	book := Book{
		BookID: nextBID,
		Title:  title,
		Author: author,
	}
	nextBID++
	books[book.BookID] = book
	return book
}

func createMember(name string) Member {
	member := Member{
		MemberID: nextMID,
		Name:     name,
	}
	nextMID++
	members[member.MemberID] = member
	return member
}

func borrowBook(BookID int, member Member) error {
	book, exists := books[BookID]
	if !exists {
		return errors.New("bok not found")
	}
	if book.Borrowed {
		return errors.New("book already borrowed")
	}
	book.Borrowed = true
	books[BookID] = book
	return nil
}

func returnBook(BookID int) error {
	book, exists := books[BookID]
	if !exists {
		return errors.New("book not found")
	}
	if !book.Borrowed {
		return errors.New("book was not borrowed")
	}
	book.Borrowed = false
	books[BookID] = book
	return nil
}

func main() {
	// Application entry point

	r := gin.Default()

	r.POST("/books", func(c *gin.Context) {
		var req struct {
			Title  string `json:"title"`
			Author string `json: "author"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		book := createBook(req.Title, req.Author)
		c.JSON(http.StatusCreated, book)
	})

	r.POST("/members", func(c *gin.Context) {
		var req struct {
			Name string `json: "name"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		member := createMember(req.Name)
		c.JSON(http.StatusCreated, member)
	})

	// Borrow book
	r.POST("/borrow", func(c *gin.Context) {
		var req struct {
			BookID   int `json:"book_id"`
			MemberID int `json:"member_id"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		member, exists := members[req.MemberID]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
			return
		}
		if err := borrowBook(req.BookID, member); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "book borrowed"})
	})

	r.POST("/return", func(c *gin.Context) {
		var req struct {
			BookID int `json:"book_id"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := returnBook(req.BookID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "book returned"})
	})
	// List all members
	r.GET("/members", func(c *gin.Context) {
		list := []Member{}
		for _, member := range members {
			list = append(list, member)
		}
		c.JSON(http.StatusOK, list)
	})
	//List all books

	r.GET("/books", func(ctx *gin.Context) {
		list := []Book{}
		for _, book := range books {
			list = append(list, book)
		}
		ctx.JSON(http.StatusOK, list)
	})

	r.Run(":8080")

}
