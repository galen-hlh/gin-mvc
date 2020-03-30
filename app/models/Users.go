package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-restful-api/app/compoments/mysql"
)

type Users struct {
	*Model
	Nickname string
	Phone    string
	Account  string
	Password string
}

func (u *Users) TableName() string {
	return "users"
}

func (u *Users) BeforeCreate(scope *gorm.Scope) error {
	u.Model.BeforeCreate(scope)
	//Model.BeforeCreate(scope)
	return nil
}

func (u *Users) Add(nickName string, phone string, account string, password string) {
	db := mysql.GetDb()

	a := db.Debug().Create(&Users{
		Nickname: nickName,
		Phone:    phone,
		Account:  account,
		Password: password,
	})

	fmt.Println(a)
}
