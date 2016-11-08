package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //easy to remember sql driver package
	// no unused error
)

type Car struct {
	gorm.Model
	Brand string
	Price uint
}

type AirCraft struct { // change to air_crafts / mush_air_crafts
	gorm.Model // not neccessary
	Engine     string
	Price      uint
}

func (Car) TableName() string {
	return "mycar"
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(mysql.yce:3306)/ormtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "mush_" + defaultTableName
	}

	db.CreateTable(&AirCraft{})
}
