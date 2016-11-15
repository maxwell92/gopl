package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Factory struct {
	gorm.Model
	Car   Car `gorm:"ForeignKey:CarID"`
	CarID int
}

// `gorm:"ForeignKey:CarID;AssociationForeignKey:ID"`
// How is it be done if the Id is in gorm.Model ?

type Car struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(mysql.yce:3306)/ormtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	/*
		It didn't work below.

		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Car{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Factory{})
		car := new(Car)
		factory := new(Factory)
		db.Model(&car).Related(&factory)
		fmt.Println("Related finished!")
	*/

	/*
		var factory Factory
		var car Car
		db.Model(&car).Related(&factory, "Factory")
		db.Model(&car).Related(&factory)
		fmt.Println("Related Finished!")
	*/

	/*
	 * Association Mode
	 */
	var user user

	db.Model = (&user).Association("Languages")
	// Query
	db.Model(&user).Association("Language").Find(&languages)
	// Append
	db.Model(&user).Association("Languages").Append(Language{Name: "DE"})
	// Delete
	db.Model(&user).Association("Languages").Delete(languageZH, languageEN)
	// Replace
	db.Model(&user).Association("Languages").Replace(Language{Name: "DE"}, languageEN)
	// Count
	db.Model(&user).Association("Languages").Count()
	// Clear
	db.Model(&user).Association("Languages").Clear()
}
