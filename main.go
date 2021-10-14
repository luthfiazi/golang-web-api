package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router :=gin.Default()

	//membuat root url yang disarankan
	router.GET("/", rootHandler)  

	router.GET("/artikel", rootArtikel)

	router.GET("/books/:id/:title/:author", bookHandler)

	router.GET("/query", queryHandler)

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