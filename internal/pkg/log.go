package pkg

import (
	"fmt"
)

const (
	LogLevelErr = 1
	LogLevelWarn = 2
	LogLevelInfo = 3
	LogLevelDebug = 4
)

var (
	gLogLevel = LogLevelInfo
)

func InitLogLevel(level int) {
	gLogLevel = level
}


func Error(format string, argv ...interface{}) {
	fmt.Printf("%s Error msg: " + format+"\n", GetCurTimeStr(), argv)
}


func Warn(format string, argv ...interface{}) {
	if gLogLevel <= LogLevelWarn {
		fmt.Printf("%s Warn msg: " + format+"\n", GetCurTimeStr(), argv)
	}
}

func Info(format string, argv ...interface{}) {
	if gLogLevel <= LogLevelInfo {
	    fmt.Printf("%s Info msg: " + format+"\n", GetCurTimeStr(), argv)
	}
}

func Debug(format string, argv ...interface{}) {
	if gLogLevel <= LogLevelDebug {
	    fmt.Printf("%s Debug msg: " + format+"\n", GetCurTimeStr(), argv)
	}
}




