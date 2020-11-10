package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginMiddler() gin.HandlerFunc {
	return func(c *gin.Context) {

		username := c.DefaultPostForm("username", "username_nil")
		password := c.DefaultPostForm("password", "password_nil")
		if username == "username_nil" && password == "password_nil" {
			cookie, err := c.Request.Cookie("VASessionId")
			if err != nil {
				fmt.Println("get cookie error", err)
				c.HTML(http.StatusOK, "login.html", gin.H{"loggin": "登录"})
			//	c.HTML(http.StatusOK, "error_return_index.html", gin.H{"loggin": "登录"})
				c.Abort()
				return
			}
			value := cookie.Value
			if value == "logging" {
				c.Next()
				return
			}
			c.HTML(http.StatusUnauthorized, "error.html", gin.H{"msg": "Vpn Audit sys Unauthorized...您已经是登出状态;请重新登录"})
			// c.JSON(http.StatusUnauthorized, gin.H{
			//	"error": "Vpn Audit sys Unauthorized...您已经是登出状态;请重新登录！",
			//})
			c.Abort()
			return
		}

	}
}
