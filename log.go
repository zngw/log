// @Title 日志封装
// @Description $
// @Author  55
// @Date  2021/9/7
package log

// 日志对象
var logger *Logger

// 获取默认日志对象
func Log() (*Logger){
	if logger == nil {
		logger = New("")
	}

	return logger
}

// 简易初始化，兼容之前版本
// logFile: 日志文件
// tags: 日志显示tag
func Init(logFile string, tags []string) (err error) {
	Log().cleanLogWay()

	if len(logFile) > 0 {
		Log().setLogFile(logFile, 30)
	} else {
		Log().setLogConsole(true)
	}

	Log().setTags(tags)

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
	Log().cleanLogWay()

	if logWay == "all" {
		Log().setLogFile(logFile, maxDays)
		Log().setLogConsole(disableLogColor)
	} else if logWay == "file" {
		Log().setLogFile(logFile, maxDays)
	} else if logWay == "console" {
		Log().setLogConsole(disableLogColor)
	}

	Log().setLogLevel(logLevel)
	Log().setTags(tags)
}

// wrap log
func Error(tag, format string, v ...interface{}) {
	Log().Error(tag,format,v...)
}

func Warn(tag, format string, v ...interface{}) {
	Log().Warn(tag,format,v...)
}

func Info(tag, format string, v ...interface{}) {
	Log().Info(tag,format,v...)
}

func Debug(tag, format string, v ...interface{}) {
	Log().Debug(tag,format,v...)
}

func Trace(tag, format string, v ...interface{}) {
	Log().Trace(tag,format,v...)
}
