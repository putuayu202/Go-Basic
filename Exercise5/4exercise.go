package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first                       string
		last                        string
		favorite_ice_creame_flavors []string
	}{
		first:                       "dwiky",
		last:                        "danov",
		favorite_ice_creame_flavors: []string{"chocolate", "vanilla"},
	}
	p2 := struct {
		first                       string
		last                        string
		favorite_ice_creame_flavors []string
	}{
		first:                       "danov",
		last:                        "dwiky",
		favorite_ice_creame_flavors: []string{"coffe", "mint"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
