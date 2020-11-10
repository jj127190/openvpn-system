package router

import (
	//"fmt"
	// "VpnAudit/common/middleware"
	"VpnAudit/handler"
	"VpnAudit/common/utils"
	"github.com/gin-gonic/gin"
)

func Apistribute(r *gin.Engine){
	
	
	r.NoRoute(utils.NoResponse)
	
	//接口相关
	api := r.Group("/api/v1")
	// api.Use(middleware.LoginMiddler())
	{
		// index
		api.GET("/test",handler.TestHandler)
		// 首页
	}
}