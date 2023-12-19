package main

import "fmt"

var hello = "Hello"

func main(){
	name := "Anderson" //Declaração
	age := 28  //Declaração

	fmt.Printf("%s %s\n Age: %d\n", hello, name, age)

	name, age = "Elliot", 29 //Atribuição

	fmt.Printf("%s\n%d", name, age)

}

/*
	Terminologia:
	- keywords (palavras-chave) são termos reservados
	- operadores, operandos
	- statement (declaração, afirmação) -> uma linha de código, uma instrução
	que forma uma ação, e é geralmente formado de expressões
	- expressão -> qualquer coisa que "produz um resultado"

	:= utilizado para criar novas variáveis, dentro de code blocks
	= para atribuir valores a variáveis já existentes
*/