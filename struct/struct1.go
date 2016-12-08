package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name   string `json:"name"`
	Booked bool   `json:"Booked"`
	Sold   bool   `json:"sold"`
}

func main() {
	Benz := &Car{
		Name:   "奔驰",
		Booked: true,
		Sold:   false,
	}

	info, _ := json.Marshal(Benz)
	fmt.Println(string(info))
}
