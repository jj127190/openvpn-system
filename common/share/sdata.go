package share

import (
	"VpnAudit/conf"
	"VpnAudit/common/logs"
	"fmt"
	"VpnAudit/common"
	"github.com/BurntSushi/toml"
)

var Conf conf.Config

var Logger common.LogMess
func init() {
	var confpath = "./conf/conf.toml"
	if _, err := toml.DecodeFile(confpath, &Conf); err != nil {
		fmt.Println("failed to read conf......")
		fmt.Println(err)
	}
	if Conf.LogInfo.LogStat == "on"{
		Logger = logs.NewLogger(Conf.LogInfo.LogName,Conf.LogInfo.LogPath,Conf.LogInfo.LogLevel)


	}
	// Logger.Warn("vpn audit system da。。。")
	// defer Logger.FileClose()
}
