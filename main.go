package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"echo-playground/services"
)

const totalSchoolDays = 180

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
	e.GET("/students/:id/sms-report", getStudentSMSReport)
	e.GET("/flagged-student-sms-reports", handleFlaggedStudentSMSReports)

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

func calculateAttendanceRate(absences int) string {
	if absences >= totalSchoolDays {
		return "0%"
	}
	attendanceRate := (float64(totalSchoolDays-absences) / float64(totalSchoolDays)) * 100
	return fmt.Sprintf("%.2f%%", math.Max(0, attendanceRate))
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

func getStudentReport(c echo.Context) error {
	studentData, err := readStudentData(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if studentData == nil {
		return c.JSON(http.StatusNotFound, "No student found")
	}
	reportPrompt := createEnhancedReportPrompt(studentData)
	response, err := services.ChatWithOpenAI(reportPrompt)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error generating report: %v", err))
	}
	return c.JSON(http.StatusOK, map[string]string{"report": response})
}

func createEnhancedReportPrompt(studentData map[string]string) string {
	return fmt.Sprintf("Generate a comprehensive student report for %s, a grade %s student in the school year %s. Include details on gender, average daily attendance (ADA) at %s%%, absence count of %s, infraction count of %s, GPA of %s, and other relevant academic and social performance indicators.",
		studentData["student_name"],
		studentData["grade_level"],
		studentData["school_year"],
		studentData["avg_daily_attendance"],
		studentData["absence_count"],
		studentData["infraction_count"],
		studentData["gpa"])
}

func getStudentSMSReport(c echo.Context) error {
	studentData, err := readStudentData(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if studentData == nil {
		return c.JSON(http.StatusNotFound, "No student found")
	}
	smsReport := createSMSReportPrompt(studentData)
	return c.String(http.StatusOK, smsReport)
}

func createSMSReportPrompt(studentData map[string]string) string {
	relationship := "child"
	if studentData["gender"] == "Male" {
		relationship = "son"
	} else if studentData["gender"] == "Female" {
		relationship = "daughter"
	}

	message := fmt.Sprintf("Hello! Just a quick update on your %s's school attendance. We've noted an attendance rate of %s with %s days absent and %s infractions recently. Please reach out if you need any support or information.",
		relationship, studentData["attendance_rate"], studentData["absence_count"], studentData["infraction_count"])
	log.Printf("Generated SMS report for student %s: %s", studentData["student_id"], message)
	return message
}

func handleFlaggedStudentSMSReports(c echo.Context) error {
	const (
		highAbsenceThreshold    = 20
		highInfractionThreshold = 5
	)
	records, err := readCSV("data/data.csv")
	if err != nil {
		log.Printf("Error reading CSV: %v", err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var smsReports []string
	for _, record := range records {
		absenceCount, _ := strconv.Atoi(record[5])
		infractionCount, _ := strconv.Atoi(record[6])
		if absenceCount > highAbsenceThreshold || infractionCount > highInfractionThreshold {
			attendanceRate := calculateAttendanceRate(absenceCount)

			studentData := map[string]string{
				"student_id":       record[0],
				"student_name":     record[1],
				"gender":           record[19],
				"absence_count":    record[5],
				"infraction_count": record[6],
				"attendance_rate":  attendanceRate,
			}
			log.Printf("Flagged student: %v", studentData)
			smsReport := createSMSReportPrompt(studentData)
			smsReports = append(smsReports, smsReport)
		}
	}

	if len(smsReports) == 0 {
		return c.JSON(http.StatusOK, "No flagged students found for SMS reports.")
	}
	return c.JSON(http.StatusOK, smsReports)
}

func handleFlaggedStudents(c echo.Context) error {
	const (
		highAbsenceThreshold    = 20
		highInfractionThreshold = 5
	)
	records, err := readCSV("data/data.csv")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var flaggedStudents []map[string]string
	for _, record := range records {
		absenceCount, _ := strconv.Atoi(record[5])    // Correct index for absence count
		infractionCount, _ := strconv.Atoi(record[6]) // Correct index for infraction count
		if absenceCount > highAbsenceThreshold || infractionCount > highInfractionThreshold {
			student := map[string]string{
				"student_id":       record[0],
				"student_name":     record[1],
				"gender":           record[19],
				"absence_count":    record[5],
				"infraction_count": record[6],
			}
			flaggedStudents = append(flaggedStudents, student)
		}
	}
	if len(flaggedStudents) == 0 {
		return c.JSON(http.StatusOK, "No flagged students found.")
	}
	return c.JSON(http.StatusOK, flaggedStudents)
}

func handleFlaggedStudentMessages(c echo.Context) error {
	const (
		highAbsenceThreshold    = 20
		highInfractionThreshold = 5
	)
	records, err := readCSV("data/data.csv")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var flaggedStudents []string
	for _, record := range records {
		absenceCount, _ := strconv.Atoi(record[5])    // Index for absence count
		infractionCount, _ := strconv.Atoi(record[6]) // Index for infraction count
		if absenceCount > highAbsenceThreshold || infractionCount > highInfractionThreshold {
			flaggedStudents = append(flaggedStudents, record[0]) // Appending student_id
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
	markdownOutput.WriteString("Personalized messages for flagged students:\n\n")
	for _, studentID := range flaggedStudents {
		personalizedMessage := strings.ReplaceAll(messageTemplate, "[Student's Name]", studentID)
		markdownOutput.WriteString(fmt.Sprintf("### Message for Student ID: %s\n\n%s\n\n---\n\n", studentID, personalizedMessage))
	}

	return c.String(http.StatusOK, markdownOutput.String())
}

func readStudentData(id string) (map[string]string, error) {
	records, err := readCSV("data/data.csv")
	if err != nil {
		log.Printf("Error reading CSV: %v", err)
		return nil, err
	}
	for _, record := range records {
		if record[0] == id {
			ada := record[4]
			if ada == "" {
				ada = "N/A" // Handle missing ADA
			} else {
				adaValue, err := strconv.ParseFloat(ada, 64)
				if err != nil {
					log.Printf("Error parsing ADA for student %s: %v", id, err)
					ada = "Invalid" // Handle invalid ADA
				} else {
					// Reformat ADA as a percentage string
					ada = fmt.Sprintf("%.2f%%", adaValue)
				}
			}

			studentData := map[string]string{
				"student_id":                    record[0],
				"student_name":                  record[1],
				"school_year":                   record[2],
				"grade_level":                   record[3],
				"avg_daily_attendance":          ada,
				"absence_count":                 record[5],
				"infraction_count":              record[6],
				"gpa":                           record[7],
				"participation_extracurricular": record[8],
				"parental_involvement_score":    record[9],
				"online_learning_engagement":    record[10],
				"reading_proficiency_level":     record[11],
				"math_proficiency_level":        record[12],
				"homework_submission_rate":      record[13],
				"class_participation_score":     record[14],
				"teacher_feedback_score":        record[15],
				"social_engagement_score":       record[16],
				"emotional_wellbeing_index":     record[17],
				"technology_proficiency_score":  record[18],
				"student_gender":                record[19],
			}
			log.Printf("Read student data for ID %s: %v", id, studentData)
			return studentData, nil
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
