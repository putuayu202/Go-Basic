package main

import(
	"fmt"
)

func main(){
	bd := 2002
	for {
		if bd > 2021{
			break
		}
		fmt.Println(bd)
		bd++
	}	
}