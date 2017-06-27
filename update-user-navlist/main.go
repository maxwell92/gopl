package main

import (
	config "app/backend/common/yce/config"
	myuser "app/backend/model/mysql/user"
	"github.com/maxwell92/gokits/mysql"
	"log"
)

func init() {
	config.Instance().Load(false)
	mysql.MysqlInstance().Open()
}

func main() {
	userList, err := myuser.QueryAllUsers()
	if err != nil {
		log.Printf("QueryAllUsers: error=%s", err)
		panic(err)
	}

	updateNavList(userList)
}

func updateNavList(userList []myuser.User) {
	op := &NewOperationConfig{
		Handle: true,
		Delete: true,
	}

	for _, user := range userList {
		err := user.QueryUserByUserName(user.Name)
		if err != nil {
			log.Printf("QueryUserByUserName Error: error=%s", err)
			panic(err)
		}

		//	log.Printf("updateNavList: userName=%s, navList=%s\n", user.Name, user.NavList)

		navList := DecodeNavList(user.NavList)
		if navList.List[0].Operation != nil {
			oldValue := navList.List[0].Operation.(bool)
			log.Printf("updateNavList: userName=%s, oldValue=%v\n", user.Name, oldValue)
			navList.List[0].Operation = op
			newValue := navList.List[0].Operation.(*NewOperationConfig)
			log.Printf("updateNavList: userName=%s, newValue=%v\n", user.Name, newValue)

			newNavList := navList.EncodeNavList()
			user.NavList = newNavList
			user.InsertUser(1)
		}
	}
}
