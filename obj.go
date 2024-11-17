// @Title 以对象格式创建日志
// @Description
// @Author  55
// @Date  2021/12/5
package log

import (
	"github.com/zngw/golib/log"
	"strings"
)

type Logger struct {
	name string      // 输出对象名
	log  *log.Logger // 日志对象
}

// 默认初始化
func New(name string) (logger *Logger) {
	logger = &Logger{
		log: log.New(log.Option{
			LogPath:    "logs/file.log",
			MaxDays:    30,
			CallerSkip: 1,
		}),
	}

	logger.name = name
	return
}

// 初始化日志参数
// logWay: file-输出到文件；console-输出到控制台
// logFile: 日志文件
// logLevel: 日志等级
// maxDays: 日志保留天数
// disableLogColor: 是否显示颜色
// tags: 日志显示tag
func (logger *Logger) Init(logWay string, logFile string, logLevel string, maxDays int64, disableLogColor bool, tags []string) {
	if logWay == "console" {
		logFile = "console"
	}

	logger.log.WithOptions(log.Option{
		LogPath:         logFile,
		LogLevel:        logLevel,
		Tags:            strings.Join(tags, ","),
		MaxDays:         int(maxDays),
		DisableLogColor: disableLogColor,
		CallerSkip:      1,
	})
}

// wrap log
func (logger *Logger) Error(tag, format string, v ...interface{}) {
	v = append([]interface{}{format}, v...)
	logger.log.Error(tag, v...)
}

func (logger *Logger) Warn(tag, format string, v ...interface{}) {
	v = append([]interface{}{format}, v...)
	logger.log.Warn(tag, v...)
}

func (logger *Logger) Info(tag, format string, v ...interface{}) {
	v = append([]interface{}{format}, v...)
	logger.log.Info(tag, v...)
}

func (logger *Logger) Debug(tag, format string, v ...interface{}) {
	v = append([]interface{}{format}, v...)
	logger.log.Debug(tag, v...)
}

func (logger *Logger) Trace(tag, format string, v ...interface{}) {
	v = append([]interface{}{format}, v...)
	logger.log.Trace(tag, v...)
}
