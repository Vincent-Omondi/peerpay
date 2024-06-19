// main.go
package main

import (
	"log"
	"net/http"
	"path/filepath"

	"peerpay/backend/database"
	"peerpay/backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Specify the path to the .env file
	err := godotenv.Load(filepath.Join("..", ".env")) // Adjust the path if necessary
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the database
	database.Connect()
	database.Migrate()

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Trust specific proxies
	router.SetTrustedProxies(nil)

	// Register routes
	routes.AuthRoutes(router)
	routes.ProfileRoutes(router)
	routes.SharesRoutes(router)
	routes.AdminRoutes(router)

	// Print all routes
	for _, route := range router.Routes() {
		log.Printf("Route: %s %s", route.Method, route.Path)
	}

	// Handle 404 errors
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"Peerpay": "Not Found"})
	})

	log.Println("Starting server on :8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
