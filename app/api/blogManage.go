package api

import (
	"github.com/gogf/gf/net/ghttp"
	"www.lffwl.com/app/model"
	"www.lffwl.com/app/service"
	"www.lffwl.com/library/response"
)

var BlogManage = blogManage{}

type blogManage struct {
}

// @summary 博客列表
// @tags    博客管理
// @produce json
// @param   entity query model.BlogManageIndexReq false "列表请求"
// @router  /manage/blog [Get]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *blogManage) Index(r *ghttp.Request) {

	var req *model.BlogManageIndexReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//获取列表
	if list, err := service.BlogManage.Index(r.Context(), req.BlogManageIndexInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", list)
	}
}

// @summary 博客新增
// @tags    博客管理
// @produce json
// @param   entity query model.BlogManageStoreReq false "新增请求"
// @router  /manage/blog [Post]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *blogManage) Store(r *ghttp.Request) {

	var req *model.BlogManageStoreReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//获取列表
	if err := service.BlogManage.Store(r.Context(), req.BlogManageStoreInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// @summary 博客更新
// @tags    博客管理
// @produce json
// @param   id	path	int    true	"ID"
// @param   entity query model.BlogManageStoreReq false "新增请求"
// @router  /manage/blog/{id} [Put]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *blogManage) Update(r *ghttp.Request) {

	var req *model.BlogManageUpdateReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//获取列表
	if err := service.BlogManage.Update(r.Context(), req.BlogManageUpdateInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// @summary 博客删除
// @tags    博客管理
// @produce json
// @param   id	path	int    true	"ID"
// @router  /manage/blog/{id} [Delete]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *blogManage) Delete(r *ghttp.Request) {

	var req *model.BlogManageDeleteReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//获取列表
	if err := service.BlogManage.Delete(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// @summary 博客详情
// @tags    博客管理
// @produce json
// @param   id	path	int    true	"ID"
// @router  /manage/blog/{id} [Get]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *blogManage) Show(r *ghttp.Request) {

	var req *model.BlogManageDeleteReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//详情
	if info, err := service.BlogManage.Show(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", info)
	}
}
