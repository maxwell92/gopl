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
	// Connecting to Database
	db, err := gorm.Open("mysql", "root:root@tcp(mysql.yce:3306)/ormtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Connecting Database error")
	}
	defer db.Close()

	// Query
	stu := new(Student)
	db.First(&stu)
	fmt.Printf("First: %s\n", stu.Name)

	someone := new(Student)
	db.Last(&someone)
	fmt.Printf("Last: %s\n", someone.Name)

	find := new(Student)
	db.Find(&find)
	fmt.Printf("Find: %s\n", find.Name)
	fmt.Printf("Find: %s\n", find)

	ten := new(Student)
	db.First(&ten, 10)
	fmt.Printf("First: %s\n", ten.Name)

	// Plain SQL
	u := new(Student)
	// get first match
	db.Where("name = ?", "Jack").First(&u)
	fmt.Printf("Where First: %s\n", u.Name)

	db.Where("name = ?", "Jack").Find(&u)
	fmt.Printf("Where Find: %s\n", u.Name)

	m := new(Student)
	db.Where("name <> ?", "Jack").Find(&m)
	fmt.Printf("Where <>: %s\n", m.Name)

	// IN
	j := new(Student)
	db.Where("name in (?)", []string{"Jack"}).Find(&j)
	fmt.Printf("Where in: %s\n", j.Name)

	// LIKE
	l := new(Student)
	db.Where("name LIKE ?", "%ack%").Find(&l)
	fmt.Printf("Where Like: %s\n", l.Name)

	// AND
	a := new(Student)
	db.Where("name LIKE ? AND age >= ?", "%ack%", "15").Find(&a)
	fmt.Printf("Where And: %s\n", a.Name)

	// Time
	today := time.Now().String
	t := new(Student)
	/*
	 db.Where("name LIKE ?", "%ack%").First(&t)
	 fmt.Printf("Where time: %s\n", t.created_at)
	*/

	/*
	 db.Where("created_at > ?", today).First(&t)
	 fmt.Printf("Where Time: %s\n", t.Name)
	*/
}
