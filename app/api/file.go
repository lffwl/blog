package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"www.lffwl.com/app/service"
	"www.lffwl.com/library/response"
)

var FileManage = fileManage{}

type fileManage struct {
}

// Upload 文件上传
func (a *fileManage) Upload(r *ghttp.Request) {
	file := r.GetUploadFile("image")
	if file == nil {
		response.JsonExit(r, 1, "请选择需要上传的文件")
	}
	//获取列表
	if list, err := service.File.Upload(r.Context(), file); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		r.Response.WriteJson(g.Map{
			"errno": 0,
			"data": []g.Map{
				{
					"url":  list.Url,
					"alt":  "",
					"href": "",
				},
			},
		})
		r.Exit()
	}
}
