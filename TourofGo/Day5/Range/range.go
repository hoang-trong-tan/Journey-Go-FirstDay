package main

import "fmt"

var reaoe = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range reaoe {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
