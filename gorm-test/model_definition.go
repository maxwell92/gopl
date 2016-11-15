package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"` // Default size
	Num      int    `gorm:"AUTO_INCREMENT"`

	CreditCard CreditCard // 1 -> 1
	Emains     []Email    // 1 -> n

	BillingAddress   Address
	BillingAddressId sql.NullInt64

	ShippingAddress   Address
	ShippingAddressId int

	IgnoreMe int        `gorm:"-"`
	Language []Language `gorm:"many2many:user_languages;"` // How many2many ?
}

type Address struct {
	Id       int
	Address1 string         `gorm:"not null;unique"`
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"`
	Code string `gorm:"index:idx_name_code"`
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}
