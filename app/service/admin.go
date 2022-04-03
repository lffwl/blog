package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"www.lffwl.com/app/dao"
	"www.lffwl.com/app/model"
	"www.lffwl.com/library/service"
)

// Admin服务
var Admin = adminService{}

type adminService struct{}

// 管理员列表
func (s *adminService) Index(ctx context.Context, input model.AdminGetListInput) (output *model.AdminGetListOutput, err error) {

	var m = dao.WwwAdmin.Ctx(ctx)

	output = &model.AdminGetListOutput{
		Page:  input.Page,
		Limit: input.Limit,
	}

	listModel := m.Fields(model.AdminGetList{}).Page(input.Page, input.Limit)

	if err := listModel.Scan(&output.List); err != nil {
		return output, err
	}

	//总条数
	output.Total, err = m.Count()

	return
}

// 管理员新增
func (s *adminService) Store(ctx context.Context, input model.AdminPostStoreInput) error {

	// 账号唯一性数据检查
	if err := s.CheckUserName(ctx, input.UserName, 0); err != nil {
		return err
	}

	//密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return err
	}

	//处理密码
	input.HashPassword = string(hash)

	m := dao.WwwAdmin.Ctx(ctx)

	//保存
	if err := service.PublicService.Store(m, gconv.Map(input)); err != nil {
		return err
	}
	return nil
}

// 检查账号是否符合规范(目前仅检查唯一性)
func (s *adminService) CheckUserName(ctx context.Context, userName string, id int) error {

	m := dao.WwwAdmin.Ctx(ctx)

	if id != 0 {
		m = m.Where("id <> ?", id)
	}

	if i, err := m.Unscoped().FindCount("user_name", userName); err != nil {
		return err
	} else if i > 0 {
		return errors.New(g.I18n().Tf(ctx, `admin.userName.exist`, userName))
	}

	return nil
}

// 验证密码是否匹配
func (s *adminService) CheckPwd(password, loginPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(loginPassword)) //验证（对比）
	if err != nil {
		return false
	}
	return true
}

//检查ID是否存在
func (s *adminService) CheckIdExist(ctx context.Context, id int) error {
	m := dao.WwwAdmin.Ctx(ctx).WherePri(id)
	//检查数据是否存在
	if i, err := m.Count(); err != nil {
		return err
	} else if i == 0 {
		return errors.New(g.I18n().Tf(ctx, `id.not.exist`, id))
	}
	return nil
}

// 更新管理员
func (s *adminService) Update(ctx context.Context, input model.AdminPutUpdateInput) error {

	//检查数据是否存在
	if err := s.CheckIdExist(ctx, input.Id); err != nil {
		return err
	}

	//排除的字段
	var fieldsEx []interface{}

	// 账号唯一性数据检查
	if err := s.CheckUserName(ctx, input.UserName, input.Id); err != nil {
		return err
	}

	//密码是否为空处理
	if input.Password != "" {

		//密码加密
		hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost) //加密处理
		if err != nil {
			return err
		}

		//处理密码
		input.HashPassword = string(hash)

	} else {
		//为空不更新密码
		fieldsEx = append(fieldsEx, dao.WwwAdmin.C.Password, dao.WwwAdmin.C.HashPassword)
	}

	if err := service.PublicService.Update(dao.WwwAdmin.Ctx(ctx), gconv.Map(input), fieldsEx); err != nil {
		return err
	}

	return nil
}

// 删除管理员
func (s *adminService) Delete(ctx context.Context, id int) error {

	//检查数据是否存在
	if err := s.CheckIdExist(ctx, id); err != nil {
		return err
	}

	m := dao.WwwAdmin.Ctx(ctx)

	//删除
	if err := service.PublicService.Delete(m, id); err != nil {
		return err
	}

	return nil
}

// 设置角色
func (s *adminService) SetRole(ctx context.Context, req model.AdminPostSetRoleReq) error {

	//检查ID是否存在
	if err := s.CheckIdExist(ctx, req.Id); err != nil {
		return err
	}

	//切割字符串
	apis := strings.Split(req.Roles, `,`)

	//检查角色集合是否存在
	if err := Role.CheckIdsExist(ctx, apis); err != nil {
		return err
	}

	var data []model.AdminPostSetRoleData

	for _, value := range apis {

		item := model.AdminPostSetRoleData{}
		item.AdminId = req.Id
		item.RoleId = gconv.Int(value)

		data = append(data, item)

	}

	m := dao.WwwAdminRole.Ctx(ctx)

	//事务
	if err := m.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		//删除
		if _, err := m.Where("admin_id = ?", req.Id).Delete(); err != nil {
			return err
		}

		// 保存
		if _, err := m.Data(data).Save(); err != nil {
			return err
		}

		return nil

	}); err != nil {
		return err
	}

	return nil

}

//获取管理员角色集合
func (s *adminService) GetAdminRoles(ctx context.Context, id int) (roles []string, err error) {

	if roles, err := dao.WwwAdminRole.Ctx(ctx).Where("admin_id IN(?)", id).Array(dao.WwwAdminRole.C.RoleId); err != nil {
		return nil, err
	} else {
		return gconv.Strings(roles), nil
	}
}

// GetRole 获取角色
func (s *adminService) GetRole(ctx context.Context, id int) (output *model.AdminGetRoleOutput, err error) {

	output = &model.AdminGetRoleOutput{}

	// 获取角色集合
	if output.Checked, err = s.GetAdminRoles(ctx, id); err != nil {
		return nil, err
	}

	if output.List, err = Role.AllIndex(ctx); err != nil {
		return nil, err
	}

	return
}
