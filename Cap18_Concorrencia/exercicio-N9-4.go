package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var contador int

const quantidadeDeGoroutinas = 5000

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
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
	
}