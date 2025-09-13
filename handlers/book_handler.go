package handlers

import (
	"library-api/models"
	"library-api/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBookHandler(c *gin.Context) {
	var req models.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := storage.AddBook(req)
	c.JSON(http.StatusCreated, book)
}

func ListBooksHandler(c *gin.Context) {
	books := []models.Book{}
	for _, b := range storage.Books {
		books = append(books, b)
	}
	c.JSON(http.StatusOK, books)
}
