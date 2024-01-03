package main

import "fmt"

func main(){
	esporteFavorito := "futebol"

	switch esporteFavorito {
		case "volei":
			fmt.Println("Meu esporte favorito é vôlei")
		case "basquete":
			fmt.Println("Meu esporte favorito é basquete")
		case "futebol":
			fmt.Println("Meu esporte favorito é futebol")
		default:
			fmt.Println("Não gosto de esporte")
	}
}