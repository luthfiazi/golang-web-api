package main

import (

	"github.com/gin-gonic/gin"
	"golang-web-api/handler"
)

func main(){
	router :=gin.Default()

	//cara membuat versioning API V1
	v1 := router.Group("/v1")

	//membuat root url yang disarankan
	v1.GET("/", handler.RootHandler)  

	v1.GET("/artikel", handler.RootArtikel)

	v1.GET("/books/:id/:title/:author", handler.BookHandler)

	v1.GET("/query", handler.QueryHandler)

	v1.POST("/books", handler.PostBooksHandler)

	//cara membuat versioning API V2
	v2 := router.Group("/v2")

	//membuat root url yang disarankan
	v2.GET("/", handler.RootHandler)  

	v2.GET("/artikel", handler.RootArtikel)

	v2.GET("/books/:id/:title/:author", handler.BookHandler)

	v2.GET("/query", handler.QueryHandler)

	v2.POST("/books", handler.PostBooksHandler)

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





