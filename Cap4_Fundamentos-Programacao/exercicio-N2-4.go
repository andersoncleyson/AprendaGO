package main

import "fmt"

var x int = 20

func main(){
	fmt.Printf("%d\t%b\t%#x\n", x ,x, x)

	y := x << 1
	fmt.Printf("y = %v\n", y)
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}