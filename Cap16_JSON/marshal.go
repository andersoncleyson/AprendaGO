package main

import (
	"encoding/json"
	"fmt"
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
		Contabancaria: 100000.50,
	}
	mrrobot := pessoa{
		Nome:          "Elliot",
		Sobrenome:     "Alderson",
		Idade:         27,
		Profissao:     "Security Enginner",
		Contabancaria: 2000.00,
	}

	j, err := json.MarshalIndent(jamesbond, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	m, err := json.MarshalIndent(mrrobot, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	os.Stdout.Write(j)
	os.Stdout.Write(m)

}
