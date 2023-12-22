package main

import "fmt"

const (
	x = 2023 + iota
	y
	z
	k
)



func main(){
	fmt.Printf("%v\n%v\n%v\n%v\n", x, y, z, k)
}