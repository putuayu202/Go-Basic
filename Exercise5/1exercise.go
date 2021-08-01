package main

import (
	"fmt"
)

type person struct {
	first                       string
	last                        string
	favorite_ice_creame_flavors []string
}

func main() {
	p1 := person{
		first:                       "dwiky",
		last:                        "danov",
		favorite_ice_creame_flavors: []string{"chocolate", "vanilla"},
	}
	p2 := person{
		first:                       "danov",
		last:                        "dwiky",
		favorite_ice_creame_flavors: []string{"coffe", "mint"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
