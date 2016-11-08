/* Hasn't succeed
 *
 */
package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(100)"`
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	// register model
	orm.RegisterModel(new(User))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@/my_db?charset=utf8", 30)
}

func main() {
	init()

	o := orm.NewOrm()

	user := User{Name: "mushroom"}

	// insert
	id, err := o.Insert(&user)
	fmt.Println(id)
	checkError(err)

	// update
	user.Name = "mushroom"
	num, err := o.Update(&user)
	fmt.Println(num)
	checkError(err)

	// read one
	u := User{Id: user.Id}
	err = o.Read(&u)
	checkError(err)

	// delete
	num, err = o.Delete(&u)
	checkError(err)
}
