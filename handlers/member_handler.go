package handlers

import (
	"library-api/models"
	"library-api/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddMemberHandler(c *gin.Context) {
	var req models.Member
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	member := storage.AddMember(req)
	c.JSON(http.StatusCreated, member)
}

func ListMembersHandler(c *gin.Context) {
	members := []models.Member{}
	for _, m := range storage.Members {
		members = append(members, m)
	}
	c.JSON(http.StatusOK, members)
}
