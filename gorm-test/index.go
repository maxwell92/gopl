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

	// Add index for columns `name` with given name `idx_user_name`
	db.Model(&User{}).AddIndex("idx_user_name", "name")
	// Add index for columns `name`, `age` with given name `idx_user_name_age`
	db.Model(&User{}).AddIndex("idx_user_name_age", "name", "age")

	// Cann't update non-unique index to unique index below AddUniqueIndex()
	// Add unique index
	db.Model(&User{}).AddUniqueIndex("idx_user_name", "name")
	db.Model(&User{}).AddUniqueIndex("unique_idx_user_name", "name")

	// Add unique index
	db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "name", "age")
}
