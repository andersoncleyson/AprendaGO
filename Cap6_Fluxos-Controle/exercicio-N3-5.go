package main

import "fmt"

func main(){
	x := 10
	for x <= 100 {
		fmt.Println(x % 4)
		x++
	}
}