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
}
