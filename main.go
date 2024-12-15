package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lamadev101/crudl/controllers"
	"github.com/lamadev101/crudl/initializers"
)

func init() {
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/blog", controllers.BlogCreate)
	r.PUT("/blog/:id", controllers.BlogUpdate)
	r.GET("/blogs", controllers.Blogs)
	r.GET("/blog/:id", controllers.Blog)
	r.DELETE("/blog/:id", controllers.BlogDelete)

	r.Run()
}
