package main
import() 

func count() int {
	print("count")	
	return 3
}

func main() {
	c := 0	
	for c < count() {
		println("b", c)
		c++
	}

	for i, c := 0, count(); i < c; i++ {
		println("a", i)
	}
}
