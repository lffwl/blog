package rule

import (
	"context"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

// 数组验证
func InArray(ctx context.Context, rule string, value interface{}, message string, data interface{}) error {
	//获取配置信息
	ruleArr := strings.Split(rule, `:`)
	config := strings.Split(ruleArr[1], `,`)

	//检查类型是否存在配置中
	methodType := g.Cfg(config[0]).GetMap(config[1])

	for _, str := range gconv.Map(value) {
		if _, ok := methodType[gconv.String(str)]; !ok {
			return errors.New(g.I18n().Tf(ctx, `rule.in-array`, str))
		}
	}

	return nil
}
