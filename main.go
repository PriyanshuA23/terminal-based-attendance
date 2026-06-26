package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const FILE_PATH = "attendance.json"

func displayHomePage(username string) {
	fmt.Println("Hello", username)
	fmt.Println("1. Mark Attendance")
	fmt.Println("2. Exit")
}

func checkIfAttendanceMarked(username string, records map[string][]string, date string) bool {
	fmt.Println(records)
	for _, name := range records[date] {
		if name == username {
			return true
		}
	}
	return false
}

func runAttendanceApplication(username string, records map[string][]string) {
	dateToday := time.Now().Format("2006-01-02")
	var isAttendanceMarked bool = checkIfAttendanceMarked(username, records, dateToday)

	for {
		displayHomePage(username)

		var userSelection string
		fmt.Scan(&userSelection)
		if userSelection == "2" {
			break
		}

		if isAttendanceMarked {
			fmt.Println("Attendance already marked 😡")
			continue
		}
		records[dateToday] = append(records[dateToday], username)
		isAttendanceMarked = true
		updatedRecord, _ := json.MarshalIndent(records, "", " ")
		os.WriteFile(FILE_PATH, updatedRecord, 0644)
		fmt.Println("Attendance is marked ✅")
		fmt.Println(records)
	}
}

func main() {
	attendanceRecords := make(map[string][]string)
	data, _ := os.ReadFile(FILE_PATH)
	json.Unmarshal(data, &attendanceRecords)

	for {
		fmt.Print("\033[H\033[2J")
		var username string
		fmt.Print("Enter username: ")
		fmt.Scan(&username)
		runAttendanceApplication(username, attendanceRecords)
	}
}
