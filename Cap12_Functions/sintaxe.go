package main

import "fmt"

func main(){
	basica()
	argumento("tarde")
	argumento("asdf")
	argumento("manhã")
	total, quantos := soma(1, 2, 3)
	fmt.Println(total, quantos)
}

func basica (){
	fmt.Println("hello, friend")
}

func argumento (x string) {
	if x == "manhã" {
		fmt.Println("Bom dia")
	} else if x == "tarde" {
		fmt.Println("Boa tarde")
	} else {
		fmt.Println("Boa noite")
	}
}

func soma (x ...int) (int, int) {
	soma := 0
	for _, v := range x {
		soma += v
	}

	return soma, len(x)
}


//func (receiver) nome(parameters) (returns) {code}