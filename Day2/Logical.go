package main

import "fmt"

func main() {
	var isWeekend = true
	var hour = 10
	if (isWeekend) && (hour >= 10) && (hour <= 20) {
		fmt.Println("We are Open")
	} else {
		fmt.Println("We are closed")
	}

	if (hour < 10) || (hour > 20) || (isWeekend) {
		fmt.Println("Store is Closed")
	} else {
		fmt.Println("Store is Open")
	}

	if !isWeekend {
		fmt.Println("It's a work day")
	} else {
		fmt.Println("It's the weekend")
	}
}
