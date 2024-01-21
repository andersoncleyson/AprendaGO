package main

import "fmt"

func main(){
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatro queijos", "marg"}
	fatia := sabores[:]
	fmt.Println(fatia)

	i := 0

	for i < len(sabores) {
		fmt.Println(sabores[i])
		i++
	}

	sabores = append(sabores[:2], sabores[3:]...)

	fmt.Println(sabores)
}