package main

import "fmt"

func main(){
	slice := []int{1, 2, 3, 4, 5}

	valor := somaSlice(slice)

	fmt.Println(valor)
}

func somaSlice(x []int) int {
	soma := 0

	for _, v := range x {
		soma += v
	}

	return soma
}