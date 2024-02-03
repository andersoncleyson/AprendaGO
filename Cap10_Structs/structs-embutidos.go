package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}

type profissional struct {
	pessoa
	titulo string
	salario int
}

func main(){
	pessoa1 := pessoa{
		nome: "Elliot",
		idade: 28,
	}

	pessoa2 := profissional {
		pessoa: pessoa {
			nome: "Maricota",
			idade: 31,
		},
		titulo: "Pizzaoila",
		salario: 10000,
	}
	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.pessoa.nome)
}