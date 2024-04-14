package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
	var contador int64
	contador = 0
	totaldegoroutines := 100

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Goroutines", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Valor final:", contador)
}
