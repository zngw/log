// @Title 日志封装
// @Description $
// @Author  55
// @Date  2021/9/7
package log

// 日志对象
var logger *Logger

func init()  {
	logger = New("")
}

// 简易初始化，兼容之前版本
// logFile: 日志文件
// tags: 日志显示tag
func Init(logFile string, tags []string) (err error) {
	logger.cleanLogWay()

	if len(logFile) > 0 {
		logger.setLogFile(logFile, 30)
	} else {
		logger.setLogConsole(true)
	}

	logger.setTags(tags)

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

// wrap log
func Error(tag, format string, v ...interface{}) {
	logger.Error(tag,format,v...)
}

func Warn(tag, format string, v ...interface{}) {
	logger.Warn(tag,format,v...)
}

func Info(tag, format string, v ...interface{}) {
	logger.Info(tag,format,v...)
}

func Debug(tag, format string, v ...interface{}) {
	logger.Debug(tag,format,v...)
}

func Trace(tag, format string, v ...interface{}) {
	logger.Trace(tag,format,v...)
}
