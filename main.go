package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main () {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main Site",
		})
	})
	r.Run(":8080")
}
