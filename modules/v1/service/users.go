package service

import (
	"t-gin/models"
	"t-gin/modules/v1/requests/users"
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
