package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	// Serve static files from Vue.js 'dist' directory
	r.Static("/", "./web/file-uploader/dist/")

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Printf("File: %v", file.Filename)
		err := c.SaveUploadedFile(file, "static/"+file.Filename)
		if err != nil {
			c.JSON(400, gin.H{"message": "Error uploading file"})
		}
		c.JSON(200, gin.H{"message": "File uploaded successfully"})
	})

	fmt.Printf("http://localhost:8082\n")
	err := r.Run(":8082")
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
