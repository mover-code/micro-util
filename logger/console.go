/***************************
@File        : console.go
@Time        : 2020/12/03 15:54:20
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : logger
****************************/

package logger

import (
    "fmt"
    "time"
)

// 在终端输出日志

// Logger 日志结构体
type Logger struct {
    level Level
}

// 格式化输出
func (l Logger) printLog(s, msg string, a ...interface{}) {
    message := fmt.Sprintf(msg, a...)
    now := time.Now()
    funcName, fileName, line := getInfo(3)
    fmt.Printf("[%s] [%s] [文件:%s 函数:%s 行数:%d] --- %s \n", now.Format("2006-01-02 15:04:05"), s, fileName, funcName, line, message)
}

// 比较级别函数
func (l Logger) enable(level Level) bool {
    return l.level <= level
}

// Debug 方法
func (l Logger) Debug(msg string, a ...interface{}) {
    if l.enable(DEBUG) {
        l.printLog("DEBUG", msg, a...)
    }
}

// Info 方法
func (l Logger) Info(msg string, a ...interface{}) {
    if l.enable(INFO) {
        l.printLog("INFO", msg, a...)
    }
}

// Waring 方法
func (l Logger) Waring(msg string, a ...interface{}) {
    if l.enable(WARING) {
        l.printLog("WARINING", msg, a...)
    }

}

// Error 方法
func (l Logger) Error(msg string, a ...interface{}) {
    if l.enable(ERROR) {
        l.printLog("ERROR", msg, a...)
    }
}

// Fatal 方法
func (l Logger) Fatal(msg string, a ...interface{}) {
    if l.enable(FATAL) {
        l.printLog("FATAL", msg, a...)
    }
}
