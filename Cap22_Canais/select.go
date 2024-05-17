package main

import (
	"fmt"
)

// - Duas funcs enviando X numeros cada uma pra um canal
// - For loop X valores, select case <-x

func main(){

	a := make(chan int)
	b := make(chan int)
	x := 500

	go func(x int){
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x/2)

	go func(x int){
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x/2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("Canal A:", v)
		case v := <-b:
			fmt.Println("Canal B:", v)
		}
	}
}