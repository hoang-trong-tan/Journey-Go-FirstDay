package main

import "fmt"

type Location struct {
	X, Y float64
}

var Cmap = map[string]Location{
	"K814":   {16.070247, 108.183382},
	"Google": {16.070247, 108.183382},
}

func main() {
	fmt.Println(Cmap)
}
