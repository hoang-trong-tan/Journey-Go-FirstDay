package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("My name: Hoang Trong Tan", "Smo smo")
	fmt.Println("My bf name: ........", "Biu biu")

	buf := make([]byte, 8)

	// var EOF = errors.New("got error babei")
	// fmt.Println(EOF)

	r := strings.NewReader("some io.Reader stream to be read\n")
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
	r1 := strings.NewReader("First reader \n")
	// r2 := strings.NewReader("second reader \n")

	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	zac := 0
	for zac < 10 {
		fmt.Println(zac)
		zac++
	}

	message := []byte("hello, Gophers!")
	err := os.WriteFile("makeGopher", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
