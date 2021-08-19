package main

import (
	"github.com/covalenthq/mq-store-agent/internal/gcp"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/cloud-storage-bucket", gcp.HandleFileUploadToBucket)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
