package main

import "fmt"

var y = 59

func main(){
    z := 20
    qualquercoisa(z)
}

func qualquercoisa(x int) {
    fmt.Println(y)
    fmt.Println(x)
}