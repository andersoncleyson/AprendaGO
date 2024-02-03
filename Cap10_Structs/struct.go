package main

import "fmt"

type cliente struct {
	nome string
	sobrenome string
	fumante bool
}

func main(){
	cliente1 := cliente {
		nome: "Angela",
		sobrenome: "Moss",
		fumante: false,
	}

	cliente2 := cliente {"Shayla", "Nico", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)

}