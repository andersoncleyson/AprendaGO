package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mudeMe(p *pessoa) {
	(*p).nome = "Elliot"
	p.sobrenome = "Anderson"
}

func main() {
	pessoa := pessoa{"Anderson", "Silva", 28}

	fmt.Println(pessoa)
	mudeMe(&pessoa)
	fmt.Println(pessoa)

}
