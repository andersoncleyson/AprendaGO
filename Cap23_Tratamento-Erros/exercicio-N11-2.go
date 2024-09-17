package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
	Last string
	Sayings []string
}

func main(){
	p1 := person{
		First: "James",
		Last: "bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("é uma cilada bino", err)
	}
	return bs, nil
}