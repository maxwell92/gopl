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

type Member struct {
	gorm.Model
	Name string
	Age  int32
}

func main() {
	// Connecting to Database
	db, err := gorm.Open("mysql", "root:root@tcp(mysql.yce:3306)/ormtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Connecting Database error")
	}
	defer db.Close()

	// Select
	a := new(Student)
	db.Select("name, age").Find(&a)
	fmt.Printf("select: %s\n", a.Name)

	b := new(Student)
	db.Select([]string{"name", "age"}).Find(&b)
	fmt.Printf("select: %s\n", b.Name)

	/*
		c := db.Table("students").Select("COALESCE(age,?)", 18).Rows()
		fmt.Printf("select: %s\n", c.Name)
	*/

	// Order
	c := new(Student)
	db.Order("age desc, name").Find(&c)
	fmt.Printf("select order: %s\n", c.Name)

	d := new(Student)
	e := new(Student)
	db.Order("age desc").Find(&d).Order("age", true).Find(&e)
	fmt.Printf("select order: %s, %s\n", d.Name, e.Name)

	// Limit
	f := make([]Student, 3)
	db.Limit(3).Find(&f)
	fmt.Printf("select limit: %s, %s, %s\n", f[0].Name, f[1].Name, f[2].Name)

	g := make([]Student, 3)
	h := make([]Student, 3)
	db.Limit(10).Find(&g).Limit(-1).Find(&h)
	fmt.Printf("select limit: %s, %s, %s\n", g[0].Name, g[1].Name, g[2].Name)
	fmt.Printf("select limit: %s, %s, %s\n", h[0].Name, h[1].Name, h[2].Name)

	// Offset
	// SQL error
	i := new(Student)
	db.Offset(2).Find(&i)
	fmt.Printf("select offset: %s\n", i.Name)

	// Count
	j := make([]Student, 3)
	var num1 int
	db.Where("name = ?", "Jack").Or("name = ?", "Justin").Find(&j).Count(&num1)
	fmt.Printf("select count: %d\n", num1)

	var num2 int
	db.Model(&Student{}).Where("name = ?", "Jack").Count(&num2)
	fmt.Printf("select count: %d\n", num2)

	var num3 int
	db.Table("students").Count(&num3)
	fmt.Printf("select count: %d\n", num3)

	// Group & Having
	// Don't know how does it works.
	//rows, err := db.Table("students").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
	//	for rows.Next() {
	//		l := new(Student)
	//		rows.Scan(&l)
	//		fmt.Printf("select group next: %s", l.Name)
	//	}

	//rows1, err := db.Table("students").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()

	//	k := make(Student)
	//	db.Table("students").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount)>?", 100).Scan(&k)
	//	fmt.Printf("select group having scan: %s\n", k.Name)

	// Join
	// rows, err := db.Table("students").Select("students.name, schools.name").Joins("left join schools on schools.student_id = students.id").Rows()

	// db.Table("students").Select("students.name, schools.name").Join("left join schools on schools.student_id = students.id").Scan(&results)
	// db.Joins("JOIN schools ON schools.student_id = students.id AND users.name = ?", "Jack").Joins("JOIN grades ON grades.students_id = students.id").Where("grades.students_id = ?", "04091079").Find(&Jack)

	// Pluck
	var ages []int32
	db.Find(&Student{}).Pluck("age", &ages)
	fmt.Printf("select pluck: %v\n", ages)

	var names []string
	// get one result
	db.Find(&Student{}).Pluck("name", &names)
	fmt.Printf("select pluck: %v\n", names)
	// only this will be just a model, names = []
	db.Model(&Student{}).Pluck("name", &names)
	fmt.Printf("select pluck: %v\n", names)
	// only this will get data, names = ["Jack", "Jimmy", "Justin", "None"]
	db.Table("users").Pluck("name", &names)
	fmt.Printf("select pluck: %v\n", names)

	// Scan
	var m Member
	db.Table("students").Select("name, age").Where("name = ?", "Jack").Scan(&m)
	fmt.Printf("scan: %s\t%d\n", m.Name, m.Age)

	var n Member
	db.Raw("SELECT name, age FROM students WHERE name=?", "Jack").Scan(&n)
	fmt.Printf("scan: %s\t%d\n", n.Name, n.Age)

	var o Student

	// Scopes
	db.Scopes(AgeGreaterThan10, NameIsNotJack).Find(&o)
	fmt.Printf("Scopes: %s\n", o.Name)

	//  var p Student
	//	db.Scopes(NameStatus("Jack")).Find(&p)
	//	fmt.Printf("Scopes: %s\n", p.Name)

	// Specifying The Table Name
	//	db.Table("users").CreateTable(&Member{})
	var q Student
	db.Table("students").Where("name = ?", "Jack").Find(&q)
	fmt.Printf("Scopes: %s, %d\n", q.Name, q.Age)
}

func AgeGreaterThan10(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 10)
}

func NameIsNotJack(db *gorm.DB) *gorm.DB {
	return db.Where("name <> ?", "Jack")
}

func NameStatus(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(AgeGreaterThan10).Where("Name = ?", name)
	}
}
