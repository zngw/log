// @Title 以对象格式创建日志
// @Description
// @Author  55
// @Date  2021/12/5
package log

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
)

type Logger struct {
	name string          // 输出对象名
	log  *logs.BeeLogger // 日志对象
	tags map[string]bool // 显示tags对象
}

// 默认初始化
func New(name string) (logger *Logger) {
	logger = &Logger{
		log:  logs.NewLogger(200),
		tags: make(map[string]bool),
	}

	logger.name = name

	// 日志输出调用的文件名和文件行号，默认为false
	logger.log.EnableFuncCallDepth(true)

	// 如果你的应用自己封装了调用 log 包,那么需要设置 SetLogFuncCallDepth
	// 默认是 2,也就是直接调用的层级,如果你封装了多层,那么需要根据自己的需求进行调整.
	logger.log.SetLogFuncCallDepth(3)

	// 默认初始化输入文件参数
	// 用于不调用Init、InitLog初始化时可直接调用log
	logger.setLogFile("logs/file.log", 30)
	logger.setLogConsole(false)

	return
}

// 初始化日志参数
// logWay: all-输出到文件和控制台;file-输出到文件；console-输出到控制台
// logFile: 日志文件
// logLevel: 日志等级
// maxDays: 日志保留天数
// disableLogColor: 是否显示颜色
// tags: 日志显示tag
func (logger *Logger) Init(logWay string, logFile string, logLevel string, maxDays int64, disableLogColor bool, tags []string) {
	logger.cleanLogWay()

	if logWay == "all" {
		logger.setLogFile(logFile, maxDays)
		logger.setLogConsole(disableLogColor)
	} else if logWay == "file" {
		logger.setLogFile(logFile, maxDays)
	} else if logWay == "console" {
		logger.setLogConsole(disableLogColor)
	}

	logger.setLogLevel(logLevel)
	logger.setTags(tags)
}

// 设置显示tag
func (logger *Logger) setTags(ts []string) {
	logger.tags = make(map[string]bool)
	if logger.tags != nil {
		for _, tag := range ts {
			logger.tags[tag] = true
		}
	}
}

// 清理输出方式
func (logger *Logger) cleanLogWay() {
	_ = logger.log.DelLogger("file")
	_ = logger.log.DelLogger("console")
}

// 设置输入文件参数
// logFile: 输出文件
// maxdays：日志文件保留天数
func (logger *Logger) setLogFile(logFile string, maxdays int64) {
	params := fmt.Sprintf(`{"filename": "%s", "maxdays": %d}`, logFile, maxdays)
	_ = logger.log.SetLogger("file", params)
}

// 设置日志终端参数
// 禁用日志文字颜色
func (logger *Logger) setLogConsole(disableLogColor bool) {
	params := ""
	if disableLogColor {
		params = fmt.Sprintf(`{"color": false}`)
	}
	_ = logger.log.SetLogger("console", params)
}

// 设置日志显示等级
// value: error, warning, info, debug, trace
func (logger *Logger) setLogLevel(logLevel string) {
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

	logger.log.SetLevel(level)
}

// 获取tag是否显示
func (logger *Logger) getTag(tag string) (msg string, show bool) {
	if len(logger.tags) > 0 {
		if _, ok := logger.tags[tag]; !ok {
			return
		}
	}

	if len(logger.name) > 0 {
		msg = "[" + logger.name + "] [Tag:" + tag + "] "
	} else {
		msg = "[Tag:" + tag + "] "
	}

	show = true

	return
}

// wrap log
func (logger *Logger) Error(tag, format string, v ...interface{}) {
	tag, show := logger.getTag(tag)
	if show {
		logger.log.Error(tag+format, v...)
	}
}

func (logger *Logger) Warn(tag, format string, v ...interface{}) {
	tag, show := logger.getTag(tag)
	if show {
		logger.log.Warn(tag+format, v...)
	}
}

func (logger *Logger) Info(tag, format string, v ...interface{}) {
	tag, show := logger.getTag(tag)
	if show {
		logger.log.Info(tag+format, v...)
	}
}

func (logger *Logger) Debug(tag, format string, v ...interface{}) {
	tag, show := logger.getTag(tag)
	if show {
		logger.log.Debug(tag+format, v...)
	}
}

func (logger *Logger) Trace(tag, format string, v ...interface{}) {
	tag, show := logger.getTag(tag)
	if show {
		logger.log.Trace(tag+format, v...)
	}
}
