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

//memperbaiki root func 
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