package main

import "fmt"

const (
	x = iota * 10
	y
	z
)

func main(){
	fmt.Println(x, y, z)
}

/*
	ref/spec#Iota
	- Numa declaração de constantes, o identificador iota
	representa números sequenciais.
	- Na prática.
		- iota, iota + 1, a = iota b c, reinicia em cada const, _
*/