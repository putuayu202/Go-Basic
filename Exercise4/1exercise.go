package main

import (
	"fmt"
)

func main() {
	x := [5]int{0, 1, 2, 3, 4}

	for i, v := range x{
		fmt.Println("value of index", i, "is", v)
	}
}
