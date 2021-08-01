package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Like", "Vanilla"}
	lbj := []string{"Lebron", "James", "Like", "Choco"}

	fusion := [][]string{jb, lbj}
	fmt.Println(fusion)
}
