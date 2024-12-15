package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lamadev101/crudl/initializers"
	"github.com/lamadev101/crudl/models"
)

func BlogCreate(c *gin.Context) {
	var requestBody struct {
		Title string
		Desc  string
		Image string
	}
	c.Bind(&requestBody)

	blog := models.Blog{
		Title: requestBody.Title,
		Desc:  requestBody.Desc,
		Image: requestBody.Image,
	}
	result := initializers.DB.Create(&blog)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"blog": blog})
}

func Blogs(c *gin.Context) {
	var blogs []models.Blog
	initializers.DB.Find(&blogs)
	c.JSON(200, gin.H{"blogs": blogs})
}

func Blog(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog
	initializers.DB.First(&blog, id)

	c.JSON(200, gin.H{"blog": blog})
}

func BlogUpdate(c *gin.Context) {
	id := c.Param("id")

	var requestBody struct {
		Title string
		Desc  string
		Image string
	}
	c.Bind(&requestBody)

	var blog models.Blog
	initializers.DB.First(&blog, id)

	result := initializers.DB.Model(&blog).Updates(models.Blog{
		Title: requestBody.Title,
		Desc:  requestBody.Desc,
		Image: requestBody.Image,
	})
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"blog": blog})
}

func BlogDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Blog{}, id)

	c.JSON(200, gin.H{"message": "Deleted successfully !!"})
}
