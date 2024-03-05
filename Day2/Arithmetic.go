package main

import "fmt"

func main() {

	var a, b, c int = 11, 5, 0

	c = a + b
	fmt.Println("a + b =", c)

	c = a - b
	fmt.Println("a - b = ", c)

	c = a * b
	fmt.Println("a * b = ", c)

	c = a / b
	fmt.Println("a / b = ", c)

	c = a % b

	fmt.Println("a % b =", c)

	a++

	fmt.Println("a++ = ", a)

	b--

	fmt.Println("b-- = ", b)

}
