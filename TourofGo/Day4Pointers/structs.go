package main

import "fmt"

type reaoe struct {
	X, Y int
}

func main() {
	v := reaoe{1, 2}
	v.X = 2

	p := &v
	*&p.X = 50
	fmt.Println(v)
}
