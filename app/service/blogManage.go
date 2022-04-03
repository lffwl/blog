package service

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"www.lffwl.com/app/dao"
	"www.lffwl.com/app/model"
	"www.lffwl.com/library/service"
)

var BlogManage = blogManage{}

type blogManage struct {
}

// 列表
func (s *blogManage) Index(ctx context.Context, input model.BlogManageIndexInput) (output *model.BlogManageIndexOutput, err error) {

	output = &model.BlogManageIndexOutput{
		Page:  input.Page,
		Limit: input.Limit,
	}

	m := dao.WwwBlog.Ctx(ctx)

	listModel := m.Fields(model.BlogManageList{}).OrderDesc(dao.WwwBlog.C.Id).Page(input.Page, input.Limit)

	if err := listModel.Scan(&output.List); err != nil {
		return output, err
	}

	//总条数
	if output.Total, err = m.Count(); err != nil {
		return output, err
	}

	if allType, err := BlogTypeManage.AllIndex(ctx); err != nil {
		return nil, err
	} else {
		output.Config = g.Map{
			"type_list": allType,
		}
	}

	return
}

// 新增
func (s *blogManage) Store(ctx context.Context, input model.BlogManageStoreInput) error {

	m := dao.WwwBlog.Ctx(ctx)

	//验证TypeId是否存在
	if err := service.PublicService.CheckFieldNotExist(dao.WwwBlogType.Ctx(ctx), dao.WwwBlogType.C.Id, input.TypeId, 0, "TypeId"); err != nil {
		return err
	}

	//验证title是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.WwwBlog.C.Title, input.Title, 0); err != nil {
		return err
	}

	if err := service.PublicService.Store(m, gconv.Map(input)); err != nil {
		return err
	}

	return nil
}

// 更新
func (s *blogManage) Update(ctx context.Context, input model.BlogManageUpdateInput) error {

	m := dao.WwwBlog.Ctx(ctx)

	//验证ID是否存在
	if err := service.PublicService.CheckIdExist(m, input.Id); err != nil {
		return err
	}

	//验证TypeId是否存在
	if err := service.PublicService.CheckFieldNotExist(dao.WwwBlogType.Ctx(ctx), dao.WwwBlogType.C.Id, input.TypeId, 0, "TypeId"); err != nil {
		return err
	}

	//验证title是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.WwwBlog.C.Title, input.Title, input.Id); err != nil {
		return err
	}

	if err := service.PublicService.Update(m, gconv.Map(input)); err != nil {
		return err
	}

	return nil
}

// 删除
func (s *blogManage) Delete(ctx context.Context, id int) error {

	m := dao.WwwBlog.Ctx(ctx)

	//验证ID是否存在
	if err := service.PublicService.CheckIdExist(m, id); err != nil {
		return err
	}

	if err := service.PublicService.Delete(m, id); err != nil {
		return err
	}

	return nil
}

// Show 详情
func (s *blogManage) Show(ctx context.Context, id int) (g.Map, error) {

	m := dao.WwwBlog.Ctx(ctx)

	//验证ID是否存在
	if err := service.PublicService.CheckIdExist(m, id); err != nil {
		return nil, err
	}

	if info, err := m.Where("id = ?", id).One(); err != nil {
		return nil, err
	} else {
		return gconv.Map(info), err
	}
}
