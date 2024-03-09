package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// Dai khai thi khai bao map se co tham so [type] = [struct] , map kh co keys, va kh the add keys

var n map[string]Vertex

func main() {
	// cho nen function map se tra ve type , va khoi tao duoc key :))
	n = make(map[string]Vertex)
	n["K814"] = Vertex{
		16.070247, 108.183382,
	}

	var place string
	fmt.Println("Ban muon tim toa do cua vi tri nao?")
	fmt.Scan(&place)

	fmt.Println(n[place])

}
