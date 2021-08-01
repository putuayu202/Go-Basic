package main

import(
	"fmt"
)

var x int
var y string
var z bool

func main(){
	x = 42
	y = "James Bond"
	z = true

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	s := fmt.Sprintf("%v\t%v\t%v\t", x,y,z)
	fmt.Println(s)
}