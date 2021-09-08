// @Title 日志封装
// @Description $
// @Author  55
// @Date  2021/9/7
package log

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
)

// 日志对象
var Log *logs.BeeLogger

// 显示tags对象
var Tags = make(map[string]bool)

// 默认初始化
func init() {
	Log = logs.NewLogger(200)

	// 日志输出调用的文件名和文件行号，默认为false
	Log.EnableFuncCallDepth(true)

	// 如果你的应用自己封装了调用 log 包,那么需要设置 SetLogFuncCallDepth
	// 默认是 2,也就是直接调用的层级,如果你封装了多层,那么需要根据自己的需求进行调整.
	Log.SetLogFuncCallDepth(3)
}

// 兼容接口
func Init(logFile string, tags []string) (err error) {
	setLogFile("file", logFile, 30, true)
	setTags(tags)

	return
}

// 初始化日志参数
func InitLog(logWay string, logFile string, logLevel string, maxDays int64, disableLogColor bool, tags []string) {
	setLogFile(logWay, logFile, maxDays, disableLogColor)
	setLogLevel(logLevel)
	setTags(tags)
}

// 设置显示tag
func setTags(tags []string)  {
	if tags != nil {
		for _, tag := range tags {
			Tags[tag] = true
		}
	}
}

// 设置日志文件参数
// logWay: file or console
func setLogFile(logWay string, logFile string, maxdays int64, disableLogColor bool) {
	if logWay == "console" {
		params := ""
		if disableLogColor {
			params = fmt.Sprintf(`{"color": false}`)
		}
		_ = Log.SetLogger("console", params)
	} else {
		params := fmt.Sprintf(`{"filename": "%s", "maxdays": %d}`, logFile, maxdays)
		_ = Log.SetLogger("file", params)
	}
}

// 设置日志显示等级
// value: error, warning, info, debug, trace
func setLogLevel(logLevel string) {
	level := 4 // warning
	switch logLevel {
	case "error":
		level = 3
	case "warn":
		level = 4
	case "info":
		level = 6
	case "debug":
		level = 7
	case "trace":
		level = 8
	default:
		level = 4
	}
	Log.SetLevel(level)
}

// 获取tag是否显示
func getTag(tag string) (msg string, show bool) {
	if _, ok := Tags[tag]; !ok {
		return
	}

	msg = "[Tag:"+tag+"] "
	show = true

	return
}

// wrap log
func Error(tag, format string, v ...interface{}) {
	tag, show := getTag(tag)
	if show {
		Log.Error(tag+format, v...)
	}
}

func Warn(tag,format string, v ...interface{}) {
	tag, show := getTag(tag)
	if show {
		Log.Warn(tag+format, v...)
	}
}

func Info(tag,format string, v ...interface{}) {
	tag, show := getTag(tag)
	if show {
		Log.Info(tag+format, v...)
	}
}

func Debug(tag,format string, v ...interface{}) {
	tag, show := getTag(tag)
	if show {
		Log.Debug(tag+format, v...)
	}
}

func Trace(tag,format string, v ...interface{}) {
	tag, show := getTag(tag)
	if show {
		Log.Trace(tag+format, v...)
	}
}
