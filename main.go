package main

import (
	"fmt"
	"time"
)

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
		fmt.Println("Attendance is marked ✅")
		fmt.Println(records)
	}
}

func main() {
	attendanceRecords := map[string][]string {}
	for {
		fmt.Print("\033[H\033[2J")
		var username string
		fmt.Print("Enter username: ")
		fmt.Scan(&username)
		runAttendanceApplication(username, attendanceRecords)
	}
}
