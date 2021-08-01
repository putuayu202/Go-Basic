package main

import (
	"fmt"
)

func main() {
	sim := "punya"
	nyetir := "jago"
	
	bisa := sim == "punya"
	gakbisa := sim == "gak punya"

	sabi := nyetir == "jago"
	gaksabi := nyetir == "gak jago"
	
	if bisa && sabi {
		fmt.Println("lu bisa nyetir")
	} else if bisa && gaksabi {
		fmt.Println("lu gak bisa nyetir")
	} else if bisa || sabi {
		fmt.Println("lu bisa nyetir")
	} else if gakbisa || sabi {
		fmt.Println("lu bisa nyetir")
	} else if !bisa {
		fmt.Println("lu gak bisa nyetir")
	}
}
