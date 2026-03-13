package pkg

import (
	"fmt"
)

const (
	LogLevelErr   = 1
	LogLevelWarn  = 2
	LogLevelInfo  = 3
	LogLevelDebug = 4
)

var (
	gLogLevel = LogLevelInfo
)

func InitLogLevel(level int) {
	gLogLevel = level
}

func Error(format string, argv ...interface{}) {
	fmt.Printf(GetCurTimeStr()+" Error msg: "+format+"\n", argv...)
}

func Warn(format string, argv ...interface{}) {
	if gLogLevel <= LogLevelWarn {
		fmt.Printf(GetCurTimeStr()+"%s Warn msg: "+format+"\n", argv...)
	}
}

func Info(format string, argv ...interface{}) {
	if gLogLevel <= LogLevelInfo {
		fmt.Printf(GetCurTimeStr()+" Info msg: "+format+"\n", argv...)
	}
}

func Debug(format string, argv ...interface{}) {
	if gLogLevel <= LogLevelDebug {
		fmt.Printf(GetCurTimeStr()+" Debug msg: "+format+"\n", argv...)
	}
}
