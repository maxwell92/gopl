package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	// change column description's data type to `text` for model `User`
	//db.Model(&Product{}).ModifyColumn("code", "int")
	//fmt.Println("modified column of  products")

	// drop column code from model `User`
	//db.Model(&User{}).DropColumn("code")

	// Add Foreign Key
	// 1st param: foreign key field
	// 2nd param: destination table(id)
	// 3rd param: ONDELETE
	// 4th param: ONUPDATE

	// hasn't been tested below
	db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")
	fmt.Println("added foreign key")
}
