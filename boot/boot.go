package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gvalid"
	"www.lffwl.com/app/rule"
	_ "www.lffwl.com/packed"
)

func init() {

	//s := g.Server()

	// swagger
	//s.Plugin(&swagger.Swagger{})

	g.Cfg().SetFileName("config.toml")

	//注册验证规则
	gvalid.RegisterRule("in-array", rule.InArray)

}
