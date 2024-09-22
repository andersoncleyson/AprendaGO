package main

import "fmt"

func main(){
	x := soma(1, 2, 3, 4, 5)
	fmt.Println(x)
}

func soma (i ...int) int {
	total := 0

	for _, v := range i {
		total += v

	}

	return total
}

func multiplica(i ...int) {
	total := 0

	for _, v := range i {
		total *= v
	}

	return total
}
