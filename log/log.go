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
	HALogger.SetLogger("console", `{"level":0}`)
	fmt.Println("Log init complete.")
}

func Info(msg string) {
	//	HALogger.Info(msg)
	fmt.Println(msg)
}

func Debug(msg string) {
	//	HALogger.Debug(msg)
	fmt.Println(msg)
}

func Warn(msg string) {
	//	HALogger.Warn(msg)
	fmt.Println(msg)
}

func Error(msg string) {
	//	HALogger.Error(msg)
	fmt.Println(msg)
}
