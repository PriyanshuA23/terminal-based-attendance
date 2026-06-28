package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"time"
)

const FILE_PATH = "attendance.json"

type AttendanceRecords map[string][]string

func displayHomePage(username string) {
	fmt.Println("Hello", username)
	fmt.Println("1. Mark Attendance")
	fmt.Println("2. Exit")
}

func checkIfAttendanceMarked(username string, records AttendanceRecords, date string) bool {
	fmt.Println(records)
	return slices.Contains(records[date], username)
}

func markAttendance(username string, records AttendanceRecords, todaysDate string) {
	records[todaysDate] = append(records[todaysDate], username)
	fmt.Println("Attendance is marked ✅")
}

func saveAttendance(records AttendanceRecords) {
	updatedRecord, _ := json.MarshalIndent(records, "", " ")
	os.WriteFile(FILE_PATH, updatedRecord, 0644)
	fmt.Println(records)
}

func runAttendanceApplication(username string, records AttendanceRecords) {
	dateToday := time.Now().Format("2006-01-02")
	var isAttendanceMarked bool = checkIfAttendanceMarked(username, records, dateToday)

	for {
		displayHomePage(username)

		var userSelection string
		fmt.Scan(&userSelection)

		if isAttendanceMarked {
			fmt.Println("Attendance already marked 😡")
			continue
		}

		switch userSelection {
		case "1":
			markAttendance(username, records, dateToday)
			saveAttendance(records)
			isAttendanceMarked = true

		case "2":
			return

		default:
			fmt.Println("Invalid option!")
		}
	}
}

func readUsername() string {
	fmt.Print("\033[H\033[2J")
	var username string
	fmt.Print("Enter username: ")
	fmt.Scan(&username)
	return username;
}

func main() {
	attendanceRecords := make(map[string][]string)
	data, _ := os.ReadFile(FILE_PATH)
	json.Unmarshal(data, &attendanceRecords)

	for {
		username := readUsername()
		runAttendanceApplication(username, attendanceRecords)
	}
}
