package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	num := math.Max(2, 10)

	switch num {
	case 10:
		fmt.Println("got 10")
	default:
		fmt.Println("nothing")
	}
}
