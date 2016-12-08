package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name   string `json:"name"`
	Booked bool   `json:"Booked,omitempty"`
	Sold   bool   `json:"sold,omitempty"`
}

func main() {
	Benz := &Car{
		Name:   "奔驰",
		Booked: true,
		Sold:   false, //添加了omitempty字段后，没值或者赋为该字段0值都不会对它进行json编码
	}

	info, _ := json.Marshal(Benz)
	fmt.Println(string(info))
}
