package main

import (
	_ "www.lffwl.com/boot"
	_ "www.lffwl.com/router"

	"github.com/gogf/gf/frame/g"
)

// @title       www.lffwl.com - 个人博客
// @version     1.0
// @schemes     http
func main() {
	g.Server().Run()
}
