package model

import (
	"github.com/gogf/gf/frame/g"
	"www.lffwl.com/library/model"
)

// 列表输出
type BlogManageIndexOutput struct {
	List   []BlogManageList `json:"list"`   // 列表
	Page   int              `json:"page"`   // 分页码
	Limit  int              `json:"limit"`  // 分页数量
	Total  int              `json:"total"`  // 数据总数
	Config g.Map            `json:"config"` // 配置
}

// 列表List
type BlogManageList struct {
	Id        string `json:"id"`         // ID
	TypeId    string `json:"type_id"`    // 博客类型ID
	Title     string `json:"title"`      // 标题
	CreatedId string `json:"created_id"` // 添加人
	CreatedAt string `json:"created_at"` // 添加时间
	UpdatedId string `json:"updated_id"` // 最近更新人
	UpdatedAt string `json:"updated_at"` // 最近更新时间
}

// 列表请求参数
type BlogManageIndexReq struct {
	BlogManageIndexInput
	model.PubicGetListPageReq
}
type BlogManageIndexInput struct {
	model.PubicGetListPageInput
}

// 新增
type BlogManageStoreReq struct {
	BlogManageStoreInput
	TypeId     string `v:"required|min:0" validate:"required"`          // 博客类型ID
	Title      string `v:"required|length:1,999" validate:"required"`   // 标题
	Content    string `v:"required|length:0,21845" validate:"required"` // 内容
	SeoKeyword string // seo关键字
	SeoDesc    string // seo描叙
}
type BlogManageStoreInput struct {
	TypeId     string // 博客类型ID
	Title      string // 标题
	Content    string // 内容
	SeoKeyword string // seo关键字
	SeoDesc    string // seo描叙
}

//更新
type BlogManageUpdateReq struct {
	BlogManageUpdateInput
	Id         int    `v:"required|min:0|integer" validate:"required"`  // ID
	TypeId     string `v:"required|min:0" validate:"required"`          // 博客类型ID
	Title      string `v:"required|length:1,999" validate:"required"`   // 标题
	Content    string `v:"required|length:0,21845" validate:"required"` // 内容
	SeoKeyword string // seo关键字
	SeoDesc    string // seo描叙
}
type BlogManageUpdateInput struct {
	Id         int    // ID
	TypeId     string // 博客类型ID
	Title      string // 标题
	Content    string // 内容
	SeoKeyword string // seo关键字
	SeoDesc    string // seo描叙
}

//删除
type BlogManageDeleteReq struct {
	Id int `v:"required|min:0|integer" validate:"required"` // ID
}
