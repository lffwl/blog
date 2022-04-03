package api

import (
	"github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"time"
	"www.lffwl.com/app/define"
	"www.lffwl.com/app/model"
	"www.lffwl.com/app/service"
	"www.lffwl.com/library/response"
)

var (
	// The underlying JWT middleware.
	AuthManage *jwt.GfJWTMiddleware
)

// Initialization function,
// rewrite this function to customized your own JWT settings.
func init() {
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:            "www_lffwl_com_manage",
		SigningAlgorithm: "HS256",
		Key:              []byte("secret key"),
		Timeout:          time.Minute * 60 * 4,
		MaxRefresh:       time.Minute * 60 * 24,
		IdentityKey:      define.JWT_ADMIN_ID,
		TokenLookup:      "header: Authorization",
		TokenHeadName:    "Bearer",
		TimeFunc: func() time.Time {
			return gconv.Time(gtime.Now())
		},
		Authenticator:   Authenticator,
		LoginResponse:   LoginResponse,
		RefreshResponse: RefreshResponse,
		LogoutResponse:  LogoutResponse,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	AuthManage = authMiddleware
}

func Authenticator(r *ghttp.Request) (interface{}, error) {
	var req *model.AuthManageLoginReq

	//参数验证
	if err := r.Parse(&req); err != nil {
		return nil, err
	}

	//获取信息
	if admin, err := service.AuthManage.Login(r.Context(), req.AuthManageLoginInput); err != nil {
		return nil, err
	} else if admin != nil {
		return admin, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	response.JsonExit(r, 0, "登录成功", g.Map{
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
}

func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	response.JsonExit(r, 0, "刷新成功", g.Map{
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
}

func LogoutResponse(r *ghttp.Request, code int) {
	response.JsonExit(r, 0, "退出成功")
}

func Unauthorized(r *ghttp.Request, code int, message string) {
	response.JsonExit(r, code, message)
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims[AuthManage.IdentityKey]
}
