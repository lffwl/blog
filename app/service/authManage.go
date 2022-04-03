package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"www.lffwl.com/app/dao"
	"www.lffwl.com/app/define"
	"www.lffwl.com/app/model"
	"www.lffwl.com/library/service"
)

var AuthManage = authManage{}

type authManage struct {
}

//登录
func (s *authManage) Login(ctx context.Context, input model.AuthManageLoginInput) (map[string]interface{}, error) {

	m := dao.WwwAdmin.Ctx(ctx)

	//验证user_name是否存在
	if err := service.PublicService.CheckFieldNotExist(m, dao.WwwAdmin.C.UserName, input.UserName, 0); err != nil {
		return nil, err
	}

	//获取详情
	adminInfo, err := m.FieldsEx(dao.WwwAdmin.C.Password).FindOne(dao.WwwAdmin.C.UserName, input.UserName)
	if err != nil {
		return nil, err
	}

	//验证密码是否错误
	if !Admin.CheckPwd(gconv.String(adminInfo[dao.WwwAdmin.C.HashPassword]), input.Password) {
		return nil, errors.New(g.I18n().Tf(ctx, `admin.password.error`))
	}

	return g.Map{
		define.JWT_ADMIN_ID:     adminInfo[dao.WwwAdmin.C.Id],
		dao.WwwAdmin.C.UserName: adminInfo[dao.WwwAdmin.C.UserName],
	}, nil
}
