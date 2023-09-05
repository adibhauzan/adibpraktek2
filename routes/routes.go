package routes

import (
	"adib2/praktek2/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/books", controllers.AddBook)

	return r
}
