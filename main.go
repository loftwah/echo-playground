package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"echo-playground/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	e := echo.New()

	e.GET("/", handleRoot)
	e.GET("/students/flagged", handleFlaggedStudents)

	e.Logger.Fatal(e.Start(":1323"))
}

func handleRoot(c echo.Context) error {
	response, err := services.ChatWithOpenAI("G'day, OpenAI!")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, response)
}

func handleFlaggedStudents(c echo.Context) error {
	const highAbsenceThreshold = 20

	fmt.Println("Opening CSV file...")
	f, err := os.Open("data/data.csv")
	if err != nil {
		fmt.Printf("Failed to open data file: %v\n", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to open data file: %v", err))
	}
	defer f.Close()

	fmt.Println("Reading CSV file...")
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read data file: %v\n", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to read data file: %v", err))
	}

	var flaggedStudents []string
	var responseMessage string
	needResponse := true

	for i, record := range records {
		if i == 0 { // Skip header line
			continue
		}

		studentID := record[0]
		absenceCount, _ := strconv.Atoi(record[3])

		if absenceCount > highAbsenceThreshold {
			flaggedStudents = append(flaggedStudents, studentID)

			// Generate message only once
			if needResponse {
				prompt := "Create a message indicating that a student has been absent too many times and may need additional support."
				fmt.Println(prompt)
				response, err := services.ChatWithOpenAI(prompt)
				if err != nil {
					fmt.Printf("Error generating message: %v\n", err)
					return c.String(http.StatusInternalServerError, fmt.Sprintf("Error generating message: %v", err))
				}
				responseMessage = response
				needResponse = false
			}
		}
	}

	if len(flaggedStudents) == 0 {
		fmt.Println("No flagged students found.")
		return c.JSON(http.StatusOK, "No flagged students found.")
	}

	fmt.Println("Returning flagged students with response message...")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"students": flaggedStudents,
		"message":  responseMessage,
	})
}
