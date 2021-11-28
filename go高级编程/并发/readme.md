
要理解 Golang 编程思维，首先要理解 Golang 这门语言的创始初衷，初衷就是为了解决好 Google 内部大规模高并发服务的问题，主要核心就是围绕高并发来开展；并且同时又不想引入面向对象那种很复杂的继承关系。首先，就是可以方便的解决好并发问题（包括高并发），那么就需要有并发思维，能够并发处理就通过并发来进行任务分配

- 这个就是涉及到了 context、 goroutine、channel（select）；
- 创建大量 goroutine， 但是需要能通过 context、 channel 建立 "父子"关系，保证子任务可以能够被回收、被主动控制（如 杀死）。

## 并发的内存模型


### Goroutine和系统线程
Goroutine是Go语言特有的并发体，是一种轻量级的线程，由go关键字启动。

```go
func f(){
	fmt.Println("f function")
}

func main()  {
    go f()
	time.Sleep(1 *time.Second)
	fmt.Println("main function")
}
```
f()函数以go开头意味着f()将作为一个单独的goroutine运行与main()同时运行。

main程序存在time.Sleep(1 *time.Second)阻塞程序执行完成，大多数情况下程序会先输出
`f function`在输出`main function`。但这不是一个优雅的做法，并不能百分百保证main()在f()执行之后执行，这个时候就需要使用go的特殊类型channel来帮助f()和main()建立通信。

### 原子操作
所谓的原子操作就是并发编程中“最小的且不可并行化”的操作

通常，如果多个并发体对同一个共享资源进行的操作是原子的话，那么同一时刻最多只能有一个并发体对该资源进行操作。从线程角度看，在当前线程修改共享资源期间，其他线程是不能访问该资源的

go语言的标准库sync/atomic包对原子操作提供了丰富的支持。
```go
import (
    "sync"
	"sync/atomic"
)

var total uint64

func worker(wg *sync.WaitFroup)  {
    defer wg.Done()  //执行完worker之后结束wg
	var i uint64
	for i=0;i<=100;i++{
		atomic.AddUint64(&total,i)
    }
}

func main()  {
	var wg sync.WaitGroup
	wg.Add(2) // 添加两个WaitGroup
	go worker(&wg)
    go worker(&wg)
    wg.wait()  //只有当所有的wg结束之后才结束当前goroutine
}
```
atomic.AddUint64 ()函数调用保证了total的读取、更新和保存是一个原子操作，因此在多线程中访问也是安全的。

### 内存一致性模型
在Go语言中，同一个Goroutine线程内部，顺序一致性的内存模型是得到保证的。但是不同的Goroutine之间，并不满足顺序一致性的内存模型，需要通过明确定义的同步事件来作为同步的参考。如果两个事件不可排序，那么就说这两个事件是并发的。为了最大化并行，Go语言的编译器和处理器在不影响上述规定的前提下可能会对执行语句重新排序（CPU也会对一些指令进行乱序执行）。

如果在一个Goroutine中顺序执行a=1; b=2;这两个语句，虽然在当前的Goroutine中可以认为a=1;语句先于b=2;语句执行，但是在另一个Goroutine中b=2;语句可能会先于a=1;语句执行，甚至在另一个Goroutine中无法看到它们的变化（可能始终在寄存器中）。也就是说在另一个Goroutine看来,a=1; b=2;这两个语句的执行顺序是不确定的。如果一个并发程序无法确定事件的顺序关系，那么程序的运行结果往往会有不确定的结果

```go
func main() {
    go println("你好，世界")
}
```
根据Go语言规范，main()函数退出时程序结束，不会等待任何后台线程。因为Goroutine的执行和main()函数的返回事件是并发的，谁都有可能先发生，所以什么时候打印、能否打印都是未知的。

用前面的原子操作并不能解决问题，因为我们无法确定两个原子操作之间的顺序。解决问题的办法就是通过同步原语来给两个事件明确排序：
```go
func main() {
    done :=make(chan int)
	go func(){
		println("你好，世界")
		done <-1 //向done发送数据
    }()
	<-done  //接收done接收到的数据
}
```
当<-done执行时，必然要求done <- 1也已经执行。根据同一个Goroutine依然满足顺序一致性规则，可以判断当done <- 1执行时，println(" 你好 , 世界 ")语句必然已经执行完成了。因此，现在的程序确保可以正常打印结果。

### Goroutine的创建
go语句会在当前Goroutine对应函数返回前创建新的Goroutine。例如:
```go
var a string

func f()  {
	print(a)
}

func hello()  {
    a ="hello,world"
	go f()
}
```
执行go f()语句创建Goroutine和hello()函数是在同一个Goroutine中执行，根据语句的书写顺序可以确定Goroutine的创建发生在hello()函数返回之前，但是新创建Goroutine对应的f()的执行事件和hello()函数返回的事件则是不可排序的，也就是并发的。调用hello ()可能会在将来的某一时刻打印“hello, world”，也很可能是在hello()函数执行完成后才打印。

### 基于通道的通信
通道（channel）是在Goroutine之间进行同步的主要方法。

在无缓存的通道上的每一次发送操作都有与其对应的接收操作相匹配，发送和接收操作通常发生在不同的Goroutine上（在同一个Goroutine上执行两个操作很容易导致死锁）。**无缓存的通道上的发送操作总在对应的接收操作完成前发生。**

```go
var done =make(chan bool)
var msg string
func aGoroutine()  {
	msg ="你好，世界"
	done <-true
}
func main()  {
    go aGoroutine()
	<-done
	println(msg)
}
```
可保证打印出“你好, 世界”。该程序首先对msg进行写入，然后在done通道上发送同步信号，随后从done接收对应的同步信号，最后执行println()函数。

**对于从无缓存通道进行的接收，发生在对该通道进行的发送完成之前。**

我们可以根据控制通道的缓存大小来控制并发执行的Goroutine的最大数目，例如：
```go
var limit =make(chan int,3)
func main()  {
	for _,w :=range work{
		go func(){
			/*
			这里limit的容量为三，可最多容纳三个w()同时执行
			*/
			limit <-1  
			w()
			<-limit
        }
    }
    select{}
}
```
最后一句select{}是一个空的通道选择语句，该语句会导致main线程阻塞，从而避免程序过早退出。还有for{}、<-make(chan int)等诸多方法可以达到类似的效果。因为main线程被阻塞了，如果需要程序正常退出的话，可以通过调用os.Exit(0)实现。

### context

context 用来解决 goroutine 之间退出通知、元数据传递的功能。

[深度解密Go语言之context](https://zhuanlan.zhihu.com/p/68792989)

### 参考
《go语言高级编程》1.5面向并发的内存模型

[Golang 编程思维和工程实战](https://mp.weixin.qq.com/s/llmE9QpnrvA02AtvfHtqJQ)

