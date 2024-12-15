package main

import (
	"fmt"

	"github.com/lamadev101/crudl/initializers"
	"github.com/lamadev101/crudl/models"
)

func init() {
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Blog{})
	fmt.Println("Successfully connected to the database!")
}
