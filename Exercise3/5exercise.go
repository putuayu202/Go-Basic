package main

import(
	"fmt"
)

func main(){
	for i:= 10; i <= 100; i++{
		fmt.Printf("when %v devided by 4, the reminder is %v\n", i, i%4)
	}
}