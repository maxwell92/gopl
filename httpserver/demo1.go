package main

import (
	"io"
	"net/http"
)

type a struct{}

func (*a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world version 1.")
}

func main() {
	http.ListenAndServe(":8080", &a{}) //第2个参数需要实现Hander的struct，a满足
}
