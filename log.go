package glog

var logger GLog = NewGLog(GLogTypeConsole, GLogLevelDebug, "", "default")
type GLog interface {
	Init() error
	LogDebug(fmt string, args...interface{})
	LogTrace(fmt string, args...interface{})
	LogInfo(fmt string, args...interface{})
	LogWarn(fmt string, args...interface{})
	LogError(fmt string, args...interface{})
	LogFatal(fmt string, args...interface{})
	SetLevel(level int)
	Close()
}

func NewGLog(logType , level int, filename, module string) GLog {
	var logger GLog
	switch (logType) {
	case GLogTypeFile:
		logger = NewGFile(level, filename, module)
	case GLogTypeConsole:
		logger = NewGConsole(level, module)
	default:
		logger = NewGFile(level, filename, module)
	}
	return logger
}


func Init(logType , level int, filename, module string) (error) {
	logger = NewGLog(logType, level, filename, module)
	return logger.Init()
}

func LogDebug(fmt string, args ...interface{}) {
	logger.LogDebug(fmt, args...)
}
func LogTrace(fmt string, args...interface{}) {
	logger.LogTrace(fmt, args...)
}
func LogInfo(fmt string, args...interface{}) {
	logger.LogInfo(fmt, args...)
}
func LogWarn(fmt string, args...interface{}) {
	logger.LogWarn(fmt, args...)
}
func LogError(fmt string, args...interface{}) {
	logger.LogError(fmt, args...)
}
func LogFatal(fmt string, args...interface{}) {
	logger.LogFatal(fmt, args...)
}
func SetLevel(level int) {
	logger.SetLevel(level)
}
func Close() {
    logger.Close()
}