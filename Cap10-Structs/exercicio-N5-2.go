package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	sorvetes []string
}

func main(){
	pessoa1 := pessoa {
		nome: "Elliot", 
		sobrenome: "Alderson", 
		sorvetes: []string{
			"morango",
			"chocolate",
		},
	}

	pessoa2 := pessoa {
		nome: "Angela",
		sobrenome: "Moss",
		sorvetes: []string{
			"Pêssego",
			"Doce de leite",
		},
	}

	pessoas := map[string]pessoa{
		"Alderson": pessoa1,
		"Moss": pessoa2,
	}

	for key, value := range pessoas {	
		fmt.Println(key)
		for _, v := range value.sorvetes{
			fmt.Printf("\t%s\n", v)
		}
		
	}

}