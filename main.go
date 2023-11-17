package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"echo-playground/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	e := echo.New()
	e.Use(middleware.Logger())  // Added logging middleware
	e.Use(middleware.Recover()) // Recovery middleware

	// Routes
	e.GET("/", handleRoot)
	e.GET("/students/:id", getStudent)
	e.GET("/students/:id/report", getStudentReport)
	e.GET("/flagged-students", handleFlaggedStudents)
	e.GET("/flagged-student-messages", handleFlaggedStudentMessages)

	// Dynamic server port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323" // Default port if not set
	}
	e.Logger.Fatal(e.Start(":" + port))
}

func handleRoot(c echo.Context) error {
	response, err := services.ChatWithOpenAI("G'day, OpenAI!")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, response)
}

func getStudent(c echo.Context) error {
	// Fetching student ID from route parameter
	id := c.Param("id")

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

	var studentData map[string]string
	for i, record := range records {
		if i == 0 { // Skip header line
			continue
		}

		if record[0] == id {
			studentData = map[string]string{
				"student_id":           record[0],
				"school_year":          record[1],
				"avg_daily_attendance": record[2],
				"absence_count":        record[3],
				"infraction_count":     record[4],
			}
			break
		}
	}

	if studentData == nil {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("No student found with ID: %s", id))
	}

	return c.JSON(http.StatusOK, studentData)
}

func getStudentReport(c echo.Context) error {
	// Fetching student ID from route parameter
	id := c.Param("id")

	// Open and read the CSV file
	f, err := os.Open("data/data.csv")
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to open data file: %v", err))
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to read data file: %v", err))
	}

	var studentData map[string]string
	found := false
	for i, record := range records {
		if i == 0 { // Skip header line
			continue
		}

		if record[0] == id {
			studentData = map[string]string{
				"student_id":           record[0],
				"school_year":          record[1],
				"avg_daily_attendance": record[2],
				"absence_count":        record[3],
				"infraction_count":     record[4],
			}
			found = true
			break
		}
	}

	if !found {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("No student found with ID: %s", id))
	}

	// Generating a prompt for OpenAI
	prompt := fmt.Sprintf("Based on the following student data: average daily attendance of %s%%, %s absences, and %s infractions in the school year %s, how is the student performing?",
		studentData["avg_daily_attendance"], studentData["absence_count"], studentData["infraction_count"], studentData["school_year"])

	response, err := services.ChatWithOpenAI(prompt)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error generating report: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"report": response,
	})
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

func handleFlaggedStudentMessages(c echo.Context) error {
	const highAbsenceThreshold = 20

	fmt.Println("Opening CSV file for messages...")
	f, err := os.Open("data/data.csv")
	if err != nil {
		fmt.Printf("Failed to open data file: %v\n", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to open data file: %v", err))
	}
	defer f.Close()

	fmt.Println("Reading CSV file for messages...")
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read data file: %v\n", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to read data file: %v", err))
	}

	// Identify flagged students
	var flaggedStudents []string
	for i, record := range records {
		if i == 0 { // Skip header line
			continue
		}

		studentID := record[0]
		absenceCount, _ := strconv.Atoi(record[3])

		if absenceCount > highAbsenceThreshold {
			flaggedStudents = append(flaggedStudents, studentID)
		}
	}

	if len(flaggedStudents) == 0 {
		fmt.Println("No flagged students found.")
		return c.String(http.StatusOK, "No flagged students found.")
	}

	// Generate one message template
	prompt := "Create a message template indicating that a student has been absent too many times and may need additional support. Include placeholders for the student's name and other details."
	fmt.Println(prompt)
	messageTemplate, err := services.ChatWithOpenAI(prompt)
	if err != nil {
		fmt.Printf("Error generating message template: %v\n", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error generating message template: %v", err))
	}

	// Replace placeholders and create markdown output
	var markdownOutput strings.Builder
	markdownOutput.WriteString("Personalized messages:\n\n")

	for _, studentID := range flaggedStudents {
		personalizedMessage := strings.ReplaceAll(messageTemplate, "[Student's Name]", studentID)
		markdownOutput.WriteString(fmt.Sprintf("### Message for Student ID: %s\n\n%s\n\n---\n\n", studentID, personalizedMessage))
	}

	// Return the markdown output as part of the HTTP response
	return c.String(http.StatusOK, markdownOutput.String())
}
