net/htpp
# server端
## 快速注册一个web服务
```go
func main() {
	http.HandleFunc("/Hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hi xiaomi's</h1> "))
	})
	
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println("http server error:", err)
	}
}
```

### http.HandleFunc

函数功能：

将handlerFunc函数绑定到pattern路径

实现流程：

http.HandleFunc：
```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
```
DefaultServeMux.HandleFunc:
```go
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	if handler == nil {
		panic("http: nil handler")
	}
	mux.Handle(pattern, HandlerFunc(handler))
}
```

HandlerFunc:
HandlerFunc为实现了Handler接口的方法
```go
type HandlerFunc func(ResponseWriter, *Request)
```

Handler interface:
```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

### http.Handle

函数功能：


实现流程：

http.Handle:
```go
func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }
```

DefaultServeMux.Handle:
```go
func (mux *ServeMux) Handle(pattern string, handler Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if pattern == "" {
		panic("http: invalid pattern")
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if _, exist := mux.m[pattern]; exist {
		panic("http: multiple registrations for " + pattern)
	}

	if mux.m == nil {
		mux.m = make(map[string]muxEntry)
	}
	e := muxEntry{h: handler, pattern: pattern}
	mux.m[pattern] = e
	if pattern[len(pattern)-1] == '/' {
		mux.es = appendSorted(mux.es, e)
	}

	if pattern[0] != '/' {
		mux.hosts = true
	}
}
```
handlerFunc底层使用的也是这个Handle函数，只是在外面封装了一个HandlerFunc.

## http.Handler相较于http.HandlerHandler的作用
在go语言高级编程 5.3 节中给了一个例子

```go
func helloHandler(w http.ResponseWriter, r *http.Request)  {
...
}

func showInfoHandler(w http.ResponseWriter, r *http.Request)  {
...
}

func showEmailhelloHandler(w http.ResponseWriter, r *http.Request)  {
...
}

func showFriendshelloHandler(w http.ResponseWriter, r *http.Request)  {
...
}
...

func main()  {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/showInfo", showInfoHandler)
    http.HandleFunc("/showEmail", showEmailHandler)
    http.HandleFunc("/showFriends", showFriendsHello)
	
	...
}

```
每次需要新的路由就直接添加路由以及对应的实现函数。

当我们把路由增加到几十个时， 老板找到你，最近找人新开发了监控系统，为了系统运行可以更加可控，需要把每个接口运行的耗时数据主动上报到我们的监控系统里。给监控系统起个名字，叫metrics。现在你需要修改代码并把耗时通过HTTP Post的方式发给metrics系统了。我们来修改一下helloHandler()：
```go
func helloHandler(w http.ResponseWriter, r *http.Request)  {
    timeStart := time.Now()
	...
	timeElpased :=time.Since(timeStart)
	logger.Println(timeElpased)
	metrics.Upload("timeHandler",timeElpased)
}
```
为单独的一个helloHandler添加监控代码此时还很简单。但是为每一个路由业务函数都添加就变得很困难。

像metrics一类的监控功能属于业务之外的需求。在httq请求前工作，我们可以采用一个函数适配器的方法来包装需要监控的函数。
```go
func helloHandler(w http.ResponseWriter, r *http.Request)  {
...
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        timeStart := time.Now()
        ...
		next.ServeHTTP(w,r)
        timeElpased :=time.Since(timeStart)
        logger.Println(timeElpased)
        metrics.Upload("timeHandler",timeElpased)
})
}

func main()  {
    http.Handler("/hello",timeMiddleware(http.HandlerFunc(helloHandler)))
	
	...
}
```
魔法就在于这个timeMiddleware。timeMiddleware参数为一个http.Handler,

http.Handler定义如下：
```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```
对于任何方法，只要实现了ServeHTTP，它就是一个合法的http.Handler。

### Handler、HandlerFunc和ServeHTTP的关系
Handler:
```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

HandlerFunc:
```go
type HandlerFunc func(w http.ResponseWriter, r *http.Request)
```

ServeHTTP:
```go
func (f HandlerFunc) ServerHttp(w http.ResponseWriter, r *http.Request){
	f(w, r);
}
```

实际上只要你的handler()函数的函数签名是`func (ResponseWriter,*Request)`(实现Handler接口)

那么这个handler和http.HandlerFunc()就有了一致的函数签名，可以将该handler()函数进行类型转换，转换为http.HandlerFunc()。而http.HandlerFunc()实现了http.Handler这个接口。在http库需要调用你的handler()函数来处理HTTP请求时，会调用HandlerFunc()的ServeHTTP()函数，可见一个请求的基本调用链是这样的：

`h=getHandler()=>h.serverHandler(w,r)=>h(w,r)`

## 创建ServeMux

// TODO 默认http.ListenAndServer可能存在的漏洞

调用http.HandleFunc()/http.Handle()都是将处理器/函数注册到ServeMux的默认对象DefaultServeMux上。使用默认对象有一个问题：不可控。

一来Server参数都使用了默认值，二来第三方库也可能使用这个默认对象注册一些处理，容易冲突。更严重的是，我们在不知情中调用http.ListenAndServe()开启 Web 服务，那么第三方库注册的处理逻辑就可以通过网络访问到，有极大的安全隐患。所以，除非在示例程序中，否则建议不要使用默认对象。

我们可以使用http.NewServeMux()创建一个新的ServeMux对象，然后创建http.Server对象定制参数，用ServeMux对象初始化Server的Handler字段，最后调用Server.ListenAndServe()方法开启 Web 服务：

```go
func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", index)
  mux.Handle("/greeting", greeting("Welcome to go web frameworks"))

  server := &http.Server{
    Addr:         ":8080",
    Handler:      mux,
    ReadTimeout:  20 * time.Second,
    WriteTimeout: 20 * time.Second,
  }
  server.ListenAndServe()
}
```


参考：
https://darjun.github.io/2021/07/13/in-post/godailylib/nethttp/