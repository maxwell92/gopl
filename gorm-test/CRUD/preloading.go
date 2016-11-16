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
	// Connect to DB
	db, err := gorm.Open("mysql", "root:root@tcp(mysql.yce:3306)/ormtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Connecting to Database Error")
	}
	defer db.Close()

	// Preloading
	var a Student
	db.Preload("Student").Find(&a) // Preload(struct)
	fmt.Printf("preloading: %s\n", a.Name)

	var b Student
	db.Where("name = ?", "Jack").Preload("Student", "Age = ?", 20).Find(&b)
	fmt.Printf("preloading: %s\n", b.Name)

	// custom preloading mysql
	c := make([]Student, 4)
	db.Preload("Student", func(db *gorm.DB) *gorm.DB {
		return db.Order("Student.age DESC")
	}).Find(&c)
	fmt.Printf("preloading: %v\n", c)
}
