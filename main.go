package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type dog struct {
	ID    string `json:"id"`
	Breed string `json:"breed"`
}

var dogs = []dog{
	{ID: "1", Breed: "Bulldog"},
	{ID: "2", Breed: "Poodle"},
	{ID: "3", Breed: "Basset"},
	{ID: "3", Breed: "Labrador Retriever"},
	{ID: "3", Breed: "Golden Retriever"},
}

func main() {
	router := gin.Default()
	router.GET("/dogs", get)
	router.GET("/dogs/:id", getById)
	router.POST("/dogs", post)

	router.Run("localhost:8080")
}

func get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dogs)
}

func post(c *gin.Context) {
	var newDog dog

	if err := c.BindJSON(&newDog); err != nil {
		return
	}

	dogs = append(dogs, newDog)
	c.IndentedJSON(http.StatusCreated, newDog)
}

func getById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range dogs {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "dog not found"})
}
