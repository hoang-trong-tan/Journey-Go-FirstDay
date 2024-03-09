package main

import "fmt"

// create a struct ?

type reaoe struct {
	X, Y, Z int
}

// how can you access struct?

func main() {
	x := reaoe{1, 2, 3}

	fmt.Println("field x:", x.X)

	// how can you pointer a struct?

	y := &x

	y.X = 3 // this was change reaoe struct X fields to 3

	fmt.Println(x)

	// define only X on struct ?

	test_define_struct := reaoe{X: 10}

	fmt.Println(test_define_struct)

	// Can you give me an array with name anime girls :)
	// okay

	anime_girls_name := [3]string{
		"Emilia",
		"Rem",
		"Ram",
	}

	fmt.Println(anime_girls_name)

	// Can you give me a slice with name anime girls :)

	slice_name_girls := []string{"Emilia", "Frieren", "Fern", "Ram", "Rem", "Est"}
	fmt.Println(slice_name_girls)

	get_name_girls := slice_name_girls[:2]
	fmt.Println("Get first and second name girls ", get_name_girls)

	// change first girl names pls
	get_name_girls[0] = "Rin"
	fmt.Println(get_name_girls[:1])

	// Please create me a slice with make

	slice_with_make := make([]int, 5, 5)
	fmt.Println(slice_with_make)
	slice_with_make = append(slice_with_make, 1, 2, 3, 4, 5)
	fmt.Println(slice_with_make)

	// create slice in slice

	slice_in_slice := [][]string{
		[]string{"_", "_", "_"},
	}
	fmt.Println(slice_in_slice)
}
