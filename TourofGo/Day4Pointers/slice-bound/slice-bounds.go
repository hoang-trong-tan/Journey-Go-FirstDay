package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	rem := s[:2] // it start like this [0:2]
	fmt.Println(rem)
	rem = s[1:] // it start like this [1:lengthofarray] this case length = 6
	fmt.Println(rem)
}
