package main

import (
	"fmt"
	"time"
)

type attendanceRecord struct {
	name string
	date string
}

func displayHomePage(username string) {
	fmt.Println("Hello", username)
	fmt.Println("1. Mark Attendance")
	fmt.Println("2. Exit")
}

func checkIfAttendanceMarked(username string, records *[]attendanceRecord) bool {
	fmt.Println(records)
	record := attendanceRecord{name: username, date: time.Now().Format("2006-01-02")}
	for _, value := range *records {
		if (value == record) {
			return true
		}
	}
	return false
}

func runAttendanceApplication(username string, records *[]attendanceRecord) {
	var isAttendanceMarked bool = checkIfAttendanceMarked(username, records)
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
		var record = attendanceRecord{name: username, date: time.Now().Format("2006-01-02")}
		*records = append(*records, record)
		isAttendanceMarked = true
		fmt.Println("Attendance is marked ✅")
		fmt.Println(records)
	}
}

func main() {
	var attendanceRecords []attendanceRecord
	for {
		fmt.Print("\033[H\033[2J")
		var username string
		fmt.Print("Enter username: ")
		fmt.Scan(&username)
		runAttendanceApplication(username, &attendanceRecords)
	}
}
