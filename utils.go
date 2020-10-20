package glog

import "runtime"

func GetLineInfo(skip int) (filename string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		fun := runtime.FuncForPC(pc)
		funcName = fun.Name()
	}
	filename = file
	lineNo = line
	return
}
