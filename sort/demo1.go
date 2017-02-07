package main

import (
	"fmt"
	"sort"
)

type Group struct {
	grades []Grade
	Sorter
}

type Grade struct {
	linux  int
	golang int
	name   string
}

func (g Group) Init() {
}

func (g Group) By() {

}

func (g Group) Sort() {
	fmt.Println("Sort")
}

type Sorter interface {
	Sort()
}

func main() {
	grp := &Group{
		grades: make([]Grade, 5),
	}

	grp.Init()
	grp.Sort()
}
