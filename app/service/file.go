package service

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"www.lffwl.com/app/model"
)

var File = file{}

type file struct {
}

// Upload 文件上传
func (s *file) Upload(ctx context.Context, file *ghttp.UploadFile) (output *model.FileUploadOutput, err error) {

	output = &model.FileUploadOutput{}

	//文件上传
	if output.Url, err = file.Save(g.Cfg().GetString("upload.path"), true); err != nil {
		return nil, err
	}

	output.Url = "/file/" + output.Url

	return
}
