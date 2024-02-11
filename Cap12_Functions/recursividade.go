package main

import "fmt"

func main(){
	x := 5
	fmt.Println(loops(x))
}

func fatorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * fatorial(x - 1)
	}
}

func loops(x int) int {
	var y = x
	for y > 1 {
		x *= y - 1
		y = y - 1
	}
	return x
}