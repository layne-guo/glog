package main

import (
	"fmt"
	"glog"
	"flag"
	"time"
)

func loggic() {
	for {
		glog.LogDebug("gjldfasfasgaslfalfabrin")
		glog.LogTrace("gjldfasfasgaslfalfabrin")
		glog.LogInfo("gjldfasfasgaslfalfabrin")
		glog.LogWarn("gjldfasfasgaslfalfabrin")
		glog.LogError("gjldfasfasgaslfalfabrin")
		glog.LogFatal("gjldfasfasgaslfalfabrin")
		time.Sleep(time.Second*3)
	}
}

/*
func testGetLine() {
	fileName, funcName, lineNo := glog.GetLineInfo(2)
	fmt.Printf("fileName:%s funcName:%s lineNo:%d\n", fileName, funcName, lineNo)
}
*/

func main() {
	var logTypeStr string
	flag.StringVar(&logTypeStr, "type", "console", "please input logger type")
	flag.Parse()
	var logType int
	if (logTypeStr == "file") {
		logType = glog.GLogTypeFile
	} else {
		logType = glog.GLogTypeConsole
	}
	err := glog.Init(logType, glog.GLogLevelDebug, "./glog.log", "glog_exp")
	//logger := glog.NewGLog(logType, glog.GLogLevelDebug, "./glog.log", "glog_exp")
	//logger := glog.NewGLog(logType, glog.GLogLevelDebug, "", "glog_exp")
	if err != nil {
		fmt.Printf("logger init failed\n")
	}
	loggic()
	glog.SetLevel(glog.GLogLevelDebug)
	glog.Close()
}