package main

import "fmt"

func main() {
	defer fmt.Println("Eu sou o primeiro (EU ACHO QUE NÃO)")
	fmt.Println("Eu sou o segundo")
	fmt.Println("Eu sou o terceiro")
}
