package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

func InitRouter(s *ghttp.Server) {
	initApiDocRouter(s)
	//s.BindHookHandler("/*", ghttp.HOOK_BEFORE_SERVE, showURL)
	initV1(s)
}

func initApiDocRouter(s *ghttp.Server) {

}

func initV1(s *ghttp.Server) {

	RouterApp := new(RouterGroup)
	PublicGroup := s.Group("")
	{
		RouterApp.InitIndexRouter(PublicGroup)

	}

}

func authHook(r *ghttp.Request) {
	switch r.Request.RequestURI { //登录相关免鉴权
	case "/v1/loginkey":
		return
	case "/v1/login":
		return
	default:
		r.Response.CORSDefault() //开启跨域
		//api.GfJWTMiddleware.MiddlewareFunc()(r) //鉴权中间件
	}
}

func showURL(r *ghttp.Request) {
	ctx := gctx.New()
	glog.Debug(ctx, "请求路径：", r.Method, r.Request.RequestURI)
	//r.Response.CORSDefault()
}
