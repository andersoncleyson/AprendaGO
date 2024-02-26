package main

import "fmt"

func qualquercoisa(x func()) {
	x()
}

func hellofriend() {
	fmt.Println("Hello, friend")
}

func main() {
	qualquercoisa(hellofriend)
}
