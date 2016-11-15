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

	// Query Chains
	// db.Where("name <> ?", "Jack").Where("age >= ? and Birthday <> ?", 20, "2016-01-01").Find(&student)
	// db.Where("Name = ?", "Jack").Or("Age >= ?", 20).Not("Birthday <> ?", "2016-01-01").Find(&student)

	// Extra Querying Option
	// db.Set("gorm:query_option", "FOR UPDATE").First(&student, 20)

	// FirstOrInit: Get first matched record, or initialized a new one with given conditions.
	a := new(Student)
	db.FirstOrInit(&a, Student{Name: "Justin"})

	db.Where(Student{Name: "Justin"}).FirstOrInit(&a)
	// db.FirstOrInit(&a, map[string]interface{}{"Name": "Justin"})
	fmt.Printf("FirstOrInit: %s\n", a.Name)

	db.Create(&a)

	// failed
	//fmt.Println(db.NewRecord(&a))

	// Attrs Age 18 doesn't affect b.Age
	b := new(Student)
	//db.Where(Student{Name: "None"}).Attrs(Student{Age: 18}).FirstOrInit(&b)
	//db.Where(Student{Name: "None_existed"}).Attrs(Student{Age: 19}).FirstOrInit(&b)
	//db.Where(Student{Name: "non_existing"}).Attrs("Age":20).FirstOrInit(&b)
	db.Where(Student{Name: "Jack"}).Attrs("Age", 18).FirstOrInit(&b)
	fmt.Printf("FirstOrInit: %s, %d\n", b.Name, b.Age)

	c := new(Student)
	db.Where(Student{Name: "Jack"}).Assign(Student{Age: 18}).FirstOrInit(&c)
	fmt.Printf("FirstOrInit: %s, %d\n", c.Name, c.Age)

	d := new(Student)
	db.FirstOrCreate(&d, Student{Name: "Jimmy"})
	fmt.Printf("FirstOrCreate: %s\n", d.Name)

	e := new(Student)
	e.Name = "Jeff"
	db.Where(Student{Name: "Jack"}).FirstOrCreate(&e)
	fmt.Printf("FirstOrCreate: %s\n", e.Name)
}
