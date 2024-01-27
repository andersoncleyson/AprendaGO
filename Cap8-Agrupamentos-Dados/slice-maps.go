package main

import "fmt"

func main(){
	amigos := map[string]int{
		"Angela": 654654,
		"Shayla": 46545646,
		"Leon": 51321654,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["Shayla"])

	amigos["gopher"] = 54654654

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"], "\n\n")

	sera, ok := amigos["fantasma"]
	fmt.Println(sera, ok)

	// comma ok idiom
	if sera, ok := amigos["asdf"]; !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(sera)
	}
}