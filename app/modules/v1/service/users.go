package service

import (
	"go-restful-api/app/models"
	"go-restful-api/app/modules/v1/rules/users"
)

type UsersService struct {
}

func (u *UsersService) Add(request users.AddRequest) {
	orm := models.Users{}

	orm.Add(request.Nickname,
		request.Phone,
		request.Account,
		request.Password)
}
