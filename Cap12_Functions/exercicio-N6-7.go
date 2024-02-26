package main

import "fmt"

func main() {
	x := 50

	func(x int) {
		fmt.Println("x vezes 100 igual a:")
		fmt.Println(x * 100)
	}(x)
}
