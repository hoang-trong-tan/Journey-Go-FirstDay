package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		pic[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			pic[y][x] = uint8((x + y) / 2) // Thay đổi công thức tại đây để thử các hàm khác
		}
	}
	return pic
}

func main() {
	fmt.Println(Pic(5, 5))
}
