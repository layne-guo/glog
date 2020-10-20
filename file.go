package glog

import(
	"fmt"
	"os"
	"sync"
	"time"
)


type GFile struct {
	fileName string
	file *os.File
	*GLogBase
	logChan chan *LogData
	wg *sync.WaitGroup
	curHour int
}

func NewGFile(level int, filename, module string) GLog {
	logger := &GFile{
		fileName :filename,
	}
	logger.GLogBase = &GLogBase{
		level: level,
		module: module,
	}
	logger.curHour = time.Now().Hour()
	logger.wg = &sync.WaitGroup{}
	logger.logChan = make(chan *LogData, 100000)
	logger.wg.Add(1)
	go logger.syncLog()
	return logger
}
// 把chan中的 日志写至文件
func (f *GFile)syncLog() {
	for data := range f.logChan {
		f.splitLog()
		f.writeLog(f.file, data)
	}

	f.wg.Done()
}

func (f *GFile) splitLog() {
	now := time.Now()
	if now.Hour() == f.curHour {
		return
	}
	f.curHour = now.Hour()
	_ = f.file.Sync()
	_ =f.file.Close()
	newFileName := fmt.Sprintf("%s-%04d-%02d-%02d-%02d", f.fileName,
		now.Year(), now.Month(),now.Day(), now.Hour()-1)
	err := os.Rename(f.fileName, newFileName)
	if err != nil {
	}
	_ = f.Init()
}

func (f *GFile) Init() (err error) {
	f.file, err = os.OpenFile(f.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return
	}
	return
}
// 将格式化后的日志数据写至 chan
func (f *GFile) writeToChan(level int, module string, format string, args ...interface{}) {
	logData := f.formatLogger(level, module, format, args...)
	select {
	case f.logChan <- logData:
	default:
	}
}

func (f *GFile)LogDebug(format string, args...interface{}) {
	if f.level > GLogLevelDebug {
		return
	}
    f.writeToChan(GLogLevelDebug, f.module, format, args...)
	//f.writeLog(f.file, logData)
}

func (f *GFile)LogTrace(format string, args...interface{}) {
	if f.level > GLogLevelTrace {
		return
	}
	f.writeToChan(GLogLevelTrace, f.module, format, args...)
	/*
	logData := f.formatLogger(GLogLevelTrace, f.module, format, args...)
	f.writeLog(f.file, logData)
	 */
}

func (f *GFile)LogInfo(format string, args...interface{}) {
	if f.level > GLogLevelInfo {
		return
	}
	f.writeToChan(GLogLevelInfo, f.module, format, args...)
	/*
	logData := f.formatLogger(GLogLevelInfo, f.module, format, args...)
	f.writeLog(f.file, logData)
	 */
}

func (f *GFile)LogWarn(format string, args...interface{}) {
	if f.level > GLogLevelWarn {
		return
	}
	f.writeToChan(GLogLevelWarn, f.module, format, args...)
	/*
	logData := f.formatLogger(GLogLevelWarn, f.module, format, args...)
	f.writeLog(f.file, logData)
	 */
}

func (f *GFile)LogError(format string, args...interface{}) {
	if f.level > GLogLevelError {
		return
	}
	f.writeToChan(GLogLevelError, f.module, format, args...)
	/*
	logData := f.formatLogger(GLogLevelError, f.module, format, args...)
	f.writeLog(f.file, logData)
	 */
}

func (f *GFile)LogFatal(format string, args...interface{}) {
	if f.level > GLogLevelFatal {
		return
	}
	f.writeToChan(GLogLevelFatal, f.module, format, args...)
	/*
	logData := f.formatLogger(GLogLevelFatal, f.module, format, args...)
	f.writeLog(f.file, logData)
	 */
}

func (f *GFile)SetLevel(level int) {
    f.level = level
}

func (f *GFile)Close() {
	if f.logChan != nil {
		close(f.logChan)
	}
	f.wg.Wait()
	if f.file != nil {
		err := f.file.Sync()
		if err != nil {
		}
		err = f.file.Close()
		if err != nil {
		}
	}
}
