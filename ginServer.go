package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type GinServer struct {
	Router *gin.Engine
}

// InstallCORS install CORS
func (s *GinServer) InstallCORS() {
	s.Router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
}

// Ping ping handler
func (s *GinServer) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Upload upload handler
func (s *GinServer) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Printf("File: %v", file.Filename)
	err := c.SaveUploadedFile(file, "static/"+file.Filename)
	if err != nil {
		c.JSON(400, gin.H{"message": "Error uploading file"})
	}
	c.JSON(200, gin.H{"message": "File uploaded successfully"})
}

// Run run server
func (s *GinServer) Run() {
	err := s.Router.Run(":8082")
	if err != nil {
		log.Printf("Error: %v", err)
	}
}

// InstallRoutes install routes
func (s *GinServer) InstallRoutes() {
	s.Router.GET("/api/ping", s.Ping)
	s.Router.POST("/api/upload", s.Upload)
}
