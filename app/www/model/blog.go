package model

import "www.lffwl.com/library/model"

// Blog列表请求接口
type BlogWwwGetIndexReq struct {
	BlogWwwGetIndexInput
	model.PubicGetListPageReq
	TypeId int    `d:"0"  v:"min:0"` // 分类
	Title  string // 标题
	Name   string // 分类名称

}
type BlogWwwGetIndexInput struct {
	model.PubicGetListPageInput
	TypeId int    // 分类
	Title  string // 标题
	Name   string // 分类名称
}

// Blog列表输出
type BlogWwwGetListOutput struct {
	List  []BlogWwwGetList // 列表
	Page  int              // 分页码
	Limit int              // 分页数量
	Total int              // 数据总数
}
type BlogWwwGetList struct {
	Id        string // ID
	TypeId    string // 博客类型ID
	Title     string // 标题
	SeoDesc   string // seo
	UpdatedAt string
	TypeName  string
}

// Blog详情请求接口
type BlogWwwGetShowReq struct {
	BlogWwwGetShowInput
	Name string `v:"required"`       // 分类名称
	Id   int    `v:"required|min:0"` // ID

}
type BlogWwwGetShowInput struct {
	Name string // 分类名称
	Id   int    // ID
}
