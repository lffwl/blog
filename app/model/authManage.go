package model

//登录请求
type AuthManageLoginReq struct {
	AuthManageLoginInput
	UserName string `v:"required|length:5,16"` // 用户名
	Password string `v:"required|length:6,16"` // 密码
}
type AuthManageLoginInput struct {
	UserName string
	Password string
}
