package main

import (
	"openvpn-system/common/share"
	// "openvpopenvpn-system/common/vpntails"
	"fmt"
	"github.com/gin-gonic/gin"
	"openvpn-system/dao"
	"openvpn-system/router"
)

func main() {
	fmt.Println(" ... Platef starting ...")
	// go vpntails.StartTail()
	defer share.Logger.FileClose()
	defer dao.DB.Close()
	defer dao.GDB.Close()
	share.Logger.Warn("[vpn Audit_system start....]")
	gin.SetMode(gin.ReleaseMode)
	Rcontext := gin.Default()
	Rcontext.LoadHTMLGlob("templates/*")
	Rcontext.Static("/assets", "./assets")
	router.Distribute(Rcontext) //事项分支...
	Rcontext.Run(share.Conf.Run.StartPort)

}
