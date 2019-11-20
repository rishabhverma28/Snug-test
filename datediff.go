package main

import (
	"os"
	"strconv"
	"strings"
)

const yearMIN int = 1901
const yearMAX int = 2999

var (
	monthDaysLeap = [...]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	monthDays     = [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	cummDays      = [...]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	cummDaysLeap  = [...]int{0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335}
)

// Creating the struct to represent the input dates
type dateType struct {
	day         int
	month       int
	year        int
	isLeapYear  bool
	sinceOrigin int
}

// Check if the day is valid
func (date *dateType) checkValidDay() bool {
	if date.isLeapYear {
		if date.day < 1 || date.day >= monthDaysLeap[date.month-1] {
			return false
		}
	} else {
		if date.day < 1 || date.day >= monthDays[date.month-1] {
			return false
		}
	}
	return true
}

// Check if the month and the year are valid
func (date *dateType) checkValidMonthYear() bool {
	if date.month < 1 || date.month > 12 {
		return false
	}
	if date.year < yearMIN || date.year > yearMAX {
		return false
	}
	validDay := date.checkValidDay()
	if !(validDay) {
		return false
	}
	return true
}

// Calculate the number of days from 01/01/0001
func (date *dateType) daysSinceOrigin() int {
	var presentDay int
	if date.isLeapYear {
		presentDay = date.day + cummDaysLeap[date.month-1]
	} else {
		presentDay = date.day + cummDays[date.month-1]
	}
	numLeapYears := (date.year-1)/4 - (date.year-1)/100 + (date.year-1)/400
	return (date.year-1)*365 + numLeapYears + presentDay - 1
}

// Check if the year is a leap year or not
func checkLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}

// Convert the string array into an integer array
func convertToInteger(date []string) [3]int {
	var date1List = [3]int{}
	k := 0
	for _, i := range date {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		date1List[k] = j
		k++
	}
	return date1List
}

// Check if the number date passed as the input is valid or not
func validateStringLength(date []string) bool {
	if len(date) != 3 || len(date[0]) != 2 || len(date[1]) != 2 || len(date[2]) != 4 {
		return false
	}
	return true
}

func minMax(a int, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a

}

func main() {
	// Make sure that the user inputs the correct number
	if len(os.Args) != 4 {
		panic("Insufficient arguments provided. \nPlease make sure the the arguments are in the format DD/MM/YYYY - DD/MM/YYYY.")
	}
	// We just want the dates
	date1 := os.Args[1]
	date2 := os.Args[3]

	// Split the dates based on the "/" separator
	date1ListString := strings.Split(date1, "/")
	date2ListString := strings.Split(date2, "/")

	// Validate the length of the date strings
	validString1 := validateStringLength(date1ListString)
	validString2 := validateStringLength(date2ListString)
	if !validString1 || !validString2 {
		panic("Please Enter the correct dates")
	}

	// Convert the string arrays to integer arrays
	var date1List = convertToInteger(date1ListString)
	var date2List = convertToInteger(date2ListString)

	// Initialize date1
	date1Type := dateType{
		day:        date1List[0],
		month:      date1List[1],
		year:       date1List[2],
		isLeapYear: checkLeapYear(date1List[2]),
	}
	// Initialize date2
	date2Type := dateType{
		day:        date2List[0],
		month:      date2List[1],
		year:       date2List[2],
		isLeapYear: checkLeapYear(date2List[2]),
	}

	// Check the validity of both the dates
	if !date1Type.checkValidMonthYear() || !date2Type.checkValidMonthYear() {
		panic("Invalid Dates. Please try again")
	}

	// Calculate the number of days since the origin
	date1Type.sinceOrigin = date1Type.daysSinceOrigin()
	date2Type.sinceOrigin = date2Type.daysSinceOrigin()

	// Print out the result
	maxVal, minVal := minMax(date2Type.sinceOrigin, date1Type.sinceOrigin)
	if maxVal-minVal == 0 {
		println(0)
	} else {
		println(maxVal - minVal - 1)
	}
}
