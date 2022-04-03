package model

//详情返回字段
type CenterShow struct {
	Id       string           `json:"id"`
	UserName string           `json:"user_name"`
	Auth     []PermissionList `json:"auth"`
}
type PermissionList struct {
	Id      int    `json:"id"`
	Pid     int    `json:"pid"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Menu    string `json:"menu"`
	Router  string `json:"router"`
	Methods string `json:"methods"`
}
