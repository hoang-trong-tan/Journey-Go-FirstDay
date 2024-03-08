package main

import "fmt"

var Remstruct struct {
	remv1 int
	remv2 bool
}

func main() {
	Emilia := []string{"thich", "khong thich", "thich", "khong thich"}

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
	}

	fmt.Println(Emilia)
	fmt.Println(s)
}
