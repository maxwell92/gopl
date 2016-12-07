package main

import "fmt"

type struct2 struct {
	number []int32
}

type struct1 struct {
	member []struct2
}

type struct3 struct {
	member *struct2
}

func main() {
	/*
		n := &struct1{
			member: []struct2{
				struct2{
					number: []int32{1, 2, 3},
				},
			},
		}

		fmt.Printf("%v\n", n)
	*/
	s3 := new(struct3)
	fmt.Println("s3: %v\ts3.member: %v\n", &s3, &s3.member)
	fmt.Println("s3: %v\ts3.member: %v\n", s3, s3.member)
}
