package model

import "www.lffwl.com/library/model"

// 列表输出
type BlogTypeManageIndexOutput struct {
	List  []BlogTypeManageList `json:"list"`  // 列表
	Page  int                  `json:"page"`  // 分页码
	Limit int                  `json:"limit"` // 分页数量
	Total int                  `json:"total"` // 数据总数
}

// 列表List
type BlogTypeManageList struct {
	Id        string `json:"id"`         // ID
	Name      string `json:"name"`       // 名称
	CreatedId string `json:"created_id"` // 添加人
	CreatedAt string `json:"created_at"` // 添加时间
	UpdatedId string `json:"updated_id"` // 最近更新人
	UpdatedAt string `json:"updated_at"` // 最近更新时间
}

// 列表请求参数
type BlogTypeManageIndexReq struct {
	BlogTypeManageIndexInput
	model.PubicGetListPageReq
}
type BlogTypeManageIndexInput struct {
	model.PubicGetListPageInput
}

//新增
type BlogTypeManageStoreReq struct {
	BlogTypeManageStoreInput
	Name       string `v:"required|length:2,16" validate:"required"` // 名称
	SeoKeyword string // seo关键字
	SeoDesc    string // seo描叙
}
type BlogTypeManageStoreInput struct {
	Name       string
	SeoKeyword string // seo关键字
	SeoDesc    string // seo描叙
}

//更新
type BlogTypeManageUpdateReq struct {
	BlogTypeManageUpdateInput
	Id         int    `v:"required|integer" validate:"required"`     // ID
	Name       string `v:"required|length:2,16" validate:"required"` // 名称
	SeoKeyword string // seo关键字
	SeoDesc    string // seo描叙
}
type BlogTypeManageUpdateInput struct {
	Id         int // ID
	Name       string
	SeoKeyword string // seo关键字
	SeoDesc    string // seo描叙
}

//删除
type BlogTypeManageDeleteReq struct {
	Id int `v:"required|integer" validate:"required"` // ID
}

// BlogTypeAllIndex 全部集合
type BlogTypeAllIndex struct {
	Id   string `json:"id"`   // ID
	Name string `json:"name"` // 名称
}
