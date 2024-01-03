package main

import "fmt"

func main(){

	quemtanoescritorio := "zezinho"

	switch quemtanoescritorio {
	case "zezinho":
		fmt.Println("hoje quem tá no escritório é o zezinho")
		fallthrough
	case "marquinhos":
		fmt.Println("sempre que o zezinho vem o marquinhos vem também")
	case "angela":
		fmt.Println("hoje quem tá no escritório é a angela")
	case "maria":
		fmt.Println("hoje que tá no escritório é a maria")
	default:
		fmt.Println("sei lá, meow")
		// isso é o padrão se não houver outra opção
	}
}