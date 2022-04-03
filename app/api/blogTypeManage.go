package api

import (
	"github.com/gogf/gf/net/ghttp"
	"www.lffwl.com/app/model"
	"www.lffwl.com/app/service"
	"www.lffwl.com/library/response"
)

var BlogTypeManage = blogTypeManage{}

type blogTypeManage struct {
}

// @summary 博客类型列表
// @tags    博客类型管理
// @produce json
// @param   entity query model.BlogTypeManageIndexReq false "列表请求"
// @router  /manage/blog-type [Get]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *blogTypeManage) Index(r *ghttp.Request) {

	var req *model.BlogTypeManageIndexReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//获取列表
	if list, err := service.BlogTypeManage.Index(r.Context(), req.BlogTypeManageIndexInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", list)
	}
}

// @summary 博客类型新增
// @tags    博客类型管理
// @produce json
// @param   entity query model.BlogTypeManageStoreReq false "新增请求"
// @router  /manage/blog-type [Post]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *blogTypeManage) Store(r *ghttp.Request) {

	var req *model.BlogTypeManageStoreReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.BlogTypeManage.Store(r.Context(), req.BlogTypeManageStoreInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")

}

// @summary 博客类型更新
// @tags    博客类型管理
// @produce json
// @param   id	path	int    true	"ID"
// @param   entity query model.BlogTypeManageUpdateReq false "新增请求"
// @router  /manage/blog-type/{id} [Put]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *blogTypeManage) Update(r *ghttp.Request) {

	var req *model.BlogTypeManageUpdateReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.BlogTypeManage.Update(r.Context(), req.BlogTypeManageUpdateInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")

}

// @summary 博客类型删除
// @tags    博客类型管理
// @produce json
// @param   id	path	int    true	"ID"
// @router  /manage/blog-type/{id} [Delete]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *blogTypeManage) Delete(r *ghttp.Request) {

	var req *model.BlogTypeManageDeleteReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.BlogTypeManage.Delete(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")

}
