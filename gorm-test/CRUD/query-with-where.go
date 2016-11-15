package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	// Struct
	user := new(Student)
	//db.Where(&Student{Name: "Jack", Age: 20}).First(&user)
	db.Where(&Student{Name: "Jack", Age: 18}).First(&user)
	fmt.Printf("Query Struct: %s\n", user.Name)

	// Map
	u := new(Student)
	db.Where(map[string]interface{}{"name": "Jack", "age": 18}).Find(&u)
	fmt.Printf("Query Map: %s\n", u.Name)

	// Slice
	// Don't know how it works.
	a := new(Student)
	//db.Where([]int64{18, 19, 20}).Find(&a)
	//db.Where([]int64{18, 19, 20}).Last(&a)
	db.Where([]int64{18, 19, 20}).First(&a)
	fmt.Printf("Query Slice: %s\n", u.Name)

	// Query with inline Condition
	// Way 1:
	b := new(Student)
	db.First(&b, 23)
	fmt.Printf("Query First: %s\n", b.Name)
	// Way 2:
	c := new(Student)
	db.Find(&c, "name = ?", "Jack")
	fmt.Printf("Query Find: %s\n", c.Name)
	// Way 3:
	d := new(Student)
	db.Find(&d, "name <> ? AND age > ?", "Jack", "20")
	fmt.Printf("Query Find: %s\n", d.Name)
	// Way 4:
	e := new(Student)
	db.Find(&e, Student{Age: 20})
	fmt.Printf("Query Find: %s\n", e.Name)
	// Way 5:
	f := new(Student)
	db.Find(&f, map[string]interface{}{"age": 20})
	fmt.Printf("Query Find: %s\n", f.Name)

	// Query with Or
	g := new(Student)
	db.Where("Name = ?", "Jack").Or("Age = 18").Find(&g)
	fmt.Printf("Query Or: %s\n", g.Name)

	// Struct
	h := new(Student)
	db.Where("Name = 'Jack'").Or(Student{Name: "Jack1"}).Find(&h)
	fmt.Printf("Query Where: %s\n", h.Name)

	// Map
	i := new(Student)
	db.Where("Name = 'Jack'").Or(map[string]interface{}{"name": "Jack2"}).Find(&i)
	fmt.Printf("Query Where: %s\n", i.Name)
}
