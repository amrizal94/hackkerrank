package main

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {
	hour, _ := strconv.Atoi(s[:len(s)-8])
	switch s[len(s)-2:] {
	case "PM":
		if hour == 12 {
			return s[:len(s)-2]
		}
		hour += 12
	default:
		if hour == 12 {
			hour = 0
		}
	}
	return fmt.Sprintf("%02d%s", hour, s[2:len(s)-2])
}

func main() {
	fmt.Println(timeConversion("07:05:45PM"))
}
