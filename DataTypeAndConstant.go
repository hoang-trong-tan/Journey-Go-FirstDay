package main

import "fmt"

var gragas float64 = 50.5

const pi float32 = 3.14

func main() {
	var num int = 42
	var num2 uint8 = 10
	var num3 int64 = 1000

	var pi float32 = 3.14
	var gravity float64 = 9.81

	var complexNum complex64 = complex(3, 4)

	message := "Hello, World!"

	isTrue := true

	fmt.Println("Integer Example: ")
	fmt.Println(num, "\n", num2, "\n", num3)

	fmt.Println("Float Examples: ")
	fmt.Println(pi, "\n", gravity)

	fmt.Println("Complex Examples: ")
	fmt.Println(complexNum)

	fmt.Println("String Examples: ")
	fmt.Println(message)

	fmt.Println("boolean Examples: ")
	fmt.Println(isTrue)

	fmt.Println(gragas)

	var gragas int = int(gragas)

	fmt.Println(gragas)
}
