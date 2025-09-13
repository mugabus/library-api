package routes

import (
	"library-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//books
	r.POST("/books", handlers.AddBookHandler)
	r.GET("/books", handlers.ListBooksHandler)

	//members

	r.POST("/members", handlers.AddMemberHandler)
	r.GET("/members", handlers.ListMembersHandler)
	//todo:Add Borrow/Return endpoints
	return r
}
