package view

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gpage"
)

var PublicView = view{}

type view struct {
}

// RenderTpl 渲染模板
func (v *view) RenderTpl(r *ghttp.Request, tpl string, data ...g.Map) {

	var viewData g.Map

	// 检查是否传了值
	if len(data) > 0 {
		viewData = data[0]
	}

	var seo = g.Cfg().GetString("viewer.SeoTitle")

	// seo
	if viewData["Title"] == nil {
		viewData["Title"] = seo[3:]
	} else {
		viewData["Title"] = gconv.String(viewData["Title"]) + seo
	}

	// seo
	if viewData["SeoKey"] == nil {
		viewData["SeoKey"] = seo[3:]
	} else {
		viewData["SeoKey"] = gconv.String(viewData["SeoKey"])
	}

	// seo
	if viewData["SeoDesc"] == nil {
		viewData["SeoDesc"] = seo[3:]
	} else {
		viewData["SeoDesc"] = gconv.String(viewData["SeoDesc"])
	}

	r.Response.WriteTpl(tpl, viewData)
	r.Exit()
}

// RenderError 错误
func (v *view) RenderError(r *ghttp.Request, err string) {
	r.Response.WriteTpl("404.html", g.Map{
		"title": err,
	})
	r.Exit()
}

// WrapContent 分页
func (v *view) WrapContent(page *gpage.Page) string {

	page.SpanStyle = ""
	page.LinkStyle = ""

	content := page.GetContent(3)

	return content
}
