package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 4, 7, 42}
	y := []int{4, 3, 1, 7, 69}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[0:2], x[4:]...)

	fmt.Println(x)
}
