package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("not Print")

	case 1 == 2:
		fmt.Println("not Print")
	case 4 == 4:
		fmt.Println("Print it god dammit")
	case true:
		fmt.Println("also Print it god dammit")
	}
}
