package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	gorm.Model
	Name      string `gorm:"default:'None'"`
	Age       int32
	Birthday  string
	DeletedAt string
}

type User struct {
	Id   int32
	Name string `gorm:"default:'None'"`
	Age  int32
}

func main() {
	// Connect to DB
	db, err := gorm.Open("mysql", "root:root@tcp(mysql.yce:3306)/ormtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Connecting to Database Error")
	}
	defer db.Close()

	// Delete
	a := new(User)
	db.Delete(&a)
	fmt.Printf("Delete: %s\n", a.Name)

	// Batch Delete
	db.Where("name LIKE ?", "%Jen%").Delete(Student{})
	fmt.Printf("Delete Over")

	db.Delete(User{}, "name LIKE ?", "%None%")
	fmt.Printf("Delete Over")

	// Soft Delete
	b := new(Student)
	db.Delete(&b)
	db.Where("id = 1").Find(&b)
	fmt.Printf("Soft Delete: %s\n", b.Name)

	Jeff := &User{
		Id:   10010,
		Name: "Jeff",
		Age:  18,
	}
	db.Create(&Jeff)
	fmt.Printf("Soft Create: %s\n", Jeff.Name)

	d := new(User)
	db.Where("id = 10010").Find(&d)
	fmt.Printf("Soft Find: %s\n", d.Name)

	fmt.Println(db.NewRecord(d))

}
