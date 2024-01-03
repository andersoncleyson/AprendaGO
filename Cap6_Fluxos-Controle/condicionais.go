package main

import "fmt"

func main(){
	if x := 10; x > 100{
		fmt.Println("X é maior que 100")
	} else if x < 100{
		fmt.Println("x não é maior que 100")
	} else {
		fmt.Println("Sei não")
	}
}