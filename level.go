package glog

const (
	GLogLevelDebug = iota
	GLogLevelTrace
	GLogLevelWarn
	GLogLevelError
	GLogLevelInfo
	GLogLevelFatal
)

const (
	GLogTypeFile = iota
	GLogTypeConsole
)

func getLevelStr(level int) string {
	switch level {
	case GLogLevelDebug:
		return "DEBUG"
	case GLogLevelTrace:
		return "TRACE"
	case GLogLevelInfo:
		return "INFO"
	case GLogLevelWarn:
		return "WARN"
	case GLogLevelError:
	    return "ERROR"
	case GLogLevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}