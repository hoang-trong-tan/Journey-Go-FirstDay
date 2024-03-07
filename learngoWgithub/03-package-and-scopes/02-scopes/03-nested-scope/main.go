package main

import "fmt"

var declareMeAgain = 1

func nested() {
	var declareMeAgain = 2
	fmt.Println("Var change in nested func: ", declareMeAgain)
}

func main() {
	fmt.Println("Inside main", declareMeAgain)

	nested()

	// var declareMeAgain can't be effected

	fmt.Println("Inside main", declareMeAgain)

}
