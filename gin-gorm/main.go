package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinhbuihong/goweb/gin-gorm/controllers"
	"github.com/thinhbuihong/goweb/gin-gorm/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello"})
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	//group test group router
	test := r.Group("/test")
	{
		test.POST("/mul", controllers.Multipart)
		test.POST("/file", controllers.UpFile)
		test.GET("/read/:file", controllers.ReadFile)
	}

	r.Static("/upload", "./upload/")

	r.Run(":3000")
}
