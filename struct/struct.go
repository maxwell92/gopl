package main

import "fmt"

type struct2 struct {
	number []int32
}

type struct1 struct {
	member []struct2
}

func main() {
	n := &struct1{
		member: []struct2{
			struct2{
				number: []int32{1, 2, 3},
			},
		},
	}

	fmt.Printf("%v\n", n)
}
