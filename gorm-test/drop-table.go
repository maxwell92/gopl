package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(mysql.yce:3306)/ormtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	// Drop table way 1:
	db.DropTable(&Product{})

	// Drop table way 2:
	db.DropTable("products")

	// Drop table way 3:
	db.DropTableIfExists(&Product{}, "products")

	fmt.Println("drop table products")
}
