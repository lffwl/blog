package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"www.lffwl.com/app/dao"
	"www.lffwl.com/app/define"
	"www.lffwl.com/app/model"
)

var CenterManage = centerManage{}

type centerManage struct {
}

//个人详情
func (s *centerManage) Show(ctx context.Context) (output *model.CenterShow, err error) {

	// 获取当前登录的管理员ID
	admin_id := gconv.Int(ctx.Value(define.JWT_ADMIN_ID))

	output = &model.CenterShow{}

	//获取管理员基本信息
	if err = dao.WwwAdmin.Ctx(ctx).Fields(dao.WwwAdmin.C.Id, dao.WwwAdmin.C.UserName).FindScan(&output, admin_id); err != nil {
		return nil, err
	}

	//获取接口信息
	if output.Auth, err = s.GetAdminApis(ctx, admin_id); err != nil {
		return nil, err
	}

	return

}

// 获取管理员权限
func (s *centerManage) GetAdminApis(ctx context.Context, id int) (auth []model.PermissionList, err error) {

	var roles []string

	// 获取角色集合
	if roles, err = Admin.GetAdminRoles(ctx, id); err != nil {
		return auth, err
	}

	// 验证是否绑定了角色
	if len(roles) == 0 {
		return auth, errors.New(g.I18n().Tf(ctx, `admin.not-bind.role`))
	}

	//不是超级管理员需要获取权限
	if !Role.CheckIsSuperRole(roles) {

		var apis []string

		// 获取Api集合
		if apis, err = Role.RolesIdGetApiIds(ctx, roles); err != nil {
			return auth, err
		}

		// 获取Api集合
		if auth, err = Api.RoleIdsGetApiList(ctx, apis); err != nil {
			return auth, err
		}

	} else {

		// 获取Api集合
		if auth, err = Api.GetAllApiList(ctx); err != nil {
			return auth, err
		}
	}

	return
}
