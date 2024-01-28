package main

import "fmt"

func main(){
	pessoas := [][]string {
		[]string{"Elliot", "Alderson", "Hackear"},
		[]string{"Shayla", "Nico", "Vender arte na praia"},
		[]string{"Angela", "Moss", "Trabalhar na Ecorp"},
	}

	for i, _ := range pessoas {
		for v, _ := range pessoas {
			fmt.Printf("%s ", pessoas[i][v])
			
		}
		fmt.Printf("\n")
	}
}