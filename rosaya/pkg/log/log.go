package log

import (
	"fmt"
	"time"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func Infof(format string, args ...interface{}) {
	InfoMsg := fmt.Sprintf(format, args...)
	InfoMsg = fmt.Sprintf("[INFO] [%s] %s", time.Now().Format("2006-01-02 15:04:05"), InfoMsg)
	InfoMsg = fmt.Sprintf(InfoColor, InfoMsg)
	fmt.Println(InfoMsg)
}

func Info(msg string) {
	InfoMsg := fmt.Sprintf("[INFO] [%s] %s", time.Now().Format("2006-01-02 15:04:05"), msg)
	InfoMsg = fmt.Sprintf(InfoColor, InfoMsg)
	fmt.Println(InfoMsg)
}

func Error(msg string) {
	ErrorMsg := fmt.Sprintf("[ERROR] [%s] %s", time.Now().Format("2006-01-02 15:04:05"), msg)
	ErrorMsg = fmt.Sprintf(ErrorColor, ErrorMsg)
	fmt.Println(ErrorMsg)
}

func Errorf(format string, args ...interface{}) {
	ErrorMsg := fmt.Sprintf(format, args...)
	ErrorMsg = fmt.Sprintf("[ERROR] [%s] %s", time.Now().Format("2006-01-02 15:04:05"), ErrorMsg)
	ErrorMsg = fmt.Sprintf(ErrorColor, ErrorMsg)
	fmt.Println(ErrorMsg)
}
