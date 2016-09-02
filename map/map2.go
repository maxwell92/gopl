package main

import (
	"fmt"
)

func main() {
	amap := make(map[string]string, 0)
	amap["xxx"] = "xxx"
	fmt.Println(amap["xxx"])
	//fmt.Println(amap)
	//fmt.Printf("%v", amap)
}
