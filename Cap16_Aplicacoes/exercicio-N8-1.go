package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"first"`
	Age   int    `json:"age"`
}

func main() {
	u1 := user{"James", 32}
	u2 := user{"Moneypenny", 27}
	u3 := user{"M", 54}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	people, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", people)

	var titles []struct{ First string }
	if err := json.Unmarshal(people, &titles); err != nil {
		fmt.Println(err)
	}
	fmt.Println(titles)
}
