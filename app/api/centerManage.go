package api

import (
	"github.com/gogf/gf/net/ghttp"
	"www.lffwl.com/app/service"
	"www.lffwl.com/library/response"
)

var CenterManage = centerManage{}

type centerManage struct {
}

func (a *centerManage) Show(r *ghttp.Request) {

	//获取列表
	if data, err := service.CenterManage.Show(r.Context()); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok", data)
	}
}
