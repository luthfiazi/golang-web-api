package handler

import (
	"fmt"
	"net/http"

	"golang-web-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//get API
func RootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"name": "Luthfi Azizi",
		"bio": "A story of coding",
	})
}

func RootArtikel(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"title": "Learning Golang",
		"content": "Learning Go language programming is fund and amazing",
	})
}

//get API by Id atau parameter variabel lebih dari 1
func BookHandler(c *gin.Context){
	id := c.Param("id")
	title := c.Param("title")
	author := c.Param("author")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"title": title,
		"author": author,
	})
}

//get API by query string variabel lebih dari 1
func QueryHandler(c *gin.Context){
	title := c.Query("title")
	description := c.Query("description")

	c.JSON(http.StatusOK, gin.H{"title":title, "description":description})
}

func PostBooksHandler(c *gin.Context){
	var bookInput book.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMessages := []string{}
		for _, e:= range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}