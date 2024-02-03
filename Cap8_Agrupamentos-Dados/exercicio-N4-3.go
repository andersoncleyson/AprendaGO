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

	for n < 10 {
		slice = append(slice, rand.Intn(100))
		n++
	}
	fmt.Println(slice)
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:8])
	fmt.Println(slice[2:len(slice) - 1])
}