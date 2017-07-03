package main

import (
	"fmt"
)

func helperContain(list []int32, id int32) bool {
	for _, elem := range list {
		if elem == id {
			return true
		}
	}

	return false
}

func helperDedup(list []int32) []int32 {
	result := make([]int32, 0)
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {

		}
	}
	return result
}

func helperUnion(list1, list2 []int32) []int32 {
	all := make([]int32, 0)
	all = append(all, list1...)
	for _, elem := range list2 {
		if helperContain(all, elem) {
			continue
		}
		all = append(all, elem)
	}
	return all
}

func helperIntersection(list1, list2 []int32) []int32 {
	all := helperUnion(list1, list2)

	intersection := make([]int32, 0)
	for _, elem := range all {
		if helperContain(list1, elem) && helperContain(list2, elem) {
			intersection = append(intersection, elem)
		}
	}
	return intersection
}

func main() {
	list1 := []int32{1, 2, 3}
	list2 := []int32{2, 3, 4, 5}

	fmt.Println(list1)
	fmt.Println(list2)

	fmt.Println(helperUnion(list1, list2))
	fmt.Println(helperIntersection(list1, list2))
}
