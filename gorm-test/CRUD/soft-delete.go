package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Car struct {
	gorm.Model
	Name string
}

func main() {
	// Connect to DB
	db, err := gorm.Open("mysql", "root:root@tcp(mysql.yce:3306)/ormtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Connecting to Database Error")
	}
	defer db.Close()
	a := &Car{
		Name: "RangeRover",
	}

	db.Create(&a)
	fmt.Println(db.NewRecord(a))
	/*
		db.Delete(&a)
		var b Car
		db.Where("name = ?", "RangeRover").Find(&b)
		fmt.Printf("%s\n", b.Name)
	*/
	var e Car
	db.Where("name = ?", "RangeRover").Find(&e)
	//fmt.Printf("%s\n", e.deleted_at) // if deleted_at is not member in Struct Car, it won't be printed.
}
