package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"www.lffwl.com/app/api"
	"www.lffwl.com/app/define"
	"www.lffwl.com/app/service"
	serviceWeb "www.lffwl.com/app/www/service"
	"www.lffwl.com/app/www/www"
)

// 管理员验证
func ManageAuth(r *ghttp.Request) {
	api.AuthManage.MiddlewareFunc()(r)
	//设置上下文变量
	r.SetCtxVar(define.JWT_ADMIN_ID, r.GetParam(define.JWT_ADMIN_ID))
	r.Middleware.Next()
}

func init() {
	s := g.Server()

	if g.Cfg().GetString("ssl.certFile") != "" {
		s.EnableHTTPS(g.Cfg().GetString("ssl.certFile"), g.Cfg().GetString("ssl.keyFile"))
	}

	//设置全局中间件
	s.BindMiddlewareDefault(
		//service.Middleware.CORS,
		service.Middleware.ZHCN,
	)

	//manage接口模块
	s.Group("/manage", func(manage *ghttp.RouterGroup) {

		//授权
		manage.Group("/", func(manageAuth *ghttp.RouterGroup) {

			//jwt验证
			manageAuth.Middleware(ManageAuth)

			//管理员管理
			manageAuth.Group("/admin", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Admin.Index)
				group.POST("/", api.Admin.Store)
				group.PUT("/:id", api.Admin.Update)
				group.DELETE("/:id", api.Admin.Delete)
				group.POST("/:id/set-role", api.Admin.SetRole)
				group.GET("/:id/get-role", api.Admin.GetRole)
			})

			//接口管理
			manageAuth.Group("/api", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Api.Index)
				group.POST("/", api.Api.Store)
				group.PUT("/:id", api.Api.Update)
				group.DELETE("/:id", api.Api.Delete)
			})

			//角色管理
			manageAuth.Group("/role", func(group *ghttp.RouterGroup) {
				group.GET("/", api.Role.Index)
				group.POST("/", api.Role.Store)
				group.PUT("/:id", api.Role.Update)
				group.DELETE("/:id", api.Role.Delete)
				group.POST("/:id/set-www", api.Role.SetApi)
				group.GET("/:id/get-www", api.Role.GetApi)
			})

			//博客类型管理
			manageAuth.Group("/blog-type", func(group *ghttp.RouterGroup) {
				group.GET("/", api.BlogTypeManage.Index)
				group.POST("/", api.BlogTypeManage.Store)
				group.PUT("/:id", api.BlogTypeManage.Update)
				group.DELETE("/:id", api.BlogTypeManage.Delete)
			})

			//博客管理
			manageAuth.Group("/blog", func(group *ghttp.RouterGroup) {
				group.GET("/", api.BlogManage.Index)
				group.POST("/", api.BlogManage.Store)
				group.PUT("/:id", api.BlogManage.Update)
				group.DELETE("/:id", api.BlogManage.Delete)
				group.GET("/:id", api.BlogManage.Show)
			})

			//管理员中心
			manageAuth.Group("/center", func(group *ghttp.RouterGroup) {
				group.GET("/", api.CenterManage.Show)
			})

			//管理员中心
			manageAuth.Group("/file", func(group *ghttp.RouterGroup) {
				group.POST("/upload", api.FileManage.Upload)
			})
		})

		//授权
		manage.Group("/auth", func(group *ghttp.RouterGroup) {
			group.POST("/", api.AuthManage.LoginHandler)
			group.POST("/refresh_token", api.AuthManage.RefreshHandler)
			group.GET("/logout", api.AuthManage.LogoutHandler)
		})

	})

	//网站模块
	s.Group("/", func(web *ghttp.RouterGroup) {

		//中间件
		web.Middleware(serviceWeb.Public.WebBase)

		web.GET("/*name", www.Home.Index)
		web.GET("/:name/{id}.html", www.Home.Show)
		web.GET("/about.html", www.Home.About)
	})

}
