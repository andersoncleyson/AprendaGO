package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}

func (p *pessoa) falar(){
	fmt.Println(p.nome, "hello, friend")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(algumaCoisa humanos){
	algumaCoisa.falar()
}

func main(){
	p1 := pessoa{"Elliot", 28}

	p1.falar() // <- É um shortcut para (&p1).falar()

	(&p1).falar() // <- É a maneira mais correta

	dizerAlgumaCoisa(&p1) // <- Funciona!

	// dizerAlgumaCoisa(p1) <- Forma errada
}