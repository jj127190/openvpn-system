package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"openvpn-system/common/share"
	"openvpn-system/router"
)

// getcwd  C:\Users\saber\Desktop\资料\go\gopath\src\VpnAudit\server\main

func main() {
	fmt.Println("Blingabc BI_Platef starting ...")
	defer share.Logger.FileClose()
	share.Logger.Warn("main...sever 22222222222222222222222222")
	gin.SetMode(gin.ReleaseMode)
	Rcontext := gin.Default()
	// Rcontext.LoadHTMLGlob("templates/*")
	// Rcontext.Static("/assets", "./assets")
	router.Apistribute(Rcontext) //事项分支...
	Rcontext.Run(":8081")
}
