package common

import (
	"os"
	"github.com/sirupsen/logrus"
	
)



type LogMess interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	FileClose()
}

type LogInfo struct {
	LogName    string
	LogPath    string
	FileSystem *os.File
	Logger     *logrus.Logger
}
