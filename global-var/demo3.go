package main

import "fmt"

//./demo3.go:5: syntax error: non-declaration statement outside function body
//str := "hello world"

//./demo3.go:9: syntax error: non-declaration statement outside function body
//var str string
//str = "hello world"

//./demo3.go:14: syntax error: non-declaration statement outside function body
var str = make(map[string]string)
str["hello"] = "world"

func main() {
	fmt.Println(str)
}
