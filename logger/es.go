/*
 * @Author: small_ant xms.chnb@gmail.com
 * @Time: 2023-02-28 11:34:17
 * @LastAuthor: small_ant xms.chnb@gmail.com
 * @lastTime: 2023-02-28 14:18:45
 * @FileName: es
 * @Desc: logs to elasticSearch
 *
 * Copyright (c) 2023 by small_ant, All Rights Reserved.
 */

package logger

import (
    "context"
    "fmt"
    "time"

    "github.com/olivere/elastic/v7"
)

type EsLogger struct {
    Client   *elastic.Client
    level    Level
    Uri      []string
    Index    string
    UserName string
    Password string
    ctx      context.Context
}

type Info struct {
    FuncName string                `json:"funcName"`
    FileName string                `json:"fileName"`
    Line     int64                 `json:"line"`
    Message  string                `json:"message"`            // 微博内容
    Created  time.Time             `json:"created,omitempty"`  // 创建时间
    Location string                `json:"location,omitempty"` //位置
    Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

const mapping = `{"mappings":{"properties":{"funcName":{"type":"keyword"},"message":{"type":"text"},"fileName":{"type":"keyword"},"created":{"type":"date"},"line":{"type":"keyword"},"location":{"type":"geo_point"},"suggest_field":{"type":"completion"}}}}`

func (e *EsLogger) Init() {
    cli, err := elastic.NewClient(
        elastic.SetURL(e.Uri...), elastic.SetSniff(false),
        elastic.SetBasicAuth(e.Index, e.Password))
    if err != nil {
        panic(fmt.Sprintf("连接失败: %v\n", err))
    } else {
        e.Client = cli
        ctx := context.Background()
        e.ctx = ctx
        exists, err := e.Client.IndexExists(e.Index).Do(ctx)
        if err != nil {
            panic(fmt.Sprintf("索引查询失败: %v\n", err))
        } else {
            if !exists {
                _, err := e.Client.CreateIndex(e.Index).BodyString(mapping).Do(ctx)
                if err != nil {
                    panic(err)
                }
            }
        }
    }
}

func (l *EsLogger) printLog(s, msg string, a ...interface{}) {
    message := fmt.Sprintf(msg, a...)
    now := time.Now()
    funcName, fileName, line := getInfo(3)
    data := &Info{
        FuncName: funcName,
        FileName: fileName,
        Line:     int64(line),
        Created:  now,
        Message:  message,
    }
    go func() {
        // put1, err :=
        l.Client.Index().
            Index(l.Index). // 设置索引名称
            BodyJson(data). // 指定前面声明的微博内容
            Do(l.ctx)       // 执行请求，需要传入一个上下文对象          Id("1").        // 设置文档id
    }()
}

// 比较级别函数
func (l *EsLogger) enable(level Level) bool {
    return l.level <= level
}

// Debug 方法
func (l *EsLogger) Debug(msg string, a ...interface{}) {
    if l.enable(DEBUG) {
        l.printLog("DEBUG", msg, a...)
    }
}

// Info 方法
func (l *EsLogger) Info(msg string, a ...interface{}) {
    if l.enable(INFO) {
        l.printLog("INFO", msg, a...)
    }
}

// Waring 方法
func (l *EsLogger) Waring(msg string, a ...interface{}) {
    if l.enable(WARING) {
        l.printLog("WARINING", msg, a...)
    }

}

// Error 方法
func (l *EsLogger) Error(msg string, a ...interface{}) {
    if l.enable(ERROR) {
        l.printLog("ERROR", msg, a...)
    }
}

// Fatal 方法
func (l *EsLogger) Fatal(msg string, a ...interface{}) {
    if l.enable(FATAL) {
        l.printLog("FATAL", msg, a...)
    }
}
