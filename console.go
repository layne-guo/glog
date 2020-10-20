package glog

import (
	"os"
)

type GConsole struct {
	*GLogBase
}

func NewGConsole(level int, module string) GLog {
	logger := &GConsole{}
	logger.GLogBase = &GLogBase{
		level: level,
		module: module,
	}
	return logger
}

func (c *GConsole) Init() (err error) {
	return nil
}

func (c *GConsole)LogDebug(format string, args...interface{}) {
	if c.level > GLogLevelDebug {
		return
	}
	logData := c.formatLogger(GLogLevelDebug, c.module, format, args...)
    c.writeLog(os.Stdout, logData)
}

func (c *GConsole)LogTrace(format string, args...interface{}) {
	if c.level > GLogLevelTrace {
		return
	}
	logData := c.formatLogger(GLogLevelTrace, c.module, format, args...)
	c.writeLog(os.Stdout, logData)

}

func (c *GConsole)LogInfo(format string, args...interface{}) {
	if c.level > GLogLevelInfo {
		return
	}
	logData := c.formatLogger(GLogLevelInfo, c.module, format, args...)
	c.writeLog(os.Stdout, logData)
}

func (c *GConsole)LogWarn(format string, args...interface{}) {
	if c.level > GLogLevelWarn {
		return
	}
	logData := c.formatLogger(GLogLevelWarn, c.module, format, args...)
	c.writeLog(os.Stdout, logData)

}

func (c *GConsole)LogError(format string, args...interface{}) {
	if c.level > GLogLevelError {
		return
	}
	logData := c.formatLogger(GLogLevelError, c.module, format, args...)
	c.writeLog(os.Stdout, logData)

}

func (c *GConsole)LogFatal(format string, args...interface{}) {
	if c.level > GLogLevelFatal {
		return
	}
	logData := c.formatLogger(GLogLevelFatal, c.module, format, args...)
	c.writeLog(os.Stdout, logData)

}

func (c *GConsole)SetLevel(level int) {
	c.level = level
}

func (c *GConsole)Close() {
	return
}
