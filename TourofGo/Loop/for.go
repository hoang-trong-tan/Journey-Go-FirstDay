package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 3000000000; i++ {
		sum += i
	}
	fmt.Println(sum)

	for { // infinity for
		sum += sum
	}
}
