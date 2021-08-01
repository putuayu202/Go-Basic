package main

import (
	"fmt"
)

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := x[0:5]
	b := x[5:10]
	c := x[3:8]
	d := x[2:7]
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
