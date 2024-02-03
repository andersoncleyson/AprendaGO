package main

import "fmt"

func main(){
	slice := []string{"banana", "maçã", "jaca", "pêssego"}

	/*for x, y := range slice {
		fmt.Printf("%d %s\n", x, y)

	}*/

	//slice[4] = "melancia"
	slice = append(slice, "melancia")

	for x, y := range slice {
		fmt.Printf("%d %s\n", x, y)

	}

	for _, i := range slice {
		fmt.Printf("%s\n", i)
	}


}