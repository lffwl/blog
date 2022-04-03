package service

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"www.lffwl.com/app/dao"
	"www.lffwl.com/app/model"
	"www.lffwl.com/library/service"
)

var BlogTypeManage = blogTypeManage{}

type blogTypeManage struct {
}

// 列表
func (s *blogTypeManage) Index(ctx context.Context, input model.BlogTypeManageIndexInput) (output *model.BlogTypeManageIndexOutput, err error) {

	output = &model.BlogTypeManageIndexOutput{
		Page:  input.Page,
		Limit: input.Limit,
	}

	m := dao.WwwBlogType.Ctx(ctx)

	listModel := m.Fields(model.BlogTypeManageList{}).Page(input.Page, input.Limit)

	if err := listModel.Scan(&output.List); err != nil {
		return output, err
	}

	// 总条数
	if output.Total, err = m.Count(); err != nil {
		return output, err
	}

	return
}

// 新增
func (s *blogTypeManage) Store(ctx context.Context, input model.BlogTypeManageStoreInput) error {

	m := dao.WwwBlogType.Ctx(ctx)

	// 检查name是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.WwwBlogType.C.Name, input.Name, 0); err != nil {
		return err
	}

	if err := service.PublicService.Store(m, gconv.Map(input)); err != nil {
		return err
	}

	return nil
}

// 更新
func (s *blogTypeManage) Update(ctx context.Context, input model.BlogTypeManageUpdateInput) error {

	m := dao.WwwBlogType.Ctx(ctx)

	// 检查Id是否存在
	if err := service.PublicService.CheckIdExist(m, input.Id); err != nil {
		return err
	}

	// 检查name是否存在
	if err := service.PublicService.CheckFieldExist(m, dao.WwwBlogType.C.Name, input.Name, input.Id); err != nil {
		return err
	}

	if err := service.PublicService.Update(m, gconv.Map(input)); err != nil {
		return err
	}

	return nil
}

// 删除
func (s *blogTypeManage) Delete(ctx context.Context, id int) error {

	m := dao.WwwBlogType.Ctx(ctx)

	// 检查Id是否存在
	if err := service.PublicService.CheckIdExist(m, id); err != nil {
		return err
	}

	//删除
	if err := service.PublicService.Delete(m, id); err != nil {
		return err
	}

	return nil
}

// AllIndex 全部列表
func (s *blogTypeManage) AllIndex(ctx context.Context) (*g.Map, error) {

	m := dao.WwwBlogType.Ctx(ctx)

	listModel := m.Fields(model.BlogTypeAllIndex{})

	var list []model.BlogTypeAllIndex

	if err := listModel.Scan(&list); err != nil {
		return nil, err
	}

	output := g.Map{}

	for _, v := range list {
		output[v.Id] = v.Name
	}

	return &output, nil
}
