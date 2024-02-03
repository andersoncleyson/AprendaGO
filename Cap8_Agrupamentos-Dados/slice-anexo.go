package main

import "fmt"

func main(){
	umaslice := []int{1, 2, 3, 4, 5}
	outraslice := []int{6, 7, 8, 9, 10}

	fmt.Println(umaslice)

	umaslice = append(umaslice, outraslice...) //... unfurl

	fmt.Println(umaslice)
}