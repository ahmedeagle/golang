package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Struct for Gin binding
// Example 1: Using `binding` for validation

type FeedbackGin struct {
	Feedback string `json:"feedback" binding:"required,min=1"`
}

// Struct for Validator.v10
// Example 2: Using `validate` for validation

type FeedbackValidator struct {
	Feedback string `json:"feedback" validate:"required,min=1"`
}

func main() {
	r := gin.Default()

	// Example 1: Gin binding validation
	r.POST("/gin-binding", func(c *gin.Context) {

		var feedback FeedbackGin

		if err := c.ShouldBindJSON(&feedback); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Validated with Gin."})
	})

	// Example 2: Using go-playground/validator for manual validation
	validate := validator.New()
	r.POST("/validator", func(c *gin.Context) {
		var feedback FeedbackValidator

		if err := c.ShouldBindJSON(&feedback); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(&feedback); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Validated with go-playground/validator."})
	})

	r.Run() // Start the Gin server
}
