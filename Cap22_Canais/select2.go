package main

import (
	"fmt"
)

// Func 1 recebe X valores de canal, depois manda qualquer coisa para chan quit
// Func 2 for infinito, select: case envia pra canal, case recebe quit

func main(){
	canal := make(chan int)
	quit := make(chan int)
	go recebeQuit(canal, quit)
	enviaPraCanal(canal, quit)
}

func recebeQuit(canal chan int, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido:", <-canal)
	}
	quit <- 0
}

func enviaPraCanal(canal chan int, quit chan int){
	qualquercoisa := 1
	for {
		select {
		case canal <- qualquercoisa:
			qualquercoisa++
		case <-quit:
			return

		}
	}
}