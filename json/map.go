package main

import (
	"encoding/json"
	"fmt"
)

/*
type Dict struct {
	Map map[string]string `json:"map"`
}
*/

/*
func main() {
	d := new(Dict)
	d.Map = make(map[string]string)

	d.Map["hello"] = "hi"

	data, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	dd := new(Dict)
	err = json.Unmarshal([]byte(data), dd)
	if err != nil {
		panic(err)
	}

	fmt.Println(dd.Map)
}
*/

type Dic map[string]string

func main() {
	// d := new(Dic)
	d := make(map[string]string)
	d["hello"] = "hi"

	data, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	dd := new(Dic)
	// dd := make(map[string]string)
	err = json.Unmarshal([]byte(data), dd)
	if err != nil {
		panic(err)
	}

	fmt.Println(dd)
}
