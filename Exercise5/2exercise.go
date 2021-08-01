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

	p3 := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range p3{
		fmt.Printf("%v \n \t first name : %v\n \t last name : %v\n \t favorite ice cream flavors : \n", k, v.first, v.last)
		for i, ice:= range v.favorite_ice_creame_flavors {
			fmt.Println("\t\t\t\t", i+1, ice)
		}
	}

}
