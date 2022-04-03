package api

import (
	"github.com/gogf/gf/net/ghttp"
	"www.lffwl.com/app/model"
	"www.lffwl.com/app/service"
	"www.lffwl.com/library/response"
)

var Api = apiManage{}

type apiManage struct {
}

// @summary 接口列表
// @tags    接口管理
// @produce json
// @param   entity query model.ApiGetIndexReq false "列表请求"
// @router  /manage/www [Get]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiManage) Index(r *ghttp.Request) {
	var (
		req *model.ApiGetIndexReq
	)

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//获取列表
	if list, err := service.Api.Index(r.Context(), req.ApiGetIndexInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", list)
	}
}

// @summary 接口新增
// @tags    接口管理
// @produce json
// @param   entity query model.ApiPostStoreReq false "新增请求"
// @router  /manage/www [Post]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiManage) Store(r *ghttp.Request) {
	var req *model.ApiPostStoreReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.Api.Store(r.Context(), req.ApiPostStoreInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// @summary 接口更新
// @tags    接口管理
// @produce json
// @param   id	path	int    true	"ID"
// @param   entity query model.ApiPutUpdateReqSwag false "更新请求"
// @router  /manage/www/{id} [Put]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiManage) Update(r *ghttp.Request) {
	var req *model.ApiPutUpdateReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.Api.Update(r.Context(), req.ApiPutUpdateInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// @summary 接口删除
// @tags    接口管理
// @produce json
// @param   id	path	int    true	"ID"
// @router  /manage/www/{id} [Delete]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiManage) Delete(r *ghttp.Request) {
	var req *model.ApiDeleteReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.Api.Delete(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}
