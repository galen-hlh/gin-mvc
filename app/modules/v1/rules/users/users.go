package users

// 列表
type ListRequest struct {
	StartTime string `form:"start_time" binding:"required"`
	EndTime   string `form:"end_time" binding:"required,date=2006-01-02 15:04:05"`
}

func (f *ListRequest) SetMessage() map[string]string {
	return map[string]string{
		"StartTime.required": "请输入开始时间",
		"EndTime.required":   "请输入结束时间",
		"EndTime.date":       "结束时间格式不正确",
	}
}

// 添加
type AddRequest struct {
	Nickname string `form:"nickname" binding:"required"`
	Phone    string `form:"phone" binding:"required"`
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (f *AddRequest) SetMessage() map[string]string {
	return map[string]string{
		"Account.required":  "请输入账号",
		"Nickname.required": "请输入昵称",
		"Phone.required":    "请输入手机号",
		"Password.required": "请输入密码",
	}
}
