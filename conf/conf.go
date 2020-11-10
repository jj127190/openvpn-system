package conf

const (
	CsrfToken = "ClMG1ms5uWRSW2YZSALRMkwYIwl4URvG"
)


type Config struct {
	Run RunStat
	LogInfo  LogConf
}

type LogConf struct {
	LogName string
	LogPath string
	LogStat string
	LogLevel int32
}

type RunStat struct {
	StartPort string
}

