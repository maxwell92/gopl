package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	gorm.Model
	Name     string `gorm:"default:'None'"`
	Age      int32
	Birthday string
}

func main() {
	// Connecting to Database
	db, err := gorm.Open("mysql", "root:root@tcp(mysql.yce:3306)/ormtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Connecting Database error")
	}
	defer db.Close()

	// Select
	a := new(Student)
	db.Select("name, age").Find(&a)
	fmt.Printf("select: %s\n", a.Name)

	b := new(Student)
	db.Select([]string{"name", "age"}).Find(&b)
	fmt.Printf("select: %s\n", b.Name)

	/*
		c := db.Table("students").Select("COALESCE(age,?)", 18).Rows()
		fmt.Printf("select: %s\n", c.Name)
	*/

	// Order
	c := new(Student)
	db.Order("age desc, name").Find(&c)
	fmt.Printf("select order: %s\n", c.Name)

	d := new(Student)
	e := new(Student)
	db.Order("age desc").Find(&d).Order("age", true).Find(&e)
	fmt.Printf("select order: %s, %s\n", d.Name, e.Name)

	// Limit
	f := make([]Student, 3)
	db.Limit(3).Find(&f)
	fmt.Printf("select limit: %s, %s, %s\n", f[0].Name, f[1].Name, f[2].Name)

	g := make([]Student, 3)
	h := make([]Student, 3)
	db.Limit(10).Find(&g).Limit(-1).Find(&h)
	fmt.Printf("select limit: %s, %s, %s\n", g[0].Name, g[1].Name, g[2].Name)
	fmt.Printf("select limit: %s, %s, %s\n", h[0].Name, h[1].Name, h[2].Name)

	// Offset
	// SQL error
	i := new(Student)
	db.Offset(2).Find(&i)
	fmt.Printf("select offset: %s\n", i.Name)

	// Count
	j := make([]Student, 3)
	var num1 int
	db.Where("name = ?", "Jack").Or("name = ?", "Justin").Find(&j).Count(&num1)
	fmt.Printf("select count: %d\n", num1)

	var num2 int
	db.Model(&Student{}).Where("name = ?", "Jack").Count(&num2)
	fmt.Printf("select count: %d\n", num2)

	var num3 int
	db.Table("students").Count(&num3)
	fmt.Printf("select count: %d\n", num3)
}
