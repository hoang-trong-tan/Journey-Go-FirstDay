package main

import "fmt"

var anime = []string{"em1", "em2", "em3"}

func main() {
	for i, _ := range anime {
		fmt.Println(i)
	}

	for _, v := range anime {
		fmt.Println(v)
	}
}
