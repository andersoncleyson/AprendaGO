package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	sorvete []string

}

func main(){
	pessoa1 := pessoa {
		nome: "Elliot",
		sobrenome: "Alderson",
		sorvete: []string{
			"Chocolate",
			"Morango",
		},
	}

	pessoa2 := pessoa {
		nome: "Angela",
		sobrenome: "Moss",
		sorvete: []string{
			"Pêssego",
			"Doce de leite",
		},
	}

	fmt.Printf("Nome: %s\n", pessoa1.nome)
	fmt.Printf("Sobrenome: %s\n", pessoa1.sobrenome)
	fmt.Printf("Sabores de sorvete:\n")
	for _, value := range pessoa1.sorvete {
		fmt.Printf("\t%s\n", value)
	}

	fmt.Printf("Nome: %s\n", pessoa2.nome)
	fmt.Printf("Sobrenome: %s\n", pessoa2.sobrenome)
	fmt.Printf("Sabores de sorvete:\n")
	for _, value := range pessoa2.sorvete {
		fmt.Printf("\t%s\n", value)
	}

}