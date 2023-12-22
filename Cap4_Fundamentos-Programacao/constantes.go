package main

import "fmt"

const x int = 10
const y = 10

var (
	a int
	b float64
	c bool
)
//var y = 10

//var c int
//var d float64

func main(){
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %v %v\n", a, b, c)
}