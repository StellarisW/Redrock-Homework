package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"main/app/api"
)

type IndexRouter struct{}

func (r *IndexRouter) InitIndexRouter(Router *ghttp.RouterGroup) {
	indexRouter := Router.Group("/")
	indexApi := api.Index()
	{
		indexRouter.GET("/", indexApi.Get)
		indexRouter.POST("/", indexApi.Update)
		//baseRouter.POST("login", baseApi.Login)
		//baseRouter.POST("captcha", baseApi.Captcha)
	}
}
