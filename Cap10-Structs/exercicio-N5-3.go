package main

import "fmt"

type veiculo struct {
	portas int
	cor string
}

type caminhonete struct {
	veiculo
	quatroRodas bool

}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main(){
	caminhao := caminhonete {
		veiculo: veiculo {
			portas: 2,
			cor: "prata",
		},
		quatroRodas: true,
	}

	carro := sedan {
		veiculo: veiculo {
			portas: 4,
			cor: "preto",

		},
		modeloLuxo: true,
	}

	fmt.Println(carro.veiculo)
	fmt.Println(caminhao.veiculo)
}