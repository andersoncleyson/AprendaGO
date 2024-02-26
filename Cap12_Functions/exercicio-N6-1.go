package main

import "fmt"

func main() {
	fmt.Println(retornaInt())
	fmt.Println(retornaIntString())
}

func retornaInt() int {
	return 15
}

func retornaIntString() (int, string) {
	return 10, "Hello, friend"
}
