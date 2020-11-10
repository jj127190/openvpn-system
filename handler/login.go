package handler

import (
	//"openvpn-system/common/share"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"openvpn-system/dao"
)

// 0表示立即销毁cookie

// -1表示关闭浏览器销毁cookie

// 大于1是设置的销毁时间，单位是秒，例如1小时后销毁，就可以设置60*60
func LogOut(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/welcome")
	c.SetCookie("VASessionId", "logout", 1, "/", "localhost", false, true)
	// c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Welcome(c *gin.Context) {
	cookie, err := c.Cookie("VASessionId")
	if err != nil {
		fmt.Println("cookie get err:", err)
	}
	fmt.Printf("Cookie value: %s \n", cookie)
	c.SetCookie("VASessionId", "logout", 3600, "/", "localhost", false, true)

	c.HTML(http.StatusOK, "login.html", gin.H{"loggin": "登录"})

}

func LoginDu(c *gin.Context) {
	username := c.DefaultPostForm("username", "username_nil")
	password := c.DefaultPostForm("password", "password_nil")
	//share.Logger.Info("登录:", username, password)
	fmt.Println("登录:", username, password)
	DBpassword, err := dao.QueryPass(username)
	fmt.Println(DBpassword)
	if err != nil {
		fmt.Println("查询失败;err:", err)
		c.HTML(http.StatusOK, "error_return_index.html", gin.H{"msg": "账号或密码输入错误！"})
		return
	}
	if password == DBpassword {
		c.SetCookie("VASessionId", "logging", 3600, "/", "localhost", false, true)
		//c.HTML(http.StatusOK, "index.html", gin.H{"index": "index"})
		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/rendto/index?username=%s", username))
		//c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("http://localhost:8080/rendto/index?username=%s",username))
	} else {
		c.HTML(http.StatusOK, "error_return_index.html", gin.H{"msg": "账号或密码输入错误！"})
	}
	//if username == "saber" && password == "123" {
	//	c.SetCookie("VASessionId", "logging", 3600, "/", "localhost", false, true)
	//c.HTML(http.StatusOK, "index.html", gin.H{"index": "index"})
	//	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/rendto/index")
	//} else {
	//		c.HTML(http.StatusOK, "error_return_index.html", gin.H{"msg": "账号或密码输入错误！"})
	//	}
	//todo 从数据库查找,并且加salt

}
