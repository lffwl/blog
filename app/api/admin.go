package api

import (
	"github.com/gogf/gf/net/ghttp"
	"www.lffwl.com/app/model"
	"www.lffwl.com/app/service"
	"www.lffwl.com/library/response"
)

var Admin = adminManage{}

type adminManage struct {
}

// @summary 管理员列表
// @tags    管理员管理
// @produce json
// @param   entity query model.AdminGetListReq false "列表请求"
// @router  /manage/admin [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *adminManage) Index(r *ghttp.Request) {
	var (
		req *model.AdminGetListReq
	)
	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	//获取列表
	if list, err := service.Admin.Index(r.Context(), req.AdminGetListInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", list)
	}
}

// @summary 添加管理员
// @tags    管理员管理
// @produce json
// @accept mpfd
// @param   entity formData model.AdminPostStoreSwag ture "添加管理员"
// @router  /manage/admin [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *adminManage) Store(r *ghttp.Request) {
	var (
		req *model.AdminPostStoreReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Admin.Store(r.Context(), req.AdminPostStoreInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")

}

// @summary 更新管理员
// @tags    管理员管理
// @produce json
// @accept x-www-form-urlencoded
// @param   id	path	int    true	"管理员ID"
// @param   entity formData model.AdminPostUpdateSwag ture "更新管理员"
// @router  /manage/admin/{id} [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *adminManage) Update(r *ghttp.Request) {
	var (
		req *model.AdminPutUpdateReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Admin.Update(r.Context(), req.AdminPutUpdateInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")

}

// @summary 删除管理员
// @tags    管理员管理
// @produce json
// @param   id	path	int    true	"管理员ID"
// @router  /manage/admin/{id} [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *adminManage) Delete(r *ghttp.Request) {
	var (
		req model.AdminDeleteReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Admin.Delete(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")

}

// @summary 管理员角色设置
// @tags    管理员管理
// @produce json
// @param   id	path	int    true	"管理员ID""
// @param   entity formData model.AdminPostSetRoleReq ture "更新管理员"
// @router  /manage/admin/{id}/set-role [Post]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *adminManage) SetRole(r *ghttp.Request) {

	var req *model.AdminPostSetRoleReq

	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.Admin.SetRole(r.Context(), *req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// @summary 管理员角色设置查看
// @tags    管理员管理
// @produce json
// @param   id	path	int    true	"管理员ID""
// @param   entity formData model.AdminPostSetRoleReq ture "更新管理员"
// @router  /manage/admin/{id}/get-role [Post]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *adminManage) GetRole(r *ghttp.Request) {

	var req *model.AdminDeleteReq

	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if list, err := service.Admin.GetRole(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", list)
	}
}
