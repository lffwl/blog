// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// WwwApiDao is the manager for logic model data accessing and custom defined data operations functions management.
type WwwApiDao struct {
	gmvc.M               // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      wwwApiColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB        // DB is the raw underlying database management object.
	Table  string        // Table is the underlying table name of the DAO.
}

// WwwApiColumns defines and stores column names for table www_api.
type wwwApiColumns struct {
	Id        string // ID
	Name      string // 名称
	Key       string // 权限标识
	Router    string // 请求地址
	Methods   string // 请求类型
	Pid       string // 上级ID
	CreatedId string // 添加人
	CreatedAt string // 添加时间
	UpdatedId string // 最近更新人
	UpdatedAt string // 最近更新时间
	DeletedId string // 软删除人
	DeletedAt string // 软删除时间
}

// NewWwwApiDao creates and returns a new DAO object for table data access.
func NewWwwApiDao() *WwwApiDao {
	columns := wwwApiColumns{
		Id:        "id",
		Name:      "name",
		Key:       "key",
		Router:    "router",
		Methods:   "methods",
		Pid:       "pid",
		CreatedId: "created_id",
		CreatedAt: "created_at",
		UpdatedId: "updated_id",
		UpdatedAt: "updated_at",
		DeletedId: "deleted_id",
		DeletedAt: "deleted_at",
	}
	return &WwwApiDao{
		C:     columns,
		M:     g.DB("default").Model("www_api").Safe(),
		DB:    g.DB("default"),
		Table: "www_api",
	}
}