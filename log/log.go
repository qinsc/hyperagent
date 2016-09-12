package log

import (
	"github.com/astaxie/beego/logs"
)

var (
	bglog *logs.BeeLogger
)

func init() {
	bglog = logs.NewLogger(2)
	bglog.EnableFuncCallDepth(true)
	bglog.SetLogFuncCallDepth(3)
	//	bglog.SetLogger("console", ``)
	bglog.SetLogger("file", `{"filename":"hyperagent.log"}`)
}

func Info(format string, v ...interface{}) {
	bglog.Info(format, v...)
}

func Debug(format string, v ...interface{}) {
	bglog.Debug(format, v...)
}

func Warn(format string, v ...interface{}) {
	bglog.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	bglog.Error(format, v...)
}
