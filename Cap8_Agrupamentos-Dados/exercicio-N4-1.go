package main

import (
	"fmt"
	"math/rand"
	"time"
	
)

var array [5]int
func main(){
	
	rand.Seed(time.Now().UnixNano())

	for i, _ := range slice {
		array[i] = rand.Intn(100)
		fmt.Printf("%d %d\n", i, array[i])
	}

	fmt.Printf("%T\n", array)
}