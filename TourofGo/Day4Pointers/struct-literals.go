package main

import "fmt"

type VerTex struct {
	X, Y int
}

var (
	v1 = VerTex{1, 2}
	v2 = VerTex{X: 5}
	v3 = VerTex{}
	p  = &VerTex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
