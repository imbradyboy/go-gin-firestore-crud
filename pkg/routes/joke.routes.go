package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imbradyboy/go-gin-firestore-crud/pkg/controllers"
)

// manual route handler
func InitJokeRoutes(router *gin.Engine) {
	router.GET("/joke", controllers.GetAllJokes)
	router.GET("/joke/:id", controllers.GetJokeById)
	router.POST("/joke", controllers.AddJoke)
	router.PUT("/joke/:id", controllers.UpdateJoke)
	router.DELETE("/joke/:id", controllers.DeleteJoke)
}
