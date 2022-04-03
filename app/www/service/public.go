package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Public = public{}

type public struct {
}

// WebBase 控制器中间件
func (m *public) WebBase(r *ghttp.Request) {

	// 分类列表
	typeList, _ := BlogType.List(r.Context())

	// 设置上下文对象 - 分类列表
	r.SetCtxVar("type_list", typeList)

	r.Assigns(g.Map{
		"type_list": typeList,
	})

	r.Middleware.Next()

}
