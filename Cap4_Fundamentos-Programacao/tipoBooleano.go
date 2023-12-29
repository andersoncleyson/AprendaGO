package main

import "fmt"

var x bool

func main(){
	x = 10 == 100
	fmt.Println(x)
	x = 10 <= 100
	fmt.Println(x)
	x = 10 >= 100
	fmt.Println(x)
	x = 10 != 100
	fmt.Println(x)
	x = 10 > 100
	fmt.Println(x)
	x = 10 < 100
	fmt.Println(x)

}

/*
Tipo booleano

- Agora vamos explorar os tipos de maneira mais detalhada. ref/spec.
A começar pelo bool.
- O tipo bool é um tipo binário, que só pode conter um dos dois valores true e false.
- Sempre que você ver operadores relacionais (==, <=, >=, !=, <, >), o resuldatdo
da expressão será um valor booleano.
- Booleans são fundamentais nas tomadas de decisões em lógica condicional,
declarações switch, declarações if, fluxo de controle, etc.
- Na prática
	- zero value
	- Atribuindo um valor
	- Bool como resultado de operadores relacionais
*/