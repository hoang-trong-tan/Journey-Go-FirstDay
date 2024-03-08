package main

import "fmt"

func main() {
	// cach tao slice
	var slice = []string{} // slice nil

	fmt.Println(slice)

	slice1 := []string{}

	slice1 = append(slice1, "pt1", "pt2")

	fmt.Println(slice1) // print pt1 pt2

	slice2 := slice1[1:]

	fmt.Println(slice2, len(slice2), cap(slice2)) // print pt2 , length 1, cap : 1

	slicewMake := make([]int, 5) // create slice []int length =5 , cap : 5

	fmt.Println(slicewMake, len(slicewMake), cap(slicewMake))

}
