package main

import "fmt"

var x [5]int

func main(){
	for i := 0; i < 5; i++{
		x[i] = i + 1
	}
	
	y := 0
	for y < 5{
		fmt.Printf("%b\t%#U\n", x[y], x[y])
		y++
	}

	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	
}