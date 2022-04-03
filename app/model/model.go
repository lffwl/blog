// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model

import (
	"github.com/gogf/gf/os/gtime"
)

// WwwAdmin is the golang structure for table www_admin.
type WwwAdmin struct {
	Id           int         `orm:"id,primary"       json:"id"`           //
	UserName     string      `orm:"user_name,unique" json:"userName"`     // 用户名
	Password     string      `orm:"password"         json:"password"`     // 明文密码
	HashPassword string      `orm:"hash_password"    json:"hashPassword"` // hash密码
	CreatedId    int         `orm:"created_id"       json:"createdId"`    // 添加人
	CreatedAt    *gtime.Time `orm:"created_at"       json:"createdAt"`    // 添加时间
	UpdatedId    int         `orm:"updated_id"       json:"updatedId"`    // 最近更新人
	UpdatedAt    *gtime.Time `orm:"updated_at"       json:"updatedAt"`    // 最近更新时间
	DeletedId    int         `orm:"deleted_id"       json:"deletedId"`    // 软删除人
	DeletedAt    *gtime.Time `orm:"deleted_at"       json:"deletedAt"`    // 软删除时间
}

// WwwAdminRole is the golang structure for table www_admin_role.
type WwwAdminRole struct {
	AdminId int `orm:"admin_id,primary" json:"adminId"` //
	RoleId  int `orm:"role_id,primary"  json:"roleId"`  //
}

// WwwApi is the golang structure for table www_api.
type WwwApi struct {
	Id        int         `orm:"id,primary" json:"id"`        // ID
	Name      string      `orm:"name"       json:"name"`      // 名称
	Key       string      `orm:"key"        json:"key"`       // 权限标识
	Router    string      `orm:"router"     json:"router"`    // 请求地址
	Methods   string      `orm:"methods"    json:"methods"`   // 请求类型
	Pid       int         `orm:"pid"        json:"pid"`       // 上级ID
	CreatedId int         `orm:"created_id" json:"createdId"` // 添加人
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` // 添加时间
	UpdatedId int         `orm:"updated_id" json:"updatedId"` // 最近更新人
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` // 最近更新时间
	DeletedId int         `orm:"deleted_id" json:"deletedId"` // 软删除人
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deletedAt"` // 软删除时间
}

// WwwBlog is the golang structure for table www_blog.
type WwwBlog struct {
	Id         int         `orm:"id,primary"  json:"id"`         // ID
	TypeId     int         `orm:"type_id"     json:"typeId"`     // 博客类型ID
	Title      string      `orm:"title"       json:"title"`      // 标题
	Content    string      `orm:"content"     json:"content"`    // 内容
	SeoKeyword string      `orm:"seo_keyword" json:"seoKeyword"` // seo关键字
	SeoDesc    string      `orm:"seo_desc"    json:"seoDesc"`    // seo描叙
	CreatedId  int         `orm:"created_id"  json:"createdId"`  // 添加人
	CreatedAt  *gtime.Time `orm:"created_at"  json:"createdAt"`  // 添加时间
	UpdatedId  int         `orm:"updated_id"  json:"updatedId"`  // 最近更新人
	UpdatedAt  *gtime.Time `orm:"updated_at"  json:"updatedAt"`  // 最近更新时间
	DeletedId  int         `orm:"deleted_id"  json:"deletedId"`  // 软删除人
	DeletedAt  *gtime.Time `orm:"deleted_at"  json:"deletedAt"`  // 软删除时间
}

// WwwBlogType is the golang structure for table www_blog_type.
type WwwBlogType struct {
	Id         int         `orm:"id,primary"  json:"id"`         // ID
	Name       string      `orm:"name"        json:"name"`       // 名称
	SeoKeyword string      `orm:"seo_keyword" json:"seoKeyword"` // seo关键字
	SeoDesc    string      `orm:"seo_desc"    json:"seoDesc"`    // seo描叙
	CreatedId  int         `orm:"created_id"  json:"createdId"`  // 添加人
	CreatedAt  *gtime.Time `orm:"created_at"  json:"createdAt"`  // 添加时间
	UpdatedId  int         `orm:"updated_id"  json:"updatedId"`  // 最近更新人
	UpdatedAt  *gtime.Time `orm:"updated_at"  json:"updatedAt"`  // 最近更新时间
	DeletedId  int         `orm:"deleted_id"  json:"deletedId"`  // 软删除人
	DeletedAt  *gtime.Time `orm:"deleted_at"  json:"deletedAt"`  // 软删除时间
}

// WwwRole is the golang structure for table www_role.
type WwwRole struct {
	Id        int         `orm:"id,primary" json:"id"`        // ID
	Name      string      `orm:"name"       json:"name"`      // 名称
	CreatedId int         `orm:"created_id" json:"createdId"` // 添加人
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` // 添加时间
	UpdatedId int         `orm:"updated_id" json:"updatedId"` // 最近更新人
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` // 最近更新时间
	DeletedId int         `orm:"deleted_id" json:"deletedId"` // 软删除人
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deletedAt"` // 软删除时间
}

// WwwRoleApi is the golang structure for table www_role_api.
type WwwRoleApi struct {
	ApiId  int `orm:"api_id,primary"  json:"apiId"`  //
	RoleId int `orm:"role_id,primary" json:"roleId"` //
}