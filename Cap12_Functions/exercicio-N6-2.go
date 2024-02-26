package main

import "fmt"

func main() {
	slice := []int{2, 3, 5, 5}
	fmt.Println(soma(slice))
	fmt.Println((soma2(1, 2, 3)))
}

func soma(x []int) int {
	valor := 0
	for _, v := range x {
		valor += v
	}

	return valor
}

func soma2(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}

	return total
}
