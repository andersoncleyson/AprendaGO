package main

import "fmt"

func main(){

	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t%#x\t%#U\n", i, i, i)
	}
	
}

// - Base-10: decimal, 0-9
// - Base-2: binário, 0-1
// - Base-16: hexadecimal, 0-f
