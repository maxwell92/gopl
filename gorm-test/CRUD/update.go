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

type Users struct {
	Id   int32
	Name string
	Age  int32
}

func main() {
	// Connect to DB
	db, err := gorm.Open("mysql", "root:root@tcp(mysql.yce:3306)/ormtest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Connecting to Database Error")
	}
	defer db.Close()

	// update all fields
	var a Student
	db.First(&a)

	a.Name = "Jen"
	a.Age = 32
	db.Save(&a)
	fmt.Printf("Save: %s\n", a.Name)

	var b Users
	db.First(&b)

	b.Name = "Jen"
	b.Age = 32
	db.Save(&b)
	fmt.Printf("Save: %s\n", b.Name)

	// update changed fields
	var c Student
	// db.Model(&c).Update("name", "Smith")
	db.Model(&c).Where("name = ?", "Jen").Update("name", "Jack")
	fmt.Printf("Update: %s\n", c.Name)

	// update all changed record. if two record change the same column, it fails
	// update multiple attr, only update those changed fields
	var d Users
	db.Model(&d).Update(map[string]interface{}{"name": "Sam"})
	fmt.Printf("Update: %s\n", d.Name)

	// update
	// update multiple attr, only update those changed and non blank fields
	var e Users
	db.Model(&e).Update(Users{Name: "Maxwell"})
	fmt.Printf("Update: %s\n", e.Name)

	// update selected fields
	var f Users
	db.Model(&f).Select("name").Updates(map[string]interface{}{"name": "Amy", "age": 20})
	fmt.Printf("Update: %s, %d\n", f.Name, f.Age)

	// Omit the changes
	var g Users
	db.Model(&g).Omit("name").Updates(map[string]interface{}{"name": "Jack", "age": 18})
	fmt.Printf("Update: %s, %d\n", g.Name, g.Age)

	// update changed fields without callbacks
	var h Users
	db.Model(&h).UpdateColumn("name", "Jay")
	fmt.Printf("UpdateColumn: %s\n", h.Name)
	// db.Model(&h).UpdateColumn(User{Name:"Jeniffer", Age:23})

	// Batch Updates
	db.Table("students").Where("age <> ?", 20).Updates(map[string]interface{}{"age": 100})

	//db.Model(Student{}).Updates(Users{Age: 99}) // should be careful
	db.Model(Student{}).Updates(Student{Age: 99})

	// Get updated records count with RowsAffected
	n := db.Model(Users{}).Updates(Users{Age: 35}).RowsAffected
	fmt.Printf("RowsAffected: %d\n", n)

	// change updaing values in callbacks
}

func (user *Users) BeforeSave(scope *gorm.Scope) (err error) {
	return nil
}

// doesn't work
func (user *Users) BeforeUpdate(scope *gorm.Scope) (err error) {
	scope.SetColumn("Class", 15) // hasn't this column, but it won't be a fatal error
	scope.SetColumn("Age", 15)
	scope.SetColumn("Name", "HotDog")
	return nil
}
