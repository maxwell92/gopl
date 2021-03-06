package main

import "fmt"

type Car struct {
	Name string
	Type []SubType
}

type SubType struct {
	Name       string
	Id         []string
	Competitor map[string]string
	Owner      *Person
}

type Person struct {
	Name   string
	Gender string
}

func main() {
	BMW := make(map[string]string, 2)
	BMW["X series"] = "X3"
	BMW["5 series"] = "5 GT"

	me := &Person{
		Name:   "liyao",
		Gender: "man",
	}

	benz := new(Car)
	benz.Name = "Benz"
	benz.Type = make([]SubType, 2)
	benz.Type[0].Name = "M series"
	benz.Type[0].Id = make([]string, 2)
	benz.Type[0].Id[0] = "320"
	benz.Type[0].Id[1] = "350"
	benz.Type[0].Competitor = BMW
	benz.Type[0].Owner = me
	benz.Type[1].Name = "G series"
	benz.Type[1].Id = make([]string, 2)
	benz.Type[1].Id[0] = "50"
	benz.Type[1].Id[1] = "55"
	benz.Type[1].Owner = me

	fmt.Println(benz)

}
