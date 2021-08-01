package main

import(
	"fmt"
)

func main(){
	switch{
	case false:
		fmt.Println("not print")
	case (2 == 4):
		fmt.Println("also not print")
	case (3 == 4):
		fmt.Println("print this")
		fallthrough
	case false:
		fmt.Println("print this also")
	default : 
		fmt.Println("ini default")
	}
	
}