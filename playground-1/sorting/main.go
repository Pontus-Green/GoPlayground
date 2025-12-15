package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{5, 2, 1, 3}

	fmt.Println("unsorted arr", arr)
	slices.Sort(arr)
	slices.Reverse(arr)
	fmt.Println("sorted arr", arr)
}
