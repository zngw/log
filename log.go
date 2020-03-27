/**
* File:log.go
* Copyright: Copyright (c) 2019
* Created on 2019-11-11
* Author:zengwu
* Version 1.0
* Title: 日志类
 */
package log

import (
	"fmt"
	"github.com/cihub/seelog"
	"runtime"
	"strings"
)

// 输入日志的标签
var showTag map[string]bool
var logFolder string

// 初始化日志
// file 日志文件名 ""为logs/log.info.2016-02-01;tags 显示的tag, nil为不设置显示tags
func Init(file string, tags []string) (err error) {
	// 配置日志文件，运行文件所在目录/logs/文件名
	if len(file) > 0 {
		fileConfig = strings.Replace(fileConfig, "./logs/log", file, -1)
	}

	logger, _ := seelog.LoggerFromConfigAsString(fileConfig)
	err = seelog.ReplaceLogger(logger)
	if err != nil {
		fmt.Println("log init error.", err)
		return 
	}

	// 根据配置显示tag
	if tags!=nil {
		showTag = make(map[string]bool)
		for _, tag := range tags {
			showTag[tag] = true
		}
	}

	return
}

// 重新设置显示标志
func ReplaceShowTag(tags []string) {
	if tags==nil {
		showTag = nil
		return
	}

	showTag = make(map[string]bool)
	for _, tag := range tags {
		showTag[tag] = true
	}
}

// 输出日志
func Trace(tag string, v ...interface{}) {
	if showTag != nil {
		if _, ok := showTag[tag]; !ok {
			return
		}
	}

	seelog.Trace("- [Tag:"+tag+"] ", v)
}

// 输出错误日志
func Error(v ...interface{}) {
	pc, filename, line, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	caller := fmt.Sprintf("- [%s:%d - %s] ", filename, line, funcName)
	_ = seelog.Error(caller, v)
}

var fileConfig = `
<!-- type 设置记录器类型 https://github.com/cihub/seelog/wiki/Logger-types-referenceminlevel 设置日志最低级别; 
maxlevel 设置日志的最高级别也可以通过 <seelog levels="trace,info,critical"> 设置日记级别 -->
<seelog type="asynctimer" asyncinterval="5000000" minlevel="trace" maxlevel="error">
	<!-- <outputs> formatid 指定日志输出的格式(格式在<formats>标签中定义) -->
    <outputs formatid="trace">
		<filter levels="trace,info,debug">
			<!-- <console> 标签表示输出到终端 -->
			<console formatid="colored" />
			<!-- <rollingfile>滚动文件(定期清除过期日志) formatid: 指定日志格式; type="size/date" 按大小/按日期; maxsize: 单日志文件最大大小; maxrools: 最大文件数 -->
            <rollingfile formatid="trace" type="date" filename="./logs/log.info" datepattern="2006-01-02" fullname="true" maxrolls="5"/>
        </filter>
        <filter levels="error,warn">
			<console formatid="coloredErr" />
            <rollingfile formatid="trace" type="date" filename="./logs/log.err" datepattern="2006-01-02" fullname="true" maxrolls="30"/>
        </filter>
    </outputs>
   
    <formats>
 		<!-- <formats> 定制日志的输出格式https://github.com/cihub/seelog/wiki/Format-reference -->
        <format id="trace" format="%Date(2006-01-02 15:04:05.000) [%LEVEL] %Msg%n" />
        <format id="err" format="%Date(2006-01-02 15:04:05.000) [%LEVEL] [%File:%Line-%Func] %Msg%n" />
        <format id="colored" format="%Date(2006-01-02 15:04:05.000) %EscM(42)[%LEVEL]%EscM(49) %Msg%n%EscM(0)" />
		<format id="coloredErr" format="%Date(2006-01-02 15:04:05.000) %EscM(41)[%LEVEL]%EscM(49) [%File:%Line-%Func] %Msg%n%EscM(0)"/>
    </formats>
</seelog>
`
