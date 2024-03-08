package main

import "fmt"

func main() {
	Ar7 := [7]int{1, 2, 3, 4, 5, 6, 7}

	var Slices []int = Ar7[:7]
	fmt.Println("len: ", len(Slices), " cap: ", cap(Slices))
}
