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
var logger *logs.BeeLogger

// 显示tags对象
var tags = make(map[string]bool)

// 默认初始化, 调用log文件时自动调用
func init() {
	logger = logs.NewLogger(200)

	// 日志输出调用的文件名和文件行号，默认为false
	logger.EnableFuncCallDepth(true)

	// 如果你的应用自己封装了调用 log 包,那么需要设置 SetLogFuncCallDepth
	// 默认是 2,也就是直接调用的层级,如果你封装了多层,那么需要根据自己的需求进行调整.
	logger.SetLogFuncCallDepth(3)

	// 默认初始化输入文件参数
	// 用于不调用Init、InitLog初始化时可直接调用log
	setLogFile("logs/file.log", 30)
	setLogConsole(false)
}

// 简易初始化，兼容之前版本
// logFile: 日志文件
// tags: 日志显示tag
func Init(logFile string, tags []string) (err error) {
	cleanLogWay()

	if len(logFile) > 0 {
		setLogFile(logFile, 30)
	} else {
		setLogConsole(true)
	}

	setTags(tags)

	return
}

// 初始化日志参数
// logWay: all-输出到文件和控制台;file-输出到文件；console-输出到控制台
// logFile: 日志文件
// logLevel: 日志等级
// maxDays: 日志保留天数
// disableLogColor: 是否显示颜色
// tags: 日志显示tag
func InitLog(logWay string, logFile string, logLevel string, maxDays int64, disableLogColor bool, tags []string) {
	cleanLogWay()

	if logWay == "all" {
		setLogFile(logFile, maxDays)
		setLogConsole(disableLogColor)
	} else if logWay == "file" {
		setLogFile(logFile, maxDays)
	} else if logWay == "console" {
		setLogConsole(disableLogColor)
	}

	setLogLevel(logLevel)
	setTags(tags)
}

// 设置显示tag
func setTags(ts []string) {
	tags = make(map[string]bool)
	if tags != nil {
		for _, tag := range ts {
			tags[tag] = true
		}
	}
}

// 清理输出方式
func cleanLogWay()  {
	_ = logger.DelLogger("file")
	_ = logger.DelLogger("console")
}

// 设置输入文件参数
// logFile: 输出文件
// maxdays：日志文件保留天数
func setLogFile(logFile string, maxdays int64) {
	params := fmt.Sprintf(`{"filename": "%s", "maxdays": %d}`, logFile, maxdays)
	_ = logger.SetLogger("file", params)
}

// 设置日志终端参数
// 禁用日志文字颜色
func setLogConsole(disableLogColor bool) {
	params := ""
	if disableLogColor {
		params = fmt.Sprintf(`{"color": false}`)
	}
	_ = logger.SetLogger("console", params)
}

// 设置日志显示等级
// value: error, warning, info, debug, trace
func setLogLevel(logLevel string) {
	level := logs.LevelWarn // warning
	switch logLevel {
	case "error":
		level = logs.LevelError // 3
	case "warn":
		level = logs.LevelWarn // 4
	case "info":
		level = logs.LevelInfo // 6
	case "debug":
		level = logs.LevelDebug // 7
	case "trace":
		level = logs.LevelTrace // 7
	default:
		level = logs.LevelWarn // 4
	}

	logger.SetLevel(level)
}

// 获取tag是否显示
func getTag(tag string) (msg string, show bool) {
	if len(tags) > 0 {
		if _, ok := tags[tag]; !ok {
			return
		}
	}

	msg = "[Tag:" + tag + "] "
	show = true

	return
}

// wrap log
func Error(tag, format string, v ...interface{}) {
	tag, show := getTag(tag)
	if show {
		logger.Error(tag+format, v...)
	}
}

func Warn(tag, format string, v ...interface{}) {
	tag, show := getTag(tag)
	if show {
		logger.Warn(tag+format, v...)
	}
}

func Info(tag, format string, v ...interface{}) {
	tag, show := getTag(tag)
	if show {
		logger.Info(tag+format, v...)
	}
}

func Debug(tag, format string, v ...interface{}) {
	tag, show := getTag(tag)
	if show {
		logger.Debug(tag+format, v...)
	}
}

func Trace(tag, format string, v ...interface{}) {
	tag, show := getTag(tag)
	if show {
		logger.Trace(tag+format, v...)
	}
}
