[comment]: <> (dtapps)
[![GitHub Org's stars](https://img.shields.io/github/stars/zngw)](https://github.com/zngw)

[comment]: <> (go)
[![godoc](https://pkg.go.dev/badge/github.com/zngw/log?status.svg)](https://pkg.go.dev/github.com/zngw/log)
[![oproxy.cn](https://goproxy.cn/stats/github.com/zngw/log/badges/download-count.svg)](https://goproxy.cn/stats/github.com/zngw/log)
[![goreportcard.com](https://goreportcard.com/badge/github.com/zngw/log)](https://goreportcard.com/report/github.com/zngw/log)
[![deps.dev](https://img.shields.io/badge/deps-go-red.svg)](https://deps.dev/go/github.com%2Fdtapps%2Fgo-ssh-tunnel)

[comment]: <> (github.com)
[![watchers](https://badgen.net/github/watchers/zngw/log)](https://github.com/zngw/log/watchers)
[![stars](https://badgen.net/github/stars/zngw/log)](https://github.com/zngw/log/stargazers)
[![forks](https://badgen.net/github/forks/zngw/log)](https://github.com/zngw/log/network/members)
[![issues](https://badgen.net/github/issues/zngw/log)](https://github.com/zngw/log/issues)
[![branches](https://badgen.net/github/branches/zngw/log)](https://github.com/zngw/log/branches)
[![releases](https://badgen.net/github/releases/zngw/log)](https://github.com/zngw/log/releases)
[![tags](https://badgen.net/github/tags/zngw/log)](https://github.com/zngw/log/tags)
[![license](https://badgen.net/github/license/zngw/log)](https://github.com/zngw/log/blob/master/LICENSE)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/zngw/log)](https://github.com/zngw/log)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/zngw/log)](https://github.com/zngw/log/releases)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/zngw/log)](https://github.com/zngw/log/tags)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/zngw/log)](https://github.com/zngw/log/pulls)
[![GitHub issues](https://img.shields.io/github/issues/zngw/log)](https://github.com/zngw/log/issues)
[![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/zngw/log)](https://github.com/zngw/log)
[![GitHub language count](https://img.shields.io/github/languages/count/zngw/log)](https://github.com/zngw/log)
[![GitHub search hit counter](https://img.shields.io/github/search/zngw/log/go)](https://github.com/zngw/log)
[![GitHub top language](https://img.shields.io/github/languages/top/zngw/log)](https://github.com/zngw/log)

# 说明
日志模块，详细情可以参考项目[https://github.com/zngw/golib](https://github.com/zngw/golib)

# 安装

```bash
go get -u github.com/zngw/log
```

# 初始化

日志有三种方式初始化。

## 第一种

* 不初始化，直接使用
* 这时候会用到默认init中的参数初始化
* 默认为终端、文件双重显示；日志文件为`logs/file.log`；保留30天；显示所有tags

## 第二种
* 简易初始化，兼容最初版本的日志接口，只设置日志文件和tags
* 日志文件为空时为终端显示
* 设置文件时只输出到文件，日志文件保留30天

```go
    err := log.Init("logs/file.log", nil)
    if err != nil {
        panic(err)
    }
```

## 第三种

* 完整初始化
* logWay: all-输出到文件和控制台;file-输出到文件；console-输出到控制台
* logFile: 日志文件
* logLevel: 日志等级
* maxDays: 日志保留天数
* disableLogColor: 是否显示颜色
* tags: 日志显示tag

```go
    log.InitLog("all", "logs/file.log", "info", 30, false, []string{"sys", "net"})
```

# 以创建多个日志模块
可以使用`log.New("日志模块名")`方法创建多个日志模块，为不同的功能记录不同的日志文件和格式。
如可以将错误日志独立出来，也可以根据不同的逻辑写不同的日志等

```go
	errorlog := log.New("error")
	errorlog.Init("all", "logs/error.log", "error", 30, false, nil)
```
	
# Quick-start
```go
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

	// 创建多个日志模块
	// 错误日志
	errorlog := log.New("error")
	errorlog.Init("all", "logs/error.log", "error", 30, false, nil)

	// 独立日志模块
	mylog := log.New("mylog")
	mylog.Init("all", "logs/mylog.log", "info", 30, false, []string{"sys", "net"})

	errorlog.Error("net", "Error Log")
	mylog.Info("sys", "MyLog")
}
```
