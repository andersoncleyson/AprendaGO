package main

import "fmt"

func main(){
	x := 0

	for x < 10{
		x++
		if (x % 2 != 0){
			continue
		} else {
			fmt.Println(x)
		}
	}
}