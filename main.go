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
	e.Use(middleware.Logger(), middleware.Recover())

	e.GET("/", handleRoot)
	e.GET("/students/:id", getStudent)
	e.GET("/students/:id/report", getStudentReport)
	e.GET("/flagged-students", handleFlaggedStudents)
	e.GET("/flagged-student-messages", handleFlaggedStudentMessages)
	e.GET("/health", handleHealthCheck)

	e.Logger.Fatal(e.Start(":" + getServerPort()))
}

func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	return port
}

func handleRoot(c echo.Context) error {
	response, err := services.ChatWithOpenAI("G'day, OpenAI!")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, response)
}

func handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK", "message": "Service is operational"})
}

func getStudent(c echo.Context) error {
	studentData, err := readStudentData(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if studentData == nil {
		return c.JSON(http.StatusNotFound, "No student found")
	}
	return c.JSON(http.StatusOK, studentData)
}

// triggering a build
func getStudentReport(c echo.Context) error {
	studentData, err := readStudentData(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if studentData == nil {
		return c.JSON(http.StatusNotFound, "No student found")
	}
	prompt := createReportPrompt(studentData)
	response, err := services.ChatWithOpenAI(prompt)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error generating report: %v", err))
	}
	return c.JSON(http.StatusOK, map[string]string{"report": response})
}

func handleFlaggedStudents(c echo.Context) error {
	const highAbsenceThreshold = 20
	records, err := readCSV("data/data.csv")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	var flaggedStudents []string
	for _, record := range records {
		absenceCount, _ := strconv.Atoi(record[3])
		if absenceCount > highAbsenceThreshold {
			flaggedStudents = append(flaggedStudents, record[0])
		}
	}
	if len(flaggedStudents) == 0 {
		return c.JSON(http.StatusOK, "No flagged students found.")
	}
	return c.JSON(http.StatusOK, flaggedStudents)
}

func handleFlaggedStudentMessages(c echo.Context) error {
	const highAbsenceThreshold = 20
	records, err := readCSV("data/data.csv")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	var flaggedStudents []string
	for _, record := range records {
		absenceCount, _ := strconv.Atoi(record[3])
		if absenceCount > highAbsenceThreshold {
			flaggedStudents = append(flaggedStudents, record[0])
		}
	}
	if len(flaggedStudents) == 0 {
		return c.JSON(http.StatusOK, "No flagged students found.")
	}

	prompt := "Create a message template indicating that a student has been absent too many times and may need additional support."
	messageTemplate, err := services.ChatWithOpenAI(prompt)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error generating message template: %v", err))
	}

	var markdownOutput strings.Builder
	markdownOutput.WriteString("Personalized messages:\n\n")
	for _, studentID := range flaggedStudents {
		personalizedMessage := strings.ReplaceAll(messageTemplate, "[Student's Name]", studentID)
		markdownOutput.WriteString(fmt.Sprintf("### Message for Student ID: %s\n\n%s\n\n---\n\n", studentID, personalizedMessage))
	}

	return c.String(http.StatusOK, markdownOutput.String())
}

func readStudentData(id string) (map[string]string, error) {
	records, err := readCSV("data/data.csv")
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		if record[0] == id {
			return map[string]string{
				"student_id":           record[0],
				"school_year":          record[1],
				"avg_daily_attendance": record[2],
				"absence_count":        record[3],
				"infraction_count":     record[4],
			}, nil
		}
	}
	return nil, nil
}

func createReportPrompt(studentData map[string]string) string {
	return fmt.Sprintf("Based on the following student data: average daily attendance of %s%%, %s absences, and %s infractions in the school year %s, how is the student performing?",
		studentData["avg_daily_attendance"], studentData["absence_count"], studentData["infraction_count"], studentData["school_year"])
}

func readCSV(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open data file: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read data file: %v", err)
	}
	return records[1:], nil // Skip header
}
