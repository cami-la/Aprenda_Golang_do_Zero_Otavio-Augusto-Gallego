package main

import "fmt"

func dayOfTheWeek(day int) string {
	switch day {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid day"
	}
}

func dayOfTheWeek2(day int) string {
	var result string

	switch {
	case day == 1:
		result = "Sunday"
		//fallthrough
	case day == 2:
		result = "Monday"
	case day == 3:
		result = "Tuesday"
	case day == 4:
		result = "Wednesday"
	case day == 5:
		result = "Thursday"
	case day == 6:
		result = "Friday"
	case day == 7:
		result = "Saturday"
	default:
		result = "Invalid day"
	}

	return result
}

func main() {
	day := dayOfTheWeek(3)
	fmt.Println(day)

	fmt.Println("------------------")
	fmt.Println(dayOfTheWeek2(5))
}
