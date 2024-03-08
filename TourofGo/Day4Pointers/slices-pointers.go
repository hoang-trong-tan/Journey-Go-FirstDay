package main

import "fmt"

func main() {
	Heroes_girl := [3]string{
		"Rem",
		"Ram",
		"Emilia",
	}

	slices1 := Heroes_girl[:2]
	slices2 := Heroes_girl[:2]
	fmt.Println(slices1)
	slices1[1] = "Emilia"

	slices1 = Heroes_girl[:3] // extend length 2-> 3 , capacity : 3

	slices1[2] = "Ram"

	fmt.Println(slices2)
}
