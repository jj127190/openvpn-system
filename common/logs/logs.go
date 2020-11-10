package logs

import (
	"VpnAudit/common"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var Logger common.LogMess

type LogInfo struct {
	LogName    string
	LogPath    string
	FileSystem *os.File
	Logger     *logrus.Logger
}

func GetLogInfo(name, path string, file *os.File, log *logrus.Logger) *LogInfo {
	return &LogInfo{
		LogName:    name,
		LogPath:    path,
		FileSystem: file,
		Logger:     log,
	}
}

func NewLogger(LogName, LogPath string, LogLevel int32) *LogInfo {

	
	file, err := os.OpenFile(LogPath+LogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("faile to open file...")
		fmt.Println(err)
	}

	loggers := logrus.New()

	loggers.Out = file

	if LogLevel == 0 {
		loggers.SetLevel(logrus.DebugLevel)
	}
	if LogLevel == 1 {
		loggers.SetLevel(logrus.InfoLevel)
	}
	if LogLevel == 2 {
		loggers.SetLevel(logrus.WarnLevel)
	}

	// DebugLevel # LogLevel 0:debug; 1:info; 2:warning; 3:danger
	loggers.SetFormatter(&logrus.TextFormatter{})

	logstrus := GetLogInfo(LogName, LogPath, file, loggers)
	return logstrus
}

func (l *LogInfo) Debug(rest ...interface{}) {

	if len(rest) == 1 {
		rest = append(rest, "vpn")
		rest = append(rest, "...info...")
	}

	l.Logger.WithFields(
		logrus.Fields{
			rest[1].(string): rest[2].(string),
		}).Info(rest[0].(string))

}

func (l *LogInfo) Info(rest ...interface{}) {

	if len(rest) == 1 {
		rest = append(rest, "vpn")
		rest = append(rest, "...info...")
	}
	l.Logger.WithFields(
		logrus.Fields{
			rest[1].(string): rest[2].(string),
		}).Warn(rest[0].(string))
}

func (l *LogInfo) Warn(rest ...interface{}) {

	if len(rest) == 1 {
		rest = append(rest, "vpn")
		rest = append(rest, "...info...")
	}

	l.Logger.WithFields(
		logrus.Fields{
			rest[1].(string): rest[2].(string),
		}).Warn(rest[0].(string))
}

func (l *LogInfo) FileClose() {
	l.FileSystem.Close()
}

