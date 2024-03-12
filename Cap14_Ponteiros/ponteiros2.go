package main

import "fmt"

func main() {
	x := 0

	estareceovalor(x)
	estarecebeumponteiro(&x)
}

func estareceovalor(x int) {
	x++
	fmt.Println(x)
}

func estarecebeumponteiro(x *int) {
	*x++
	fmt.Println(*x)
}
