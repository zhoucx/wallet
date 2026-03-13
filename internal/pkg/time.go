package pkg

import (
	"time"
)

func GetCurTime() time.Time {
	return time.Now()
}

func GetCurTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
