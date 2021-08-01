package main

import (
	"fmt"
)

func main() {
	s := "hello Jancuk"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs:=[]byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i:=0; i<len(s); i++{
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("\n")

	for i, v:= range s {
		fmt.Printf("at index position %d we have Hex %x\n",i,v)
	}
}