package service

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"www.lffwl.com/app/dao"
)

var BlogType = blogType{}

type blogType struct {
}

// List 列表
func (s *blogType) List(ctx context.Context) ([]g.Map, error) {

	if list, err := dao.WwwBlogType.Ctx(ctx).All(); err != nil {
		return nil, err
	} else {
		return gconv.Maps(list), nil
	}

}

// NameGetId 获取ID
func (s *blogType) NameGetId(ctx context.Context, name string) int {
	id, _ := dao.WwwBlogType.Ctx(ctx).Fields(dao.WwwBlogType.C.Id).Where("name = ?", name).Value()
	return gconv.Int(id)
}
