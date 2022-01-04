package main

import (
	db "USRMGMT/controllers"
	"USRMGMT/routes"
	"fmt"

	"github.com/gin-contrib/cors"
)

func main() {

	// Create a new table.
	db.InitDatabase()

	// gin.DisableConsoleColor()

	// Initialize the routes
	r := routes.InitializeRoutes()

	//Allow CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	// r.Use(cors.Default())

	// Start the application server.
	fmt.Println("Server is running on port 5000")
	r.Run("localhost:5000")
}
