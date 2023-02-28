# 直接写入ElasticSearch 实验

```go
    l := func NewEsLogger(level, user, password, index string, url []string)
    l.Debug("hello:%v", "nnn")
```

## 耗时38.425µs

直接打开一个协程不考虑是否写入成功至es
