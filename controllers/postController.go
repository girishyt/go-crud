package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/girishyt/go-crud/initializers"
	"github.com/girishyt/go-crud/models"
)

func PostCreate(c *gin.Context) {
	type body struct {
		Title string
		Body  string
	}
	reqBody := body{}
	c.Bind(&reqBody)
	post := models.Post{Title: reqBody.Title, Body: reqBody.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Error(result.Error)
		return
	}
	fmt.Printf("Result: %+v", result)
	c.JSON(200, gin.H{
		"post": post,
	})

}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.Error(result.Error)
		return
	}
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(404, "Record with id "+id+" not found")
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePostById(c *gin.Context) {
	// Get ID
	id := c.Param("id")

	// Get Request Body
	var post models.Post
	var reqBody models.Post

	c.Bind(&reqBody)

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, "Record Not Found with the ID - "+id)
		return
	}

	initializers.DB.Model(&post).Updates(models.Post{Title: reqBody.Title, Body: reqBody.Body})

	c.JSON(200, gin.H{
		"post": post,
	})

}

func DeletePostById(c *gin.Context) {
	id := c.Param("id")
	tx := initializers.DB.Delete(&models.Post{}, id)
	if tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, "Record Not Found with the ID - "+id)
		return
	}
	c.JSON(200, "Deleted")
}
