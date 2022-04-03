package api

import (
	"github.com/gogf/gf/net/ghttp"
	"www.lffwl.com/app/model"
	"www.lffwl.com/app/service"
	"www.lffwl.com/library/response"
)

var Role = roleManage{}

type roleManage struct {
}

// @summary 角色列表
// @tags    角色管理
// @produce json
// @param   entity query model.RoleGetIndexReq false "列表请求"
// @router  /manage/role [Get]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *roleManage) Index(r *ghttp.Request) {
	var req *model.RoleGetIndexReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	//获取列表
	if list, err := service.Role.Index(r.Context(), req.RoleGetIndexInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", list)
	}
}

// @summary 角色新增
// @tags    角色管理
// @produce json
// @param   entity query model.RolePostStoreReq false "新增请求"
// @router  /manage/role [Post]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *roleManage) Store(r *ghttp.Request) {
	var req *model.RolePostStoreReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//获取列表
	if err := service.Role.Store(r.Context(), req.RolePostStoreInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// @summary 角色更新
// @tags    角色管理
// @produce json
// @param   id	path	int    true	"ID"
// @param   entity query model.RolePutUpdateReqSwag false "更新请求"
// @router  /manage/role/{id} [Put]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *roleManage) Update(r *ghttp.Request) {
	var req *model.RolePutUpdateReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//获取列表
	if err := service.Role.Update(r.Context(), req.RolePutUpdateInput); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// @summary 角色删除
// @tags    角色管理
// @produce json
// @param   id	path	int    true	"ID"
// @router  /manage/role/{id} [Delete]
// @param   entity query model.RoleDeleteReq false "设置权限请求"
// @success 200 {object} response.JsonResponse "执行结果"
func (a *roleManage) Delete(r *ghttp.Request) {
	var req *model.RoleDeleteReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//获取列表
	if err := service.Role.Delete(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// @summary 角色权限设置
// @tags    角色管理
// @produce json
// @param   id	path	int    true	"ID"
// @router  /manage/role/{id}/set-www [Post]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *roleManage) SetApi(r *ghttp.Request) {
	var req *model.RolePostSetApiReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//获取列表
	if err := service.Role.SetApi(r.Context(), *req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "ok")
}

// @summary 角色权限详情
// @tags    角色管理
// @produce json
// @param   id	path	int    true	"ID"
// @router  /manage/role/{id}/get-www [Get]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *roleManage) GetApi(r *ghttp.Request) {
	var req *model.RoleDeleteReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	//获取列表
	if output, err := service.Role.GetApi(r.Context(), req.Id); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", output)
	}

}
