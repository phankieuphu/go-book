package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// "github.com/patrick-blip/vocher-init/models"
	"github.com/patrick-blip/vocher-init/pkg/commond/models"
	"github.com/patrick-blip/vocher-init/pkg/commond/validate"
)

func FindBook(c *gin.Context){
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data" :books})
}
func CreateBook(c *gin.Context){
	var input validate.CreateBookvalidate
	if err:= c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	// Create books
	book := models.Book{Title: input.Title,Author: input.Author}
	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data" :book})
}
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err:= models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"erorr":"Recrod not found"})
	}
	var input validate.UpdateBookInput
	if err:= c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"erorr":err.Error()})
	}
	fmt.Println(input)
	// book.Author = input.Author
	// book.Title = input.Title
	// models.DB.
	updateBook := models.Book{Title:input.Title, Author:input.Author}
	models.DB.Model(&book).Updates(&updateBook)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	  return
	}
  
	models.DB.Delete(&book)
  
	c.JSON(http.StatusOK, gin.H{"data": true})
  }