package main

import "fmt"

func main(){
	x := 335

	func(x int){
		fmt.Println(x, "vezes 456")
		fmt.Println(x * 456)
	}(x)
}