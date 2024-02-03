package main

import "fmt"

func main(){
	x := struct {
		nome string
		idade int
		cores map[string]string
		energeticos []int

	}{
		nome: "Elliot",
		idade: 28,
		cores: map[string]string{
			"Camisa": "Preta",
			"Calça": "Preta",
			"Sapato:": "Preto",
		},
		energeticos: []int{1, 2, 3, 4, 5},
	}

	fmt.Println(x)
}