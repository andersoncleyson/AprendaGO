package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`

	fmt.Println(s)

	var person []struct {
		First string
		Age   int
	}
	if err := json.Unmarshal([]byte(s), &person); err != nil {
		fmt.Println(err)
	}

	fmt.Println(person[0].First)

}
