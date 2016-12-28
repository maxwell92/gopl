package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type Student struct {
	Name string `json:"name"`
}

type Someone struct {
	Name string
	*iris.Context
}

func (s Someone) Patch() {
	r := new(Student)
	err := s.ReadJSON(&r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", r)
	fmt.Println("This is Patch")
	s.Write("Patch\n")
}

func (s Someone) Post() {
	r := new(Student)
	err := s.ReadJSON(&r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", r)
	fmt.Println("This is Post")
	s.Write("Post\n")
}

func (s Someone) Get() {
	r := new(Student)
	err := s.ReadJSON(&r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", r)
	fmt.Println("This is Get")
	s.Write("Get\n")
}

func (s Someone) Put() {
	r := new(Student)
	err := s.ReadJSON(&r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", r)
	fmt.Println("This is Put")
	s.Write("Put\n")
}

func (s Someone) Delete() {
	r := new(Student)
	err := s.ReadJSON(&r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", r)
	fmt.Println("This is Delete")
	s.Write("Delete\n")
}

func main() {
	s := &Someone{
		Name: "Jay",
	}
	iris.API("/hello", *s)
	iris.Listen(":8081")
}
