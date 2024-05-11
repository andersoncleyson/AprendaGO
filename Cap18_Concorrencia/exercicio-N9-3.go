package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int

const quantidadeDeGoroutinas = 500

func main(){
	
	criarGoroutines(quantidadeDeGoroutinas)
	wg.Wait()
	fmt.Println("Total de goroutines: \t",
	quantidadeDeGoroutinas, "\nTotal do contador:\t", contador)
}

func criarGoroutines (i int){
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func(){
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
	
}