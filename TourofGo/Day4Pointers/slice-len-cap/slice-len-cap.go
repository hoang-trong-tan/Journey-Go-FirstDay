package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 9, 11, 13}
	// slice has both length and a capacity
	printSlice(s)

	// Slice the slice to give it zero length len:7 -> len:0
	s = s[:0]
	printSlice(s)

	// extend length: 4
	s = s[:4]
	printSlice(s)

	// Drop it First two values : decrease cap=7 -> cap=5 . len= 4 -> 2
	s = s[4:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
