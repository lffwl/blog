package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 中间件管理服务
var Middleware = middlewareService{}

type middlewareService struct{}

// 设置语言 - 中文
func (s *middlewareService) ZHCN(r *ghttp.Request) {
	g.I18n().SetLanguage("zh-CN")
	r.Middleware.Next()
}

// 允许接口跨域请求
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
