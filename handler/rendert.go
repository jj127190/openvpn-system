package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//index
func IndexHandler(c *gin.Context) {
	username := c.Query("username")
	c.HTML(http.StatusOK, "index.html", gin.H{"username": username})
}

//首页
func WelcomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{"index": "index"})
}

//登录
func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"loggin": "登录"})
}

//404
func F404Handler(c *gin.Context) {
	c.HTML(http.StatusOK, "error.html", gin.H{"index": "index"})
}

//系统配置
func SysConfHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{"index": "index"})
}

//权限页面展示

func PerGrouptem(c *gin.Context) {
	c.HTML(http.StatusOK, "per_group.html", gin.H{"index": "index"})
}

//权限组添加

func Pergroupaddtem(c *gin.Context) {

	c.HTML(http.StatusOK, "per_add_temselect.html", gin.H{"index": "index"})

	// c.HTML(http.StatusOK, "per_add_tem.html", gin.H{"index": "index"})
}

//权限组编辑

func Pergroupedittem(c *gin.Context) {
	ID := c.DefaultQuery("ID", "null")
	GroupName := c.DefaultQuery("GroupName", "null")
	if ID != "null" && GroupName != "nil" {
		fmt.Println("编辑...组")
		fmt.Println(ID, GroupName)

	}

	c.HTML(http.StatusOK, "per_edit_temselect.html", gin.H{"gid": ID, "GroupName": GroupName}) //interface err struct 都可以是nil
	// c.HTML(http.StatusOK, "per_add_tem.html", gin.H{"index": "index"})
}

//日志追踪
func LogtrackHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "logtrack.html", gin.H{"index": "index"})
}

//防火墙规则sql查询
func SQLBpointHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "sqlbpoint.html", gin.H{"index": "index"})
}

//添加平台用户

func AddAccountHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-add-define.html", gin.H{"adminadd": "adminadd"})
	// c.HTML(http.StatusOK, "per-add-tem.html", gin.H{"adminadd": "adminadd"})

}

///////////////////// vpn账号相关页面
func VpnTemplateRender(c *gin.Context) {
	c.HTML(http.StatusOK, "vpnUserShowlist.html", gin.H{"vpnadmin": "vpnadmin"})

}

/////////////////////vpn账号添加
func VpnUserAddRender(c *gin.Context) {
	c.HTML(http.StatusOK, "vpnUserAdd.html", gin.H{"vpnuseradd": "vpnuseradd"})

}

func VpnPertem(c *gin.Context) {
	VpnAccount := c.DefaultQuery("username", "null")
	fmt.Println(VpnAccount)
	c.HTML(http.StatusOK, "vpnperedit.html", gin.H{"VpnAccount": VpnAccount})
}

func SQLBpointAddHandler(c *gin.Context) {
	username := c.DefaultQuery("username", "null")
	// fmt.Println(username)
	if username != "null" {
		c.HTML(http.StatusOK, "sqlbpointAdd.html", gin.H{"username": username})
	}

}

//json marshal
func JSONMarShalHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "jsonMarshal.html", gin.H{"index": "index"})
}

//日志开发
func DevlogHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "log.html", gin.H{"index": "index"})
}

// 用户管理
func AdminListHandler(c *gin.Context) {
	// fmt.Println("用户管理........列表页面")
	c.HTML(http.StatusOK, "admin_list.html", gin.H{"index": "index"})
}
