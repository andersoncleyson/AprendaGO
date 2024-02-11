package main

import "fmt"

func main(){
	x := 335

	y := func(x int) int{
		//fmt.Println(x, "vezes 456")
		return x * 456
	}
	fmt.Println(x, "vezes 456 é ", y(x))
	
}