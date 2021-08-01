package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 4, 7, 42}
	fmt.Println(x)

	for i, v := range x {
		fmt.Printf("Value of Index %v is %v\n", i, v)
	}
}
