package main

import "fmt"

type Location struct {
	X, Y string
}

var reaoe = map[string]Location{
	"reaoe": {"the", "best"},
}

func main() {
	m := make(map[string]int)

	m["Answer"] = 42

	v, ok := reaoe["reaoe"]
	fmt.Println("The value:", v, "Presents?", ok)

	delete(reaoe, "reaoe")

	v, ok = reaoe["reaoe"]
	fmt.Println("The value:", v, "Presents?", ok)

}
