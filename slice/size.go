package main

import "fmt"

func main() {
	s := make([]string, 6)
	fmt.Println(len(s))

	s[0] = "0"
	s[1] = "1"
	s[2] = "2"
	s[3] = "3"
	s[4] = "4"
	s[5] = "5"
	/*
		index := len(s)
		half := index / 2

		for i := 0; i < half+1; i++ {
			if (i + half) > len(s) {
				break
			}
			s[i] = s[i+half]
			s[i+half] = ""
		}
		index = half
	*/
	index := len(s)

	var half int

	if judge := index % 2; judge == 0 {
		half = index / 2
	} else {
		half = (index - 1) / 2
	}

	fmt.Printf("%d, %s\n", half, s)
	for i := 0; i < half; i++ {
		s = append(s, "")
	}
	s = s[half:]
	index = len(s) - half

	fmt.Printf("%v, %d, %d\n", s, index, len(s))
}
