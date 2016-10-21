package main

import (
	"encoding/json"
	"fmt"
)

type Dictionary struct {
	Dic map[int32]string `json:"dictionary"`
}

func main() {
	d := new(Dictionary)
	d.Dic = make(map[int32]string)

	d.Dic[1] = "abc"
	dj, err := json.Marshal(d.Dic)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dj))
}

// go 1.6 doesn't support map[int]string convert to json, for json didn't support int but astring as key.
// go 1.7 will convert key int automatically before marshal.
