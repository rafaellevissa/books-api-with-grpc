package routes

import (
	"crud/internal/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.POST("/", controllers.CreateBook)
			books.GET("/", controllers.ShowAllBooks)
			books.GET("/:id", controllers.ShowBook)
			books.PUT("/:id", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}

	}
	return router
}
