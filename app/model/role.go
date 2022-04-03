package model

import "www.lffwl.com/library/model"

// 角色列表输出
type RoleGetListOutput struct {
	List  []RoleGetList `json:"list"`  // 列表
	Page  int           `json:"page"`  // 分页码
	Limit int           `json:"limit"` // 分页数量
	Total int           `json:"total"` // 数据总数
}

// 角色列表输出字段
type RoleGetList struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// role列表请求接口
type RoleGetIndexReq struct {
	RoleGetIndexInput
	model.PubicGetListPageReq
}
type RoleGetIndexInput struct {
	model.PubicGetListPageInput
}

//role新增
type RolePostStoreReq struct {
	RolePostStoreInput
	Name string `v:"required|length:2,16" validate:"required"`
}
type RolePostStoreInput struct {
	Name string
}

//role更新
type RolePutUpdateReq struct {
	RolePutUpdateInput
	Id int `v:"required|min:0"`
	RolePostStoreReq
}
type RolePutUpdateInput struct {
	Id int
	RolePostStoreInput
}
type RolePutUpdateReqSwag struct {
	Name string `validate:"required"` //名称
}

//role 删除
type RoleDeleteReq struct {
	Id int `v:"required|min:0" validate:"required"`
}

// 设置api权限
type RolePostSetApiReq struct {
	Id   int    `v:"required|min:0" validate:"required"`                                  // ID
	Apis string `v:"regex:^(\\d+,)*\\d+$|methods|required#只能数字和逗号组合" validate:"required"` // apiId集合
}
type RolePostSetApiData struct {
	RoleId int
	ApiId  int
}

// 角色详情
type RoleShow struct {
	Id   int
	Name string
}

// 角色详情
type RoleGetApiOutput struct {
	Checked []string         `json:"checked"` // 选中
	List    []PermissionList `json:"list"`    // 列表
}
