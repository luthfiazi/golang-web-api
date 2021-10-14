package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router :=gin.Default()

	//membuat root url
	router.GET("/", func(c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"name": "Luthfi Azizi",
			"bio": "A story of coding",
		})
	})

	router.Run()
} 