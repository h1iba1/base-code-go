## logus库学习
github.com/sirupsen/logrus

[Go 每日一库之 logrus](https://juejin.cn/post/6844904061393698823)

### 日志级别

- Panic：记录日志，然后panic。
- Fatal：致命错误，出现错误时程序无法正常运转。输出日志后，程序退出；
- Error：错误日志，需要查看原因；
- Warn：警告信息，提醒程序员注意；
- Info：关键操作，核心流程的日志；
- Debug：一般程序中输出的调试信息；
- Trace：很细粒度的信息，一般用不到；


### 定制
#### 输出文件名
调用logrus.SetReportCaller(true)设置在输出日志中添加文件名和方法信息：
```go
func TestSetReportCaller(t *testing.T) {
	logrus.SetReportCaller(true)

	logrus.Info("info msg")
}
```
输出：
```txt
=== RUN   TestSetReportCaller
time="2021-11-28T15:42:57+08:00" level=info msg="info msg" func=baseCode/common-lib/logus.TestSetReportCaller file="/Users/xxx/Desktop/go/src/baseCode/common-lib/logus/logus_test.go:23"
--- PASS: TestSetReportCaller (0.00s)
```

输出多了两个字段file为调用logrus相关方法的文件名，method为方法名：

#### 添加字段
有时候需要在输出中添加一些字段，可以通过调用logrus.WithField和logrus.WithFields实现。
```go
func TestName(t *testing.T) {
	logrus.WithFields(logrus.Fields{
		"name": "dj",
		"age": 18,
	}).Info("info msg")
}
```

```txt
=== RUN   TestName
time="2021-11-28T15:48:17+08:00" level=info msg="info msg" age=18 name=dj
--- PASS: TestName (0.00s)
```

#### 重定向输出
默认情况下，日志输出到io.Stderr。可以调用logrus.SetOutput传入一个io.Writer参数。后续调用相关方法日志将写到io.Writer中。
现在，我们就能像上篇文章介绍log时一样，可以搞点事情了。传入一个io.MultiWriter，
同时将日志写到bytes.Buffer、标准输出和文件中：
```go

func TestSetOutput(t *testing.T) {
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	//同时将日志写到bytes.Buffer、标准输出和文件中：
	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logrus.Info("info msg")
}
```

#### 日志格式
logrus支持两种日志格式，文本和 JSON，默认为文本格式。可以通过logrus.SetFormatter设置日志格式：

```go
func TestJSONFormatter(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")
}
```

```txt
{"level":"trace","msg":"trace msg","time":"2021-11-28T15:58:34+08:00"}
{"level":"debug","msg":"debug msg","time":"2021-11-28T15:58:34+08:00"}
{"level":"info","msg":"info msg","time":"2021-11-28T15:58:34+08:00"}
{"level":"warning","msg":"warn msg","time":"2021-11-28T15:58:34+08:00"}
{"level":"error","msg":"error msg","time":"2021-11-28T15:58:34+08:00"}
{"level":"fatal","msg":"fatal msg","time":"2021-11-28T15:58:34+08:00"}
```

#### 第三方格式
除了内置的TextFormatter和JSONFormatter，还有不少第三方格式支持。我们这里介绍一个[logrus-prefixed-formatter](https://github.com/x-cray/logrus-prefixed-formatter)

```go
package log

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	Log *logrus.Entry
)

func init(){
	logger :=logrus.New()
	logger.Formatter =new(prefixed.TextFormatter)
	// 等看到debug以上的日志
	logger.Level =logrus.DebugLevel
	// 设置日志前缀
	Log =logger.WithFields(logrus.Fields{"prefix": "cybersky"})
}
```



















