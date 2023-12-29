package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main(){
    fmt.Printf("%v %T\n", x, x)
    x = 42
    fmt.Printf("%v %T\n", x, x)

    y = int(x)

    fmt.Printf("%v %T", y, y)
}

