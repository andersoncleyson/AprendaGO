package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

var wg sync.WaitGroup

func main(){
	fmt.Println("Quantidade de CPU:", runtime.NumCPU())
	fmt.Println("Número de goroutines:", runtime.NumGoroutine())
	
	wg.Add(2)

	go func1()
	go func2()

	fmt.Println("Número de goroutines:", runtime.NumGoroutine())
	
	wg.Wait()

	fmt.Println()
}

func func1(){
	for i := 0; i < 10; i++ {
		fmt.Println("func1", i)
		time.Sleep(20)
	}
	wg.Done()
}

func func2(){
	for i := 0; i < 10; i++ {
		fmt.Println("func2", i)
		time.Sleep(20)
	}
	wg.Done()
}