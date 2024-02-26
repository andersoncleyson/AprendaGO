package main

import "fmt"

func returnFunction() func(int) int {
	return func(i int) int { return 5 * i }
}
func main() {
	x := returnFunction()

	y := x(5)

	fmt.Println(y)
}
