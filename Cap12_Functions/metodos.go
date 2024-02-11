package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}

func (p pessoa) oibomdia(){
	fmt.Println(p.nome, "diz bom dia")
}

func main(){
	elliot := pessoa{"Elliot", 28}

	elliot.oibomdia()
}