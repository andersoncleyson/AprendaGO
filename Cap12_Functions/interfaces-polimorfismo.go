package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salario float64
}

type arquiteto struct {
	pessoa
	tipodeconstrucao string
	tamanhodaloucura string

}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e eu já arranquei ", x.dentesarrancados, " e ouve só: Bom dia")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	switch g.(type) {
		case dentista:
			g.oibomdia()
			fmt.Println("Eu ganho:", g.(dentista).salario)
		case arquiteto:
			g.oibomdia()
			fmt.Println("Eu construo:", g.(arquiteto).tipodeconstrucao)
	}
	
}

func main(){
	mrdente := dentista {
		pessoa: pessoa {
			nome: "Mister Dente",
			sobrenome: "da Silva",
			idade: 50,
		},
		dentesarrancados: 8987,
		salario: 10000.00,
	}

	mrpredio := arquiteto {
		pessoa: pessoa {
			nome: "Mister Prédio",
			sobrenome: "Soares",
			idade: 51,
		},
		tipodeconstrucao: "Hospícios",
		tamanhodaloucura: "Demais",
	}

	serhumano(mrdente)
	fmt.Println("")
	serhumano(mrpredio)

	
}