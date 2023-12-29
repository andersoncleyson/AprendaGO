package main

import "fmt"

func main(){
	
	s := "Hello, friend"
	sb := []byte(s)
	fmt.Printf("%v\n%T\n", sb, sb)
	
	for _, v := range sb {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}
	fmt.Println()
	
}

/*
	- %v %T
	- Raw string literals
	- Conversão para slice of bytes: []byte(x)
	- %#U, %#x

*/