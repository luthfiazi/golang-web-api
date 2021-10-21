package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main(){
	router :=gin.Default()

	//versioning API V1
	v1 := router.Group("v1")

	//membuat root url yang disarankan
	v1.GET("/", rootHandler)  

	v1.GET("/artikel", rootArtikel)

	v1.GET("/books/:id/:title/:author", bookHandler)

	v1.GET("/query", queryHandler)

	v1.POST("/books", postBooksHandler)

	//versioning API V1
	v2 := router.Group("v2")

	//membuat root url yang disarankan
	v2.GET("/", rootHandler)  

	v2.GET("/artikel", rootArtikel)

	v2.GET("/books/:id/:title/:author", bookHandler)

	v2.GET("/query", queryHandler)

	v2.POST("/books", postBooksHandler)

	//--root url yang kurang disarankan--

	// router.GET("/", func(c *gin.Context)  {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"name": "Luthfi Azizi",
	// 		"bio": "A story of coding",
	// 	})
	// })

	// router.GET("/artikel", func(c *gin.Context)  {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"title": "Learning Golang",
	// 		"content": "Learning Go language programming is fund and amazing",
	// 	})
	// })

	router.Run(":9090")
} 

//memperbaiki root func yang lebih disarankan

//get API 
func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"name": "Luthfi Azizi",
		"bio": "A story of coding",
	})
}

func rootArtikel(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"title": "Learning Golang",
		"content": "Learning Go language programming is fund and amazing",
	})
}

//get API by Id atau parameter variabel lebih dari 1
func bookHandler(c *gin.Context){
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
func queryHandler(c *gin.Context){
	title := c.Query("title")
	description := c.Query("description")

	c.JSON(http.StatusOK, gin.H{"title":title, "description":description})
}

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int`json:"price" binding:"required,number"`
}

func postBooksHandler(c *gin.Context){
	var bookInput BookInput

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