package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //easy to remember sql driver package
	// no unused error
)

type Product struct {
	gorm.Model // not neccessary
	Code       string
	Price      uint
}

type User struct {
	Id   string
	Name string
	Age  uint
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(mysql.yce:3306)/ormtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	// Create table way 1:
	db.CreateTable(&Product{})

	// Create table way 2:
	// will append "ENGINE=InnoDB" to the SQL statement when creating table `product`
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
}
