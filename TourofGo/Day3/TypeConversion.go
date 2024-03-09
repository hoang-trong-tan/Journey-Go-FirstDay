package main

import (
	"fmt"
	"math"
)

func main() {

	const onlyone = 1

	var i int = 42
	var f float64 = math.Sqrt(float64(i))
	var u uint = uint(f)

	fmt.Println(onlyone)
	fmt.Println(i, f, u)
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j <= 5-i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
