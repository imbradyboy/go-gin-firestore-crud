package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imbradyboy/go-gin-firestore-crud/pkg/models"
	"google.golang.org/grpc/codes"
)

func GetAllJokes(c *gin.Context) {
	// fetch all jokes from model
	sliceOfJokes, err := models.GetAllJokes(c.Request.Context())
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// json encode data and send back as response
	c.IndentedJSON(http.StatusOK, sliceOfJokes)
}

func GetJokeById(c *gin.Context) {
	// get id from path params
	id := c.Param("id")

	// get joke from model
	joke, err := models.GetJokeById(id, c.Request.Context())
	if err != nil {
		// check if returned not found error, otherwise send back generic error
		httpStatus := http.StatusBadRequest
		if err.Error() == codes.NotFound.String() {
			httpStatus = http.StatusNotFound
		}

		c.IndentedJSON(httpStatus, gin.H{"message": err.Error()})
		return
	}

	// json encode data and send back as response
	c.IndentedJSON(http.StatusOK, joke)
}

func AddJoke(c *gin.Context) {
	// Call BindJSON to bind the received JSON to our variable
	var newJoke models.JokeDTO
	if err := c.BindJSON(&newJoke); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// add joke to db
	addedJoke, err := models.AddJoke(newJoke, c.Request.Context())
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// json encode data and send back as response
	c.IndentedJSON(http.StatusOK, addedJoke)
}

func UpdateJoke(c *gin.Context) {
	// get id from path params
	id := c.Param("id")

	// Call BindJSON to bind the received JSON to our variable
	var updatedJoke models.JokeDTO
	if err := c.BindJSON(&updatedJoke); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// update specified joke in db
	updatedJokeResponse, err := models.UpdateJoke(id, updatedJoke, c.Request.Context())
	if err != nil {
		// check if returned not found error, otherwise send back generic error
		httpStatus := http.StatusBadRequest
		if err.Error() == codes.NotFound.String() {
			httpStatus = http.StatusNotFound
		}

		c.IndentedJSON(httpStatus, gin.H{"message": err.Error()})
		return
	}

	// json encode data and send back as response
	c.IndentedJSON(http.StatusOK, updatedJokeResponse)
}

func DeleteJoke(c *gin.Context) {
	// get id from path params
	id := c.Param("id")

	// delete joke from db
	deleteResult, err := models.DeleteJoke(id, c.Request.Context())
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// json encode data and send back as response
	c.IndentedJSON(http.StatusOK, deleteResult)
}
