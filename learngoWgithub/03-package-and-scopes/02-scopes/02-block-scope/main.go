package main

import "fmt"

func eachlv() {
	var hello = "Hello"
	const ok = true

	// but visible here
	fmt.Println(hello, ok)
}

func main() {

	// hello and ok not visible here
	eachlv()
	// fmt.Println(hello, ok)
}
