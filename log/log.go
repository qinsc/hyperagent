package log

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

var HALogger *logs.BeeLogger

func init() {
	fmt.Println("Init Log ...")
	HALogger = logs.NewLogger(2)
	HALogger.EnableFuncCallDepth(true)
	HALogger.SetLogFuncCallDepth(3)
	HALogger.SetLogger("console", ``)
	fmt.Println("Log init complete.")
}

func Info(format string, v ...interface{}) {
	HALogger.Info(format, v...)
}

func Debug(format string, v ...interface{}) {
	HALogger.Debug(format, v...)
}

func Warn(format string, v ...interface{}) {
	HALogger.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	HALogger.Error(format, v...)
}
