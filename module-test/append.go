package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 4, 7, 42}
	y := []int{4, 3, 1, 7, 69}
	x = append(x, y...)
	fmt.Println(x)

	x = x[0:4]

	fmt.Println(x)
}
