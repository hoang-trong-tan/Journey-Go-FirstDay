package main

import (
	"fmt"
	"sort"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

var min int64
var max int64

func miniMaxSum(arr []int32) {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 0; i < len(arr)-1; i++ {
		min += int64(arr[i])
	}
	for i := 1; i < len(arr); i++ {
		max += int64(arr[i])
	}
	fmt.Println(min, max)
}

func main() {
	var a int32
	test := []int32{}
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		test = append(test, a)
	}
	miniMaxSum(test)

}
