package main

import (
	"VpnAudit/common/share"
	// "VpnAudit/common/vpntails"
	"VpnAudit/dao"
	"VpnAudit/router"
	"fmt"
	"github.com/gin-gonic/gin"
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
