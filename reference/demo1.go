package main

import()

func mkSlice() []int {
	s := make([]int, 0, 5)
	s = append(s, 100)
	return s
}

func mkMap() map[int]string {
	s := make(map[int]string)
	s[1] = "Jack"
	return s
}

func main() {
	s := mkSlice()
	println(s[0])

	m := mkMap()
	println(m[1])
}
