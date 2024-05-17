package main

import (
	"fmt"
)

func main(){
	canal := make(chan int)
	go send(canal)

	recieve(canal)
}

func send(s chan<- int){
	s <- 42
}

func recieve(r <-chan int){
	fmt.Println("O valor recebido do canal foi: ", <-r)
}