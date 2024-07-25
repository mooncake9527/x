# 日志记录器

## 用法

```go
logger := xlog.NewStdLogger(os.Stdout)
// fields & valuer
logger = xlog.With(logger,
"service.name", "hellworld",
"service.version", "v1.0.0",
"ts", xlog.DefaultTimestamp,
"caller", xlog.DefaultCaller,
)
logger.Log(xlog.LevelInfo, "key", "value")

// helper
helper := xlog.NewHelper(logger)
helper.Log(xlog.LevelInfo, "key", "value")
helper.Info("info message")
helper.Infof("info %s", "message")
helper.Infow("key", "value")

// filter
log := xlog.NewHelper(xlog.NewFilter(logger,
log.FilterLevel(xlog.LevelInfo),
log.FilterKey("foo"),
log.FilterValue("bar"),
log.FilterFunc(customFilter),
))
log.Debug("debug log")
log.Info("info log")
log.Warn("warn log")
log.Error("warn log")
```

## 第三方日志库

### zap

```shell
go get -u github.com/go-kratos/kratos/contrib/log/zap/v2
```
### logrus

```shell
go get -u github.com/go-kratos/kratos/contrib/log/logrus/v2
```

### fluent

```shell
go get -u github.com/go-kratos/kratos/contrib/log/fluent/v2
```

### aliyun

```shell
go get -u github.com/go-kratos/kratos/contrib/log/aliyun/v2
```
