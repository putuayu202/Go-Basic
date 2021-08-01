package main

import (
	"fmt"
)

type vehicle struct {
	door  int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			door:  2,
			color: "yellow",
		},
		fourWheel: false,
	}

	s1 := sedan{
		vehicle: vehicle{
			door:  4,
			color: "silver",
		},
		luxury: true,
	}

	fmt.Println(t1, s1)
}
