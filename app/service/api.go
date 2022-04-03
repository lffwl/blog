package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"www.lffwl.com/app/dao"
	"www.lffwl.com/app/model"
	"www.lffwl.com/library/service"
)

var Api = apiManage{}

type apiManage struct {
}

// api列表
func (s *apiManage) Index(ctx context.Context, input model.ApiGetIndexInput) (output *model.ApiGetIndexOutput, err error) {

	m := dao.WwwApi.Ctx(ctx)

	output = &model.ApiGetIndexOutput{}

	listModel := m.Fields(model.ApiGetIndexList{})

	if err := listModel.Scan(&output.List); err != nil {
		return output, err
	}

	//总条数
	output.Total, err = m.Count()

	output.Config = g.Map{
		"methods":   g.Cfg("api").GetMap("methods"),
		"menu_type": g.Cfg("api").GetMap("menu_type"),
	}

	return
}

// 检查是否重复请求地址和类型
func (s *apiManage) CheckRouter(ctx context.Context, router, methods string, id int) error {

	m := dao.WwwApi.Ctx(ctx)

	//更新排除自己
	if id > 0 {
		m = m.Where("id <> ?", id)
	}

	if i, err := m.Where(dao.WwwApi.C.Router, router).Where(dao.WwwApi.C.Methods, methods).Count(); err != nil {
		return err
	} else if i > 0 {
		return errors.New(g.I18n().Tf(ctx, `www.router.exist`, router, methods))
	}

	return nil
}

// Api新增
func (s *apiManage) Store(ctx context.Context, input model.ApiPostStoreInput) error {

	//检查PID是否存在
	if err := s.CheckPidExist(ctx, input.Pid); err != nil {
		return err
	}

	// 检查是否重复请求地址和类型
	if input.Router != "" && input.Methods != "" {
		if err := s.CheckRouter(ctx, input.Router, input.Methods, 0); err != nil {
			return err
		}
	}

	m := dao.WwwApi.Ctx(ctx)

	//验证name是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.WwwApi.C.Name, input.Name, 0); err != nil {
		return err
	}

	//验证key是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.WwwApi.C.Key, input.Key, 0); err != nil {
		return err
	}

	// 添加
	if err := service.PublicService.Store(dao.WwwApi.Ctx(ctx), gconv.Map(input)); err != nil {
		return err
	}

	return nil

}

//检查PID是否存在
func (s *apiManage) CheckPidExist(ctx context.Context, pid int) error {

	//检查ID是否存在
	if pid > 0 {
		if err := s.CheckIdExist(ctx, pid); err != nil {
			return errors.New(g.I18n().Tf(ctx, `pid.not.exist`, pid))
		}
	}

	return nil
}

//检查ID是否存在
func (s *apiManage) CheckIdExist(ctx context.Context, id int) error {
	m := dao.WwwApi.Ctx(ctx).WherePri(id)
	//检查数据是否存在
	if i, err := m.Count(); err != nil {
		return err
	} else if i == 0 {
		return errors.New(g.I18n().Tf(ctx, `id.not.exist`, id))
	}
	return nil
}

//检查ID集合是否存在
func (s *apiManage) CheckIdsExist(ctx context.Context, ids []string) error {
	m := dao.WwwApi.Ctx(ctx).WherePri(ids)
	//检查数据是否存在
	if i, err := m.Count(); err != nil {
		return err
	} else if i != len(ids) {
		return errors.New(g.I18n().Tf(ctx, `ids.not.exist`, ids))
	}
	return nil
}

// api更新
func (s *apiManage) Update(ctx context.Context, input model.ApiPutUpdateInput) error {

	//检查ID是否存在
	if err := s.CheckIdExist(ctx, input.Id); err != nil {
		return err
	}

	//检查PID是否存在
	if err := s.CheckPidExist(ctx, input.Pid); err != nil {
		return err
	}

	// 检查是否重复请求地址和类型
	if input.Router != "" && input.Methods != "" {
		if err := s.CheckRouter(ctx, input.Router, input.Methods, input.Id); err != nil {
			return err
		}
	}

	m := dao.WwwApi.Ctx(ctx)

	//验证name是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.WwwApi.C.Name, input.Name, input.Id); err != nil {
		return err
	}

	//验证key是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.WwwApi.C.Key, input.Key, input.Id); err != nil {
		return err
	}

	if err := service.PublicService.Update(dao.WwwApi.Ctx(ctx), gconv.Map(input)); err != nil {
		return err
	}

	return nil

}

//api删除
func (s *apiManage) Delete(ctx context.Context, id int) error {

	//检查ID是否存在
	if err := s.CheckIdExist(ctx, id); err != nil {
		return err
	}

	if err := service.PublicService.Delete(dao.WwwApi.Ctx(ctx), id); err != nil {
		return err
	}

	return nil
}

// 根据角色ID集合获取接口列表
func (s *apiManage) RoleIdsGetApiList(ctx context.Context, ids []string) (list []model.PermissionList, err error) {

	model := dao.WwwApi.Ctx(ctx).Fields(model.PermissionList{}).WherePri(ids)

	if err = model.Scan(&list); err != nil {
		return nil, err
	}

	return
}

// 获取所有的接口
func (s *apiManage) GetAllApiList(ctx context.Context) (list []model.PermissionList, err error) {
	model := dao.WwwApi.Ctx(ctx).Fields(model.PermissionList{})

	if err = model.Scan(&list); err != nil {
		return nil, err
	}

	return
}
