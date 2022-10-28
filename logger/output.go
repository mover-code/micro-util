/***************************
@File        : output_file.go
@Time        : 2020/12/04 12:26:49
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 向文件中输出日志
****************************/

package logger

import (
    "fmt"
    "os"
    "path"
    "time"
)

type FileLogger struct {
    level       Level
    filePath    string //日志保存路径
    fileName    string //日志文件名
    maxFileSize int64
    fileObj     *os.File
    errFileObj  *os.File
}

func (f *FileLogger) initFile() error {
    fullFileName := path.Join(f.filePath, f.fileName)
    fileobj, err := os.OpenFile(fullFileName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(fmt.Sprintf("open file faild, err:%v\n", err))
        return err
    }
    errFileobj, err := os.OpenFile(fullFileName+"_err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(fmt.Sprintf("open file faild, err:%v\n", err))
        return err
    }
    f.fileObj = fileobj
    f.errFileObj = errFileobj
    return nil
}

// 比较级别函数
func (l *FileLogger) enable(level Level) bool {
    return l.level <= level
}

// 切割文件
func (l *FileLogger) checkSize(file *os.File) (a *os.File) {
    fileInfo, _ := file.Stat()
    // fmt.Printf("文件%s - 大小:%d\n", fileInfo.Name(), fileInfo.Size())
    if fileInfo.Size() >= l.maxFileSize {
        nowStr := time.Now().Format("20060102150405000")
        oldName := path.Join(l.filePath, fileInfo.Name())
        newName := fmt.Sprintf("%s_bak_%s.log", oldName[:len(oldName)-4], nowStr)
        l.Close(file)
        os.Rename(oldName, newName)
        fileobj, err := os.OpenFile(oldName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Printf("open new log file fialed,err:%s", err)
        }
        return fileobj
    }
    return file
}

// 关闭文件
func (f *FileLogger) Close(file *os.File) {
    file.Close()
}

// 格式化输出
func (l *FileLogger) printLog(s, msg string, a ...interface{}) {
    message := fmt.Sprintf(msg, a...)
    now := time.Now()
    funcName, fileName, line := getInfo(3)
    l.fileObj = l.checkSize(l.fileObj)
    l.errFileObj = l.checkSize(l.errFileObj)
    if s == "ERROR" || s == "FATAL" {
        // 如果要记录日志大于等于error
        fmt.Fprintf(l.errFileObj, "[%s] [%s] [文件:%s 函数:%s 行数:%d] --- %s \n", now.Format("2006-01-02 15:04:05"), s, fileName, funcName, line, message)
    } else {
        fmt.Fprintf(l.fileObj, "[%s] [%s] [文件:%s 函数:%s 行数:%d] --- %s \n", now.Format("2006-01-02 15:04:05"), s, fileName, funcName, line, message)
    }

}

// Debug 方法
func (l *FileLogger) Debug(msg string, a ...interface{}) {
    if l.enable(DEBUG) {
        l.printLog("DEBUG", msg, a...)
    }
}

// Info 方法
func (l *FileLogger) Info(msg string, a ...interface{}) {
    if l.enable(INFO) {
        l.printLog("INFO", msg, a...)
    }
}

// Waring 方法
func (l *FileLogger) Waring(msg string, a ...interface{}) {
    if l.enable(WARING) {
        l.printLog("WARINING", msg, a...)
    }

}

// Error 方法
func (l *FileLogger) Error(msg string, a ...interface{}) {
    if l.enable(ERROR) {
        l.printLog("ERROR", msg, a...)
    }
}

// Fatal 方法
func (l FileLogger) Fatal(msg string, a ...interface{}) {
    if l.enable(FATAL) {
        l.printLog("FATAL", msg, a...)
    }
}
