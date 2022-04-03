package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"www.lffwl.com/app/define"
)

var PublicService = publicService{}

type publicService struct {
}

// 检查ID是否存在
func (s *publicService) CheckIdExist(m *gdb.Model, id int) error {

	m = m.WherePri(id)

	//检查数据是否存在
	if i, err := m.Count(); err != nil {
		return err
	} else if i == 0 {
		return errors.New(g.I18n().Tf(m.GetCtx(), `id.not.exist`, id))
	}

	return nil

}

//检查字段是否存在
func (s *publicService) CheckFieldExist(m *gdb.Model, field string, value string, id int) error {

	// 检查ID是否存在
	if id != 0 {
		m = m.Where("id <> ?", id)
	}

	// 检查是否存在
	if i, err := m.FindCount(field, value); err != nil {
		return err
	} else if i > 0 {
		return errors.New(g.I18n().Tf(m.GetCtx(), `field.exist`, field, value))
	}

	return nil
}

//检查字段是否不存在
func (s *publicService) CheckFieldNotExist(m *gdb.Model, field string, value string, id int, fieldsName ...interface{}) error {

	// 检查ID是否存在
	if id != 0 {
		m = m.Where("id <> ?", id)
	}

	// 检查是否存在
	if i, err := m.FindCount(field, value); err != nil {
		return err
	} else if i == 0 {
		fieldName := field
		if len(fieldsName) > 0 {
			fieldName = gconv.String(fieldsName[0])
		}
		return errors.New(g.I18n().Tf(m.GetCtx(), `field.not.exist`, fieldName, value))
	}

	return nil
}

// 新增方法
func (s *publicService) Store(m *gdb.Model, data map[string]interface{}, fieldsEx ...interface{}) error {

	// 添加新增人
	data["CreatedId"] = gconv.Int(m.GetCtx().Value(define.JWT_ADMIN_ID))

	//是否存在需要排除字段
	if len(fieldsEx) > 0 {
		m = m.FieldsEx(strings.Join(gconv.Strings(fieldsEx[0]), ","))
	}

	//新增
	if _, err := m.Save(data); err != nil {
		return err
	}

	return nil
}

// 更新
func (s *publicService) Update(m *gdb.Model, data map[string]interface{}, fieldsEx ...interface{}) error {

	// 添加更新人
	data["UpdatedId"] = gconv.Int(m.GetCtx().Value(define.JWT_ADMIN_ID))

	var fieldsExStr string

	//是否存在需要排除字段
	if len(fieldsEx) > 0 {

		//追加ID排除
		fieldsEx[0] = append(gconv.SliceAny(fieldsEx[0]), "id")

		//转成字符串
		fieldsExStr = strings.Join(gconv.Strings(fieldsEx[0]), ",")

	} else {
		fieldsExStr = "id"
	}

	// 更新
	if _, err := m.FieldsEx(fieldsExStr).WherePri(data["Id"]).Data(data).Update(); err != nil {
		return err
	}

	return nil
}

// 删除方法
func (s *publicService) Delete(m *gdb.Model, id int) error {

	//事务
	if err := m.Transaction(m.GetCtx(), func(ctx context.Context, tx *gdb.TX) error {

		// 更新删除人 - 忽略更新时间 Unscoped
		if _, err := m.Data(g.Map{
			"DeletedId": gconv.Int(ctx.Value(define.JWT_ADMIN_ID)),
		}).Unscoped().WherePri(id).Update(); err != nil {
			return err
		}

		//删除
		if _, err := m.WherePri(id).Delete(); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
