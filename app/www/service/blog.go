package service

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"www.lffwl.com/app/dao"
	"www.lffwl.com/app/www/model"
)

var Blog = blog{}

type blog struct {
}

func (s *blog) List(ctx context.Context, input model.BlogWwwGetIndexInput) (output *model.BlogWwwGetListOutput, err error) {

	output = &model.BlogWwwGetListOutput{
		Page:  input.Page,
		Limit: input.Limit,
	}

	m := dao.WwwBlog.Ctx(ctx)

	// 是否存在分类查找
	if input.TypeId > 0 {
		m = m.Where("type_id = ?", input.TypeId)
	}

	// 是否存在标题搜索
	if input.Title != "" {
		m = m.Where("title like ?", "%"+input.Title+"%")
	}

	// 是否存在name
	if input.Name != "" {
		id := BlogType.NameGetId(ctx, input.Name)
		m = m.Where("type_id = ?", id)
	}

	//获取列表
	listModel := m.Fields(model.BlogWwwGetList{}).OrderDesc(dao.WwwBlog.C.CreatedAt).Page(input.Page, input.Limit)

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

// Show 详情
func (s *blog) Show(ctx context.Context, id int) (g.Map, error) {

	if info, err := dao.WwwBlog.Ctx(ctx).Where("id = ?", id).One(); err != nil {
		return nil, err
	} else {
		return gconv.Map(info), nil
	}

}
