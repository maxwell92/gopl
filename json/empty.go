package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name   string `json:"name"`
	Bought bool   `json:"bought"`
	Price  int    `json:"price,omitempty"`
}

func main() {
	benz := &Car{
		Name: "Benz",
	}

	bj, err := json.Marshal(benz)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bj))

	car := new(Car)
	err1 := json.Unmarshal(bj, car)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Printf("%v\n", car)
}
