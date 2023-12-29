package main

import "fmt"

type hotdog int

var b hotdog = 10

func main(){
    x := 10
    fmt.Printf("%v\n", x)
    fmt.Printf("%v\n", x)

    x = int(b) //Conversão
    fmt.Printf("%v", x)
}