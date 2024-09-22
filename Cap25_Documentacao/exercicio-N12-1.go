package main

import "fmt"

func main(){
	var idade int

	fmt.Printf("Qual idade do seu cachorro? ")
	fmt.Scan(&idade)

	fmt.Print("A idade do seu cachorro em idade humana é: ", idadeCachorro(idade), "\n")
}

func idadeCachorro(i int) int {
	return i * 7
}
