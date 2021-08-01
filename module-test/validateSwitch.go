package main

import (
	"fmt"
)

func main() {
	a := "danov"
	switch a {
	case "anjay":
		fmt.Println("bukan")
	case "beruq":
		fmt.Println("bukan")
	case "danov":
		fmt.Println("print this danov")
	default:
		fmt.Println("ini default")
	}
	
}
