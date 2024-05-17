package main

import (
	"fmt"
)

func main(){
	canal := make(chan int, 1) //buffer
	
	go func(){
		canal <- 42
	}()

	fmt.Println(<-canal)
}