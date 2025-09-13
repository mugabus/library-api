package handlers

import (
	"library-api/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BorrowRequest struct {
	BookID   int `json:"book_id"`
	MemberID int `json:"member_id"`
}

type ReturnRequest struct {
	BookID int `json:"book_id"`
}

func BorrowBookHandler(c *gin.Context) {
	var req BorrowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := storage.BorrowBook(req.BookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book borrowed successfully"})
}

func ReturnBookHandler(c *gin.Context) {
	var req ReturnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := storage.ReturnBook(req.BookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}
