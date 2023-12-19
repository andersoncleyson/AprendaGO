package main

import "fmt"

var a int
var b float64
var c string
var d bool

func main(){
    fmt.Printf("%v, %T", a, a)
    fmt.Printf("%v, %T", b, b)
    fmt.Printf("%v, %T", c, c)
    fmt.Printf("%v, %T", d, d)
}

/*
- Declaração vs. inicialização vs. atribuição de valor variáveis: caixas postais
- O que é valor zero
- Os zeros:
    - ints: 0
    - floats: 0
    - booleans: false
    - strings: ""
    - pointers, functions, interfaces, slices, channels, maps: nil
- Use := sempre que possível
- Use var para package-level scope
*/