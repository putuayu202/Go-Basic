package main

import (
	"fmt"
)

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"miss", "MoneyPenny", "Hello Bond"}

	fmt.Println([][]string{x, y})
}
