package router

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"openvpn-system/common/middleware"
	"openvpn-system/common/utils"
	"openvpn-system/handler"
)

func Distribute(r *gin.Engine) {

	r.NoRoute(utils.NoResponse)

	// 登录相关
	r.POST("/loggin", handler.LoginDu)
	r.GET("/logout", handler.LogOut)
	r.GET("/welcome", handler.Welcome)

	//页面渲染相关
	renderTem := r.Group("/rendto")
	renderTem.Use(middleware.LoginMiddler())
	{

		// index
		renderTem.GET("/index", handler.IndexHandler)
		// 首页
		renderTem.GET("/welcome", handler.WelcomeHandler)
		//登录
		renderTem.GET("/login", handler.LoginHandler)
		// 错误404
		renderTem.GET("/errorf", handler.F404Handler)

		//添加平台用户
		renderTem.GET("/addaccount", handler.AddAccountHandler)

		///////////////////////////////////////////////////
		// switch c.Request.Method {
		//  case "GET":
		// 	c.JSON(http.StatusOK, gin.H{"method": "GET"})
		//  case http.MethodPost:
		// 	c.JSON(http.StatusOK, gin.H{"method": "Post"})

		///////vpn账号管理页面
		renderTem.GET("/vpn", handler.VpnTemplateRender)
		renderTem.GET("/vpnUserAddRender", handler.VpnUserAddRender)

		//vpn权限修改页面
		renderTem.GET("/vpnperender", handler.VpnPertem)

		///////////////////////////////////////////////////

		//系统配置
		renderTem.GET("/sysconf", handler.SysConfHandler)

		//日志追踪
		renderTem.GET("/logtrack", handler.LogtrackHandler)
		//sql埋点
		renderTem.GET("/sqlbpoint", handler.SQLBpointHandler)
		// sql更改密码
		renderTem.GET("/sqlbpointadd", handler.SQLBpointAddHandler)

		//json Marshal
		renderTem.GET("/jsonmarshal", handler.JSONMarShalHandler)
		//开发日志
		renderTem.GET("/devlog", handler.DevlogHandler)

		//开发日志
		renderTem.GET("/admilis", handler.AdminListHandler)

		//权限组划分页面

		renderTem.GET("/pertem", handler.PerGrouptem)

		//添加权限组页面

		renderTem.GET("/pergroupaddtem", handler.Pergroupaddtem)

		//权限组编辑
		renderTem.GET("/pergroupedittem", handler.Pergroupedittem)

	}

	// api接口
	apiv := r.Group("/api/v1/")

	{
		apiv.POST("test", handler.TestHandler)

		apiv.POST("adminlist", handler.AdminlistHandler)

	}

	// 账号
	account := r.Group("/api/account/")
	{
		account.POST("adduserinfo", handler.AdduserinfoHandler)
		account.GET("del", handler.DeluserinfoHandler)

		account.POST("changepasswd", handler.ChangepasswdHandler)

		///////////////vpn账号
		account.POST("vpnuseradd", handler.VpnUserAddHandler)

		account.POST("vpnUserlist", handler.VpnUserlist)

	}

	domain := r.Group("/api/domain/")

	{
		domain.GET("getdomins", handler.GetDomains)
		domain.POST("addPerGroup", handler.AddPerGroup)
		domain.POST("dgshow", handler.DomainShow)
		domain.GET("winitgedomains", handler.Winitgedomains)
		domain.POST("addGroupajax", handler.AddGroupajax)
		domain.POST("editPerGroup", handler.EditPerGroup)

	}

	perg := r.Group("/api/pergroup/")
	{

		perg.POST("del", handler.Pergroupdel)
	}

}
