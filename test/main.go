// @Title
// @Description $
// @Author  55
// @Date  2021/11/7
package main

import (
	"github.com/zngw/log"
)

func main() {
	// 日志有三种方式初始化。

	// 第一种
	// 不初始化，直接使用
	// 这时候会用到默认init中的参数初始化
	// 默认为终端、文件双重显示；日志文件为`logs/file.log`；保留30天；显示所有tags

	// 第二种
	// 简易初始化，兼容最初版本的日志接口，只设置日志文件和tags
	// 日志文件为空时为终端显示，
	// 设置文件时只输出到文件，日志文件保留30天
	//err := log.Init("logs/file.log", nil)
	//if err != nil {
	//	panic(err)
	//}

	// 第三种
	// 完整初始化
	// logWay: all-输出到文件和控制台;file-输出到文件；console-输出到控制台
	// logFile: 日志文件
	// logLevel: 日志等级
	// maxDays: 日志保留天数
	// disableLogColor: 是否显示颜色
	// tags: 日志显示tag
	log.InitLog("all", "logs/file.log", "info", 30, false, []string{"sys", "net"})

	// 输出： 2021/11/07 02:53:59.148 [I] [main.go:39]  [Tag:sys] Hello World
	log.Info("sys", "Hello World")

	// 输出： 2021/11/07 02:53:59.148 [I] [main.go:42]  [Tag:net] Hello Golang
	log.Info("net", "Hello Golang")

	// 这条日志不在显示的tag内，故不输出
	log.Info("test", "Hello zngw")

	// 输出错误日志: 2021/11/07 02:53:59.148 [E] [main.go:48]  [Tag:sys] Error
	log.Error("sys", "Error")

	// 以对象方式创建多个日志模块
	mylog1 := log.New("mylog1")
	mylog1.Init("all", "logs/mylog1.log", "info", 30, false, []string{"sys", "net"})

	mylog2 := log.New("mylog2")
	mylog2.Init("all", "logs/mylog2.log", "info", 30, false, []string{"sys", "net"})

	mylog1.Info("net", "MyLog1")
	mylog2.Info("sys", "MyLog2")
}
