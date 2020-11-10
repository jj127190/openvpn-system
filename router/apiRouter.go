package router

import (
	//"fmt"
	// "openvpn-system/common/middleware"
	"github.com/gin-gonic/gin"
	"openvpn-system/common/utils"
	"openvpn-system/handler"
)

func Apistribute(r *gin.Engine) {

	r.NoRoute(utils.NoResponse)

	//接口相关
	api := r.Group("/api/v1")
	// api.Use(middleware.LoginMiddler())
	{
		// index
		api.GET("/test", handler.TestHandler)
		// 首页
	}
}
