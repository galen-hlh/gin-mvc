package service

import (
	"gin-mvc/models"
	"gin-mvc/modules/v1/requests/users"
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
