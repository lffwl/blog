package model

import (
	"github.com/gogf/gf/frame/g"
	"www.lffwl.com/library/model"
)

// api列表输出
type ApiGetIndexOutput struct {
	List   []ApiGetIndexList `json:"list"`   // 列表
	Total  int               `json:"total"`  // 数据总数
	Config g.Map             `json:"config"` // 配置
}
type ApiGetIndexList struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	Menu      string `json:"menu"`
	Router    string `json:"router"`
	Methods   string `json:"methods"`
	Pid       int    `json:"pid"`
	Type      int    `json:"type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// api列表请求接口
type ApiGetIndexReq struct {
	ApiGetIndexInput
	model.PubicGetListPageReq
}

// api列表请求参数
type ApiGetIndexInput struct {
	model.PubicGetListPageInput
	ids []string
}

//新增参数
type ApiPostStoreReq struct {
	ApiPostStoreInput
	Pid     int    `v:"integer|min:0" `                              // 上级ID
	Router  string `v:"max-length:100" `                             // 请求地址
	Name    string `v:"required|max-length:60" validate:"required" ` // 名称
	Key     string `v:"required|max-length:60" validate:"required" ` // 权限标识
	Menu    string `v:"max-length:30" `                              // 菜单地址
	Methods string `v:"in-array:www,methods" `                       // 请求类型
	Type    int    `v:"integer|min:0" `                              // 菜单类型
}
type ApiPostStoreInput struct {
	Name    string
	Key     string
	Menu    string
	Router  string
	Methods string
	Pid     int
	Type    int
}

//更新参数
type ApiPutUpdateReq struct {
	ApiPutUpdateInput
	ApiPostStoreReq
	Id int `v:"required|max-length:100"` //ID
}
type ApiPutUpdateInput struct {
	ApiPostStoreInput
	Id int
}
type ApiPutUpdateReqSwag struct {
	Pid     int    `validate:"required"`  // 上级ID
	Router  string `validate:"required"`  // 请求地址
	Methods string ` validate:"required"` // 请求类型
}

//删除参数
type ApiDeleteReq struct {
	Id int `v:"required|max-length:100"` //ID
}
