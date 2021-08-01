package main

import (
	"fmt"
)

func main() {
	x := map[string]int{
		"Lebron": 23,
		"harden": 13,
		"Curry":  30,
	}
	fmt.Println(x)
	fmt.Println(x["Curry"])


	if v, ok := x["harden"]; ok{
		fmt.Println("THIS SHOULD PRINT", v, ok)
	}

	x["Paul"] = 3

	for i, m := range x{
		fmt.Println(i, "is number", m)
	}

	delete(x, "harden")

	fmt.Println(x)

}
