package exceptions

import "go-restful-api/app/compoments/errors"

func NewUser(code int32) errors.BusinessLib {
	return &errors.Business{
		Num: code,
		Msg: userText[code],
	}
}

//用户枚举值 10000-10100
const (
	// 用户不存在
	UserNotExist = 10000
	// 用户帐号错误
	UserAccountError = 10001
	// 用户密码错误
	UserPasswordError = 10002
)

var userText = map[int32]string{
	UserNotExist:      "用户不存在",
	UserAccountError:  "用户帐号错误",
	UserPasswordError: "用户密码错误",
}
