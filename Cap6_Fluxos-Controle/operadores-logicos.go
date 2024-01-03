package main

import "fmt"

func main(){
	x := 3

	if (x == 2) || ( x == 3){
		fmt.Println("x é igual a 3")
	}

	if (x % 2 == 0) && (x % 3 == 0){
		fmt.Println("é multiplo de dois e também de 3")
	}
}