package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main(){
    s := fmt.Sprintf("%v %s %v", x, y, z)

    fmt.Printf("%v %T", s, s)
}