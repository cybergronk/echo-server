package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	routes "github.com/cybergronk/echo-server/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	var port string
	{
		// Exits when cli arguments don't match expectations.
		if len(os.Args) != 2 {
			fmt.Printf("Usage: %s <listening-port>\n", filepath.Base(os.Args[0]))
			os.Exit(1)
		}
		_, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Invalid 'port' value: %v\n", os.Args[1])
			os.Exit(1)
		}
		port = os.Args[1]
	}

	// Create a gin router with default configurations
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Redirect to health endpoint
	router.GET("/", routes.RedirectToHealth)

	// Health check endpoint
	router.GET("/health", routes.ReturnAliveStatus)

	// Echoes request body on any HTTP method
	router.Any("/echo", routes.EchoBody)

	// Starts listening
	router.Run(":" + port)
}
