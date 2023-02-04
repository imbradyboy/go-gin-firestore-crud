package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imbradyboy/go-gin-firestore-crud/pkg/config"
	"github.com/imbradyboy/go-gin-firestore-crud/pkg/routes"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// connect to Firebase services
	config.InitializeFirebaseApp()

	// initialize gin and set up routes
	router := gin.Default()
	routes.InitJokeRoutes(router)

	router.Run("localhost:8080")
}
