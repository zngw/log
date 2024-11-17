// @Title 日志封装
// @Description $
// @Author  55
// @Date  2021/9/7
package log

// 日志对象
var logger *Logger

func init() {
	logger = New("")
}

// 简易初始化，兼容之前版本
// logFile: 日志文件
// tags: 日志显示tag
func Init(logFile string, tags []string) (err error) {
	logWay := "file"
	if len(logFile) == 0 {
		logFile = "console"
		logWay = "console"
	}
	logger.Init(logWay, logFile, "trace", 30, false, tags)

	return
}

// 初始化日志参数
// logWay: file-输出到文件；console-输出到控制台
// logFile: 日志文件
// logLevel: 日志等级
// maxDays: 日志保留天数
// disableLogColor: 是否显示颜色
// tags: 日志显示tag
func InitLog(logWay string, logFile string, logLevel string, maxDays int64, disableLogColor bool, tags []string) {
	logger.Init(logWay, logFile, logLevel, maxDays, disableLogColor, tags)
}

// wrap log
func Error(tag, format string, v ...interface{}) {
	logger.Error(tag, format, v...)
}

func Warn(tag, format string, v ...interface{}) {
	logger.Warn(tag, format, v...)
}

func Info(tag, format string, v ...interface{}) {
	logger.Info(tag, format, v...)
}

func Debug(tag, format string, v ...interface{}) {
	logger.Debug(tag, format, v...)
}

func Trace(tag, format string, v ...interface{}) {
	logger.Trace(tag, format, v...)
}
