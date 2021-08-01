package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 12)
	x[0] = 10
	x[9] = 999
	fmt.Println(x)
	x = append(x, 11, 12, 13)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
