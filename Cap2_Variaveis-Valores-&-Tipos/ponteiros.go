package main

import "fmt"

func main(){
	x := 10
	p := &x

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", *p)

	*p = 15
	fmt.Printf("%v\n", x)
}