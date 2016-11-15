package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
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

	// Drop Table
	db.DropTable(&Student{})

	// Create Table
	db.CreateTable(&Student{})

	// Create Record
	Jack := &Student{
		Name:     "Jack",
		Age:      18,
		Birthday: time.Now().String(),
	}

	if db.NewRecord(Jack) {
		fmt.Println("New Record \"Jack\" Succeed")
	}

	db.Create(&Jack)

	if !db.NewRecord(Jack) {
		fmt.Println("Of course it cann't be new")
	}

	// Default Values.
	// Must be specified when creating tables.
	var Who = &Student{
		Age:      20,
		Birthday: time.Now().String(),
	}
	db.Create(&Who)
	fmt.Println("Who is created!")

	// Extra Creating Option
	// Didn't work, too
	db.Set("gorm:insert_option", "ON CONFLICT").Create(&Who)
}

// Setting Primary Key In Callbacks
// How to Use ?
func (stu *Student) BeforeCreate(scope *gorm.Scope) error {
	//scope.SetColumn("ID", uuid.New())
	scope.SetColumn("ID", "101")
	return nil
}
