package main

import (
	"encoding/json"
	"os"
)

type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {
	jamesbond := pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente Secreto",
		Contabancaria: 10000.50,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesbond)
}
