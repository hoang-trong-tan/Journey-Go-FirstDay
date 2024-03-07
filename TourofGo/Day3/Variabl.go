package main

import "fmt"

func shortfunc(a, b, c int) int {
	return a + b + c
}

func remultivalue(a, b int) (int, int) {
	return b, a
}

func nakeretrn(x int) (a int, b int) {
	a = x - 1
	b = x - 2
	return
}

var (
	Tobe  bool   = false
	Reaoe string = "Reaoe"
	bad   byte   = '1'
	rud   rune   = 50
)

func main() {

	fmt.Println("Short function like this ", shortfunc(1, 2, 3))
	fmt.Println(remultivalue(1, 2))
	fmt.Println(nakeretrn(10))
	fmt.Println(bad, rud)
}
