package main

import (
	"fmt"
)

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}
