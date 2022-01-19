## 应用场景
在go http 包的 server 中，每一个请求都有一个 goroutine 去处理。请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和rpc服务。用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息，验证相关的token,请求的截止时间。**当一个请求被取消或者超时时，所有用来处理该请求的 goroutine 的应该迅速退出，然后系统才能释放这些goroutine占用的资源。**

注意：`go1.6`及之前版本请使用`golang.org/x/net/context`。`go1.7`及之后已移到标准库`context`。

## context原理
context 的调用应该是链式的，通过 withCancel，withDeadline，withTimeout 或 withValue 派生出新的 context 。当父 Context 被取消时，其派生的所有 Context 都将取消。
通过 context.Withxxx 都将返回新的 context 和 CancelFunc 。调用 CancelFunc 将取消子代，移除父代对子代的引用，并停止所有定时器。未能调用 CancelFunc 将泄漏子代，直到父代被取消或定时器触发。`go vet`工具检查所有流程控制路径上使用`CancelFuncs`。

## 遵循规则
遵循以下规则，以保持包之间的接口一致，并启用静态分析工具以检查上下文传播。
1.不要将`context`放入结构体，相反`context`应该作为第一个参数传入，命名为`ctx`。`func DoSomething(ctx context.Context,arg Arg) error {//...use ctx...}`
2.**即使函数允许，也不要传入nil的`Context`。如果不知道用那种`context`,可以使用`context.TODO()`**。
3.使用`context`的`value`相关方法只应该用于在程序和接口中传递的和请求相关的元数据，不要用它来传递一些可选的参数。
4.相同的`Context`可以传递在不同的`goroutine`；`Context`是并发安全的。

## COntext包
### Context结构体
```go
// A Context carries a deadline, cancelation signal, and request-scoped values  
// across API boundaries. Its methods are safe for simultaneous use by multiple  
// goroutines.  
type Context interface {  
 // Done returns a channel that is closed when this Context is canceled  
 // or times out.  
 Done() <-chan struct{}  
	
 // Err indicates why this context was canceled, after the Done channel  
 // is closed.  
 Err() error  

 // Deadline returns the time when this Context will be canceled, if any.  
 Deadline() (deadline time.Time, ok bool)  
 
 // Value returns the value associated with key or nil if none.  
 Value(key interface{}) interface{}  
}
```

- Done()，返回一个channel。当times out 或者调用 cancel方法时，将会close掉。
- Err(), 返回一个错误。该 context 为什么被取消掉。
- Deadline()，返回截止时间和ok。
- Value()，返回值。

### 所有方法

```go
func Background() Context
func TODO() Context

func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context
```

上面可以看到 Context 是一个接口，想要使用就得实现其方法。在context包内部已经为我们实现好了两个空的 Context ，可以通过调用 BackGround() 和 TODO() 方法获取。一般将他们作为 Context 的根，往下派生。

### WithCancel
WithCancel 以一个新的 Done channel 返回一个父 Context 的拷贝。
```go
// WithCancel returns a copy of parent with a new Done channel. The returned
// context's Done channel is closed when the returned cancel function is called
// or when the parent context's Done channel is closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	c := newCancelCtx(parent)
	propagateCancel(parent, &c)
	return &c, func() { c.cancel(true, Canceled) }
}

// newCancelCtx returns an initialized cancelCtx.
func newCancelCtx(parent Context) cancelCtx {
	return cancelCtx{Context: parent}
}
```

此示例演示使用一个可取消的上下文，以防止 goroutine 泄漏。示例函数结束时，defer 调用 cancel 方法，gen goroutine 将返回而不泄漏。
```go
package main

import (
	"context"
	"fmt"
)

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
```

### WithDeadline 例子
```go
// WithDeadline returns a copy of the parent context with the deadline adjusted
// to be no later than d. If the parent's deadline is already earlier than d,
// WithDeadline(parent, d) is semantically equivalent to parent. The returned
// context's Done channel is closed when the deadline expires, when the returned
// cancel function is called, or when the parent context's Done channel is
// closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if cur, ok := parent.Deadline(); ok && cur.Before(d) {
		// The current deadline is already sooner than the new one.
		// 当前的截止日期已经比新的截止日期早了。
		// 新的截止日期，在父截止日期之后，直接返回一个父 Context 的拷贝
		return WithCancel(parent)
	}
	c := &timerCtx{
		cancelCtx: newCancelCtx(parent),
		deadline:  d,
	}
	propagateCancel(parent, c)
	dur := time.Until(d)
	if dur <= 0 {
		c.cancel(true, DeadlineExceeded) // deadline has already passed
		return c, func() { c.cancel(false, Canceled) }
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.err == nil {
		c.timer = time.AfterFunc(dur, func() {
			c.cancel(true, DeadlineExceeded)
		})
	}
	return c, func() { c.cancel(true, Canceled) }
}
```

#### 新的截止日期，在父截止日期之后，直接返回一个父 Context 的拷贝
```go
if cur, ok := parent.Deadline(); ok && cur.Before(d) {
		// The current deadline is already sooner than the new one.
		// 当前的截止日期已经比新的截止日期早了。
		// 新的截止日期，在父截止日期之后，直接返回一个父 Context 的拷贝
		return WithCancel(parent)
	}
```
可以清晰的看到，当派生出的子 Context 的deadline在父 Context 之后，直接返回了一个父Context的拷贝。故语义上等效为父。

#### 新的截止日期，在父截止日期之前
如果父节点的截止日期已经早于 d，WithDeadline(parent, d) 在语义上等价于父节点。返回的上下文的 Done 通道在截止日期到期、**调用返回的取消函数或父上下文的 Done 通道关闭时关闭，以先发生者为准。**

官方例子：
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
```

### WithTimeOut 例子
传递一个带有超时的上下文来告诉一个阻塞函数它应该在超时后放弃它的工作。

WithTimeout 返回 WithDeadline(parent, time.Now().Add(timeout))。

官方例子：
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
```

### WitchValue例子
```go
// WithValue returns a copy of parent in which the value associated with key is
// val.
//
// Use context Values only for request-scoped data that transits processes and
// APIs, not for passing optional parameters to functions.
//
// The provided key must be comparable and should not be of type
// string or any other built-in type to avoid collisions between
// packages using context. Users of WithValue should define their own
// types for keys. To avoid allocating when assigning to an
// interface{}, context keys often have concrete type
// struct{}. Alternatively, exported context key variables' static
// type should be a pointer or interface.
func WithValue(parent Context, key, val interface{}) Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if key == nil {
		panic("nil key")
	}
	if !reflectlite.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	return &valueCtx{parent, key, val}
}
```
WithValue 返回的父与键关联的值在 val 的副本。

使用上下文值仅为过渡进程和 Api 的请求范围的数据，而不是将可选参数传递给函数。

提供的键必须是可比性和应该不是字符串类型或任何其他内置的类型以避免包使用的上下文之间的碰撞。WithValue 用户应该定义自己的键的类型。为了避免分配分配给接口 {} 时，上下文键经常有具体类型结构 {}。另外，导出的上下文关键变量静态类型应该是一个指针或接口。

官方例子：
```go
package main

import (
	"context"
	"fmt"
)

func main() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}
```

## 使用场景
### 1. http 服务器的request互相传递数据
context还提供了valueCtx的数据结构。

这个valueCtx最经常使用的场景就是在一个http服务器中，在request中传递一个特定值，比如有一个中间件，做cookie验证，然后把验证后的用户名存放在request中。

我们可以看到，官方的request里面是包含了Context的，并且提供了WithContext的方法进行context的替换。

```go
package main

import (
	"context"
	"net/http"
)

type FooKey string

var UserName = FooKey("user-name")
var UserId = FooKey("user-id")

func foo(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), UserId, "1")
		ctx2 := context.WithValue(ctx, UserName, "yejianfeng")
		next(w, r.WithContext(ctx2))
	}
}

func GetUserName(context context.Context) string {
	if ret, ok := context.Value(UserName).(string); ok {
		return ret
	}
	return ""
}

func GetUserId(context context.Context) string {
	if ret, ok := context.Value(UserId).(string); ok {
		return ret
	}
	return ""
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome: "))
	w.Write([]byte(GetUserId(r.Context())))
	w.Write([]byte(" "))
	w.Write([]byte(GetUserName(r.Context())))
}

func main() {
	http.Handle("/", foo(test))
	http.ListenAndServe(":8080", nil)
}

```
在使用ValueCtx的时候需要注意一点，这里的key不应该设置成为普通的String或者Int类型，为了防止不同的中间件对这个key的覆盖。最好的情况是每个中间件使用一个自定义的key类型，比如这里的FooKey，而且获取Value的逻辑尽量也抽取出来作为一个函数，放在这个middleware的同包中。这样，就会有效避免不同包设置相同的key的冲突问题了。

### 2.超时请求
我们发送RPC请求的时候，往往希望对这个请求进行一个超时的限制。当一个RPC请求超过10s的请求，自动断开。当然我们使用CancelContext，也能实现这个功能（开启一个新的goroutine，这个goroutine拿着cancel函数，当时间到了，就调用cancel函数）。

鉴于这个需求是非常常见的，context包也实现了这个需求：timerCtx。具体实例化的方法是 WithDeadline 和 WithTimeout。

具体的timerCtx里面的逻辑也就是通过time.AfterFunc来调用ctx.cancel的。
官方例子：
```go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
    defer cancel()

    select {
    case <-time.After(1 * time.Second):
        fmt.Println("overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err()) // prints "context deadline exceeded"
    }
}

```
在http的客户端里面加上timeout也是一个常见的办法
```go
uri := "https://httpbin.org/delay/3"
req, err := http.NewRequest("GET", uri, nil)
if err != nil {
	log.Fatalf("http.NewRequest() failed with '%s'\n", err)
}

ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100)
req = req.WithContext(ctx)

resp, err := http.DefaultClient.Do(req)
if err != nil {
	log.Fatalf("http.DefaultClient.Do() failed with:\n'%s'\n", err)
}
defer resp.Body.Close()

```

在http服务端设置一个timeout如何做呢？
```go
package main

import (
	"net/http"
	"time"
)

func test(w http.ResponseWriter, r *http.Request) {
	time.Sleep(20 * time.Second)
	w.Write([]byte("test"))
}


func main() {
	http.HandleFunc("/", test)
	timeoutHandler := http.TimeoutHandler(http.DefaultServeMux, 5 * time.Second, "timeout")
	http.ListenAndServe(":8080", timeoutHandler)
}

```
我们看看TimeoutHandler的内部，本质上也是通过context.WithTimeout来做处理。
```go
func (h *timeoutHandler) ServeHTTP(w ResponseWriter, r *Request) {
  ...
		ctx, cancelCtx = context.WithTimeout(r.Context(), h.dt)
		defer cancelCtx()
	...
	go func() {
    ...
		h.handler.ServeHTTP(tw, r)
	}()
	select {
    ...
	case <-ctx.Done():
		...
	}
}
```

### 3.PipeLine
pipeline模式就是流水线模型，流水线上的几个工人，有n个产品，一个一个产品进行组装。其实pipeline模型的实现和Context并无关系，没有context我们也能用chan实现pipeline模型。但是对于整条流水线的控制，则是需要使用上Context的。这篇文章Pipeline Patterns in Go的例子是非常好的说明。这里就大致对这个代码进行下说明。

runSimplePipeline的流水线工人有三个，lineListSource负责将参数一个个分割进行传输，lineParser负责将字符串处理成int64,sink根据具体的值判断这个数据是否可用。他们所有的返回值基本上都有两个chan，一个用于传递数据，一个用于传递错误。（<-chan string, <-chan error）输入基本上也都有两个值，一个是Context，用于传声控制的，一个是(in <-chan)输入产品的。

我们可以看到，这三个工人的具体函数里面，都使用switch处理了case <-ctx.Done()。这个就是生产线上的命令控制。
```go
func lineParser(ctx context.Context, base int, in <-chan string) (
	<-chan int64, <-chan error, error) {
	...
	go func() {
		defer close(out)
		defer close(errc)

		for line := range in {

			n, err := strconv.ParseInt(line, base, 64)
			if err != nil {
				errc <- err
				return
			}

			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out, errc, nil
}
```

### 4.RPC调用
在主goroutine上有4个RPC，RPC2/3/4是并行请求的，我们这里希望在RPC2请求失败之后，直接返回错误，并且让RPC3/4停止继续计算。这个时候，就使用的到Context。

这个的具体实现如下面的代码。
```go
package main

import (
	"context"
	"sync"
	"github.com/pkg/errors"
)

func Rpc(ctx context.Context, url string) error {
	result := make(chan int)
	err := make(chan error)

	go func() {
		// 进行RPC调用，并且返回是否成功，成功通过result传递成功信息，错误通过error传递错误信息
		isSuccess := true
		if isSuccess {
			result <- 1
		} else {
			err <- errors.New("some error happen")
		}
	}()

	select {
		case <- ctx.Done():
			// 其他RPC调用调用失败
			return ctx.Err()
		case e := <- err:
			// 本RPC调用失败，返回错误信息
			return e
		case <- result:
			// 本RPC调用成功，不返回错误信息
			return nil
	}
}


func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// RPC1调用
	err := Rpc(ctx, "http://rpc_1_url")
	if err != nil {
		return
	}

	wg := sync.WaitGroup{}

	// RPC2调用
	wg.Add(1)
	go func(){
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_2_url")
		if err != nil {
			cancel()
		}
	}()

	// RPC3调用
	wg.Add(1)
	go func(){
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_3_url")
		if err != nil {
			cancel()
		}
	}()

	// RPC4调用
	wg.Add(1)
	go func(){
		defer wg.Done()
		err := Rpc(ctx, "http://rpc_4_url")
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()
}
```
当然我这里使用了waitGroup来保证main函数在所有RPC调用完成之后才退出。

在Rpc函数中，第一个参数是一个CancelContext, 这个Context形象的说，就是一个传话筒，在创建CancelContext的时候，返回了一个听声器（ctx）和话筒（cancel函数）。所有的goroutine都拿着这个听声器（ctx），当主goroutine想要告诉所有goroutine要结束的时候，通过cancel函数把结束的信息告诉给所有的goroutine。当然所有的goroutine都需要内置处理这个听声器结束信号的逻辑（ctx->Done()）。我们可以看Rpc函数内部，通过一个select来判断ctx的done和当前的rpc调用哪个先结束。

这个waitGroup和其中一个RPC调用就通知所有RPC的逻辑，其实有一个包已经帮我们做好了。[errorGroup](https://godoc.org/golang.org/x/sync/errgroup)。具体这个errorGroup包的使用可以看这个包的test例子。

有人可能会担心我们这里的cancel()会被多次调用，context包的cancel调用是幂等的。可以放心多次调用。

我们这里不妨品一下，这里的Rpc函数，实际上我们的这个例子里面是一个“阻塞式”的请求，这个请求如果是使用http.Get或者http.Post来实现，实际上Rpc函数的Goroutine结束了，内部的那个实际的http.Get却没有结束。所以，需要理解下，这里的函数最好是“非阻塞”的，比如是http.Do，然后可以通过某种方式进行中断。比如像这篇文章[Cancel http.Request using Context](https://medium.com/@ferencfbin/golang-cancel-http-request-using-context-1f45aeba6464)中的这个例子：

```go
func httpRequest(
  ctx context.Context,
  client *http.Client,
  req *http.Request,
  respChan chan []byte,
  errChan chan error
) {
  req = req.WithContext(ctx)
  tr := &http.Transport{}
  client.Transport = tr
  go func() {
    resp, err := client.Do(req)
    if err != nil {
      errChan <- err
    }
    if resp != nil {
      defer resp.Body.Close()
      respData, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        errChan <- err
      }
      respChan <- respData
    } else {
      errChan <- errors.New("HTTP request failed")
    }
  }()
  for {
    select {
    case <-ctx.Done():
      tr.CancelRequest(req)
      errChan <- errors.New("HTTP request cancelled")
      return
    case <-errChan:
      tr.CancelRequest(req)
      return
    }
  }
}
```
它使用了http.Client.Do，然后接收到ctx.Done的时候，通过调用transport.CancelRequest来进行结束。  
我们还可以参考[net/dail/DialContext](https://tip.golang.org/src/net/dial.go?s=11446:11534#L351)  
换而言之，如果你希望你实现的包是“可中止/可控制”的，那么你在你包实现的函数里面，最好是能接收一个Context函数，并且处理了Context.Done。
