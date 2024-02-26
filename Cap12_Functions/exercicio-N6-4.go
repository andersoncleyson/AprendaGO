package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p *pessoa) nomeCompleto() {
	fmt.Println(p.nome, p.sobrenome, p.idade)
}

func main() {
	pessoa1 := pessoa{"Elliot", "Anderson", 27}

	pessoa1.nomeCompleto()
}
