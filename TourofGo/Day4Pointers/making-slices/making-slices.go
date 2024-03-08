package main

import "fmt"

func main() {

	a := make([]int, 5) // len(a) = 5, cap = 5
	fmt.Println(cap(a))

	b := make([]int, 0, 5)

	b = b[:cap(b)]
	fmt.Println(b)

	b = b[1:]
	fmt.Println(b)
}
