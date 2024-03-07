package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := NewGinServer(
		gin.Default(),
		NewInMemoryFileRepository(),
	)
	server.InstallCORS()
	server.InstallRoutes()
	server.Router.Static("/web", "./web/file-uploader/dist/")

	fmt.Printf("http://localhost:8082/web/index.html\n")
	err := server.Router.Run(":8082")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
