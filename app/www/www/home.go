package www

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"www.lffwl.com/app/www/model"
	"www.lffwl.com/app/www/service"
	"www.lffwl.com/library/view"
)

var Home = home{}

type home struct {
}

// Index 首页
func (a *home) Index(r *ghttp.Request) {

	var (
		data       = g.Map{}
		total      int
		typeTypeKv = g.Map{}
	)

	// 遍历分类
	for _, v := range gconv.Maps(r.Context().Value("type_list")) {
		typeTypeKv[gconv.String(v["id"])] = v["name"]
	}

	var (
		req *model.BlogWwwGetIndexReq
	)

	//参数验证
	if err := r.Parse(&req); err != nil {
		view.PublicView.RenderError(r, err.Error())
	}

	//翻页大小只能是10
	req.PubicGetListPageInput.Limit = 10

	if list, err := service.Blog.List(r.Context(), req.BlogWwwGetIndexInput); err != nil {
		view.PublicView.RenderError(r, err.Error())
	} else {

		for i, v := range list.List {
			v.TypeName = gconv.String(typeTypeKv[v.TypeId])
			list.List[i] = v
		}

		data["list"] = list
		total = list.Total
	}

	data["page"] = view.PublicView.WrapContent(r.GetPage(total, req.BlogWwwGetIndexInput.Limit))

	data["title"] = req.BlogWwwGetIndexInput.Title
	data["type_id"] = req.BlogWwwGetIndexInput.TypeId
	data["name"] = req.BlogWwwGetIndexInput.Name

	view.PublicView.RenderTpl(r, "index.html", data)
}

// Show 详情页面
func (a *home) Show(r *ghttp.Request) {
	var (
		data = g.Map{}
	)

	var (
		req *model.BlogWwwGetShowReq
	)
	//参数验证
	if err := r.Parse(&req); err != nil {
		view.PublicView.RenderError(r, err.Error())
	}

	if info, err := service.Blog.Show(r.Context(), req.BlogWwwGetShowInput.Id); err != nil {
		view.PublicView.RenderError(r, err.Error())
	} else {
		data["info"] = info
		data["Title"] = info["title"]
		data["SeoKey"] = info["seo_keyword"]
		data["SeoDesc"] = info["seo_desc"]
	}

	data["name"] = req.BlogWwwGetShowInput.Name

	view.PublicView.RenderTpl(r, "show.html", data)
}

// About 关于我们
func (a *home) About(r *ghttp.Request) {

	var (
		data = g.Map{}
	)

	data["name"] = "about"

	view.PublicView.RenderTpl(r, "about.html", data)
}
