package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	slice := []int{}

	rand.Seed(time.Now().UnixNano())

	n := 0

	for n <= 10 {
		slice = append(slice, (rand.Intn(100)))
		n++
	}

	for _, v := range slice {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}