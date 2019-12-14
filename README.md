# log
封装seelog包

# Quick-start
```go
package main

import (
	"github.com/zngw/log"
)

func main() {
	// 不设置显示tag
	// log.Init("",nil)
	
	// 初始化日志
	dir, file := filepath.Split(os.Args[0])
	err := log.Init(dir+"/logs/"+file,[]string{"sys","net"})
	if err != nil {
		panic(err)
	}

	// 输出日志: 2019-11-15 01:06:01.215 [TRACE] - [Tag:sys] [Hello World]
	log.Trace("sys","Hello World")

	// 输出日志: 2019-11-15 01:06:01.215 [TRACE] - [Tag:net] [Hello Golang]
	log.Trace("net","Hello Golang")

	// 这条日志不在显示的tag内，故不输出
	log.Trace("test","Hello zngw")

	// 输出错误日志: 2019-11-15 01:06:01.215 [ERROR] [log.go:71-github.com/zngw/log.Error] - [F:/work/src/github.com/main.go:15 - main.main] [Error]
	log.Error("Error")
}

```