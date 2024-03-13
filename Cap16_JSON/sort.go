package main

import (
	"fmt"
	"sort"
)

func main(){
	ss := []string{"elliot", "angela", "shayla", "darlene"}

	fmt.Println(ss)

	sort.Strings(ss)

	fmt.Println(ss)
}