package console

import "log"

var DebugLogger *log.Logger

func DebugLog(v ...interface{}) {
	DebugLogger.Println(v)
}

func DebugLogf(format string, v ...interface{}) {
	DebugLogger.Printf(format, v)
}
