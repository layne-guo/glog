package glog

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type LogData struct {
	timeStr string
	levelStr string
	module string
	fileName string
	funcName string
	lineNo int
	data string
}

type GLogBase struct {
	level int
	module string
}

func (l *GLogBase)writeLog(file *os.File, logData *LogData) {
	fmt.Fprintf(file, "%s %s %s (%s:%s:%d) msg:%s\n",logData.timeStr, logData.levelStr, logData.module,
		logData.fileName, logData.funcName, logData.lineNo, logData.data)
}

func (l *GLogBase)formatLogger(level int, module string, format string, args...interface{}) (*LogData) {

	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05.000")
	leveStr := getLevelStr(level)
	module = module
	fileName, funcName, lineNo := GetLineInfo(5)
	fileName = filepath.Base(fileName)
	data := fmt.Sprintf(format, args...)
	return &LogData{
		timeStr: timeStr,
		levelStr: leveStr,
		module: module,
		fileName: fileName,
		funcName: funcName,
		lineNo: lineNo,
		data: data,
	}
}