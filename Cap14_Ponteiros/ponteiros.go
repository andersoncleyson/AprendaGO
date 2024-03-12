package main

import "fmt"

func main() {
	x := 10

	y := &x

	*y++

	fmt.Println(*y)
	fmt.Printf("%T, %T\n", x, y)
	fmt.Print(x, y)
}
