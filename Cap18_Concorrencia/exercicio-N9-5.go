package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int32

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
			atomic.AddInt32(&contador, 1)
			v := atomic.LoadInt32(&contador)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()
	}
	
}