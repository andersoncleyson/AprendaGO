package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main(){

}

func trabalho(s string) chan string {
	canal := make(chan string)
	go func(s string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, canal)
	return canal
}