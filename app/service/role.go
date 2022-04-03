package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"www.lffwl.com/app/dao"
	"www.lffwl.com/app/model"
	"www.lffwl.com/library/helper"
	"www.lffwl.com/library/service"
)

var Role = roleManage{}

type roleManage struct {
}

// 列表
func (s *roleManage) Index(ctx context.Context, input model.RoleGetIndexInput) (output *model.RoleGetListOutput, err error) {

	output = &model.RoleGetListOutput{
		Page:  input.Page,
		Limit: input.Limit,
	}

	m := dao.WwwRole.Ctx(ctx)

	//获取列表
	listModel := m.Fields(model.RoleGetList{}).Page(input.Page, input.Limit)

	//类型转换
	if err := listModel.Scan(&output.List); err != nil {
		return output, err
	}

	//获取条数
	if output.Total, err = m.Count(); err != nil {
		return output, err
	}

	return
}

//	检查角色名称是否唯一
func (s *roleManage) CheckName(ctx context.Context, name string, id int) error {

	m := dao.WwwRole.Ctx(ctx)

	//是否需要排除ID
	if id > 0 {
		m = m.Where("id <> ?", id)
	}

	if i, err := m.Where(dao.WwwRole.C.Name, name).Count(); err != nil {
		return err
	} else if i > 0 {
		return errors.New(g.I18n().Tf(ctx, `role.name.exist`, name))
	}

	return nil
}

//	新增
func (s *roleManage) Store(ctx context.Context, input model.RolePostStoreInput) error {

	//检查角色名称时候唯一
	if err := s.CheckName(ctx, input.Name, 0); err != nil {
		return err
	}

	//新增
	if err := service.PublicService.Store(dao.WwwRole.Ctx(ctx), gconv.Map(input)); err != nil {
		return err
	}

	return nil
}

// 检查ID是否存在
func (s *roleManage) CheckIdExist(ctx context.Context, id int) error {

	if i, err := dao.WwwRole.Ctx(ctx).WherePri(id).Count(); err != nil {
		return err
	} else if i == 0 {
		return errors.New(g.I18n().Tf(ctx, `id.not.exist`, id))
	}

	return nil
}

// 更新
func (s *roleManage) Update(ctx context.Context, input model.RolePutUpdateInput) error {

	//检查ID是否存在
	if err := s.CheckIdExist(ctx, input.Id); err != nil {
		return err
	}

	//检查角色名称时候唯一
	if err := s.CheckName(ctx, input.Name, input.Id); err != nil {
		return err
	}

	//更新
	if err := service.PublicService.Update(dao.WwwRole.Ctx(ctx), gconv.Map(input)); err != nil {
		return err
	}

	return nil
}

// 删除
func (s *roleManage) Delete(ctx context.Context, id int) error {

	//检查ID是否存在
	if err := s.CheckIdExist(ctx, id); err != nil {
		return err
	}

	if err := service.PublicService.Delete(dao.WwwRole.Ctx(ctx), id); err != nil {
		return err
	}

	return nil
}

// 设置权限
func (s *roleManage) SetApi(ctx context.Context, req model.RolePostSetApiReq) error {

	//检查ID是否存在
	if err := s.CheckIdExist(ctx, req.Id); err != nil {
		return err
	}

	//切割字符串
	apis := strings.Split(req.Apis, `,`)

	//检查ID权限是否存在
	if err := Api.CheckIdsExist(ctx, apis); err != nil {
		return err
	}

	var data []model.RolePostSetApiData

	for _, value := range apis {

		item := model.RolePostSetApiData{}
		item.RoleId = req.Id
		item.ApiId = gconv.Int(value)

		data = append(data, item)

	}

	m := dao.WwwRoleApi.Ctx(ctx)

	//事务
	if err := m.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {

		//删除
		if _, err := m.Where("role_id = ?", req.Id).Delete(); err != nil {
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

// 检查ID集合是否存在
func (s *roleManage) CheckIdsExist(ctx context.Context, ids []string) error {
	m := dao.WwwRole.Ctx(ctx).WherePri(ids)
	//检查数据是否存在
	if i, err := m.Count(); err != nil {
		return err
	} else if i != len(ids) {
		return errors.New(g.I18n().Tf(ctx, `ids.not.exist`, ids))
	}
	return nil
}

// 根据角色Id集合，获取ApiIds
func (s *roleManage) RolesIdGetApiIds(ctx context.Context, ids []string) (apis []string, err error) {
	if apis, err := dao.WwwRoleApi.Ctx(ctx).Where("role_id in (?)", ids).Array(dao.WwwRoleApi.C.ApiId); err != nil {
		return nil, err
	} else {
		return gconv.Strings(apis), nil
	}
}

// 检查是否是超级管理员
func (s *roleManage) CheckIsSuperRole(ids []string) bool {
	superRoleId := g.Cfg("admin").GetString("SuperRoleId")
	return helper.Array.In(superRoleId, ids)
}

// 角色详情
func (s *roleManage) Show(ctx context.Context, id int) (role *model.RoleShow, err error) {

	role = &model.RoleShow{}

	info, err := dao.WwwRole.Ctx(ctx).WherePri(id).Fields(dao.WwwRole.C.Id, dao.WwwRole.C.Name).One()
	if err != nil {
		return nil, err
	}

	if err = info.Struct(&role); err != nil {
		return nil, err
	}

	return
}

// GetApi 查看角色权限
func (s *roleManage) GetApi(ctx context.Context, id int) (output *model.RoleGetApiOutput, err error) {

	output = &model.RoleGetApiOutput{}

	// 获取Api集合
	if output.Checked, err = s.RolesIdGetApiIds(ctx, gconv.Strings(id)); err != nil {
		return nil, err
	}

	if output.List, err = Api.GetAllApiList(ctx); err != nil {
		return nil, err
	}

	return
}

// 全部列表
func (s *roleManage) AllIndex(ctx context.Context) (output []model.RoleGetList, err error) {

	output = []model.RoleGetList{}

	m := dao.WwwRole.Ctx(ctx)

	//获取列表
	listModel := m.Fields(model.RoleGetList{})

	//类型转换
	if err := listModel.Scan(&output); err != nil {
		return output, err
	}

	return
}
