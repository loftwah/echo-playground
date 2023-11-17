package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	// Replace with the correct import path for your services package
	"echo-playground/services"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Create a new instance of Echo
	e := echo.New()

	// Define routes
	e.GET("/", handleRoot)

	// Start the server on port 1323
	e.Logger.Fatal(e.Start(":1323"))
}

// handleRoot is the route handler for the root endpoint
func handleRoot(c echo.Context) error {
	// Send a prompt to OpenAI and get the response
	response, err := services.ChatWithOpenAI("G'day, OpenAI!")
	if err != nil {
		// If there's an error, return a 500 internal server error response
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Return the response from OpenAI
	return c.String(http.StatusOK, response)
}
