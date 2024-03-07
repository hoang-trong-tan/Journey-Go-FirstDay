package main

import "fmt"

// package level
const ok = true

func main() {
	//func level
	var hello = "Hello"

	fmt.Println(hello, ok)
}
