
## 2.1优雅的实现构造函数编程思想

```go
package main
/*
一个更为优雅的构造函数的实现方式
参考：
https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
通过这个方式可以方便构造不同对象，同时避免了大量重复代码
*/

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

type Cluster struct {
	opts options
}

type options struct {
	connectionTimeout time.Duration
	readTimeout       time.Duration
	writeTimeout      time.Duration
	logError          func(ctx context.Context, err error)
}

// Option 通过一个选项实现为一个函数指针来达到一个目的：设置选项中的数据的状态
// Golang函数指针的用法
type Option func(c *options)

// LogError 设置某个参数的一个具体实现，用到了闭包的用法。
// 不仅仅只是设置而采用闭包的目的是为了更为优化，更好用，对用户更友好
func LogError(f func(ctx context.Context, err error)) Option {

	//相当于SetOptions
	// *options 修改的是地址上的值
	return func(opts *options) {
		opts.logError = f
	}
}

// ConnectionTimeout 对关键数据变量的赋值采用一个方法来实现而不是直接设置
func ConnectionTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.connectionTimeout = d
	}
}

func WriteTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.writeTimeout = d
	}
}

func ReadTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.readTimeout = d
	}
}

// NewCluster 构造函数具体实现，传入相关Option，new一个对象并赋值
// 如果参数很多，也不需要传入很多参数，只需要传入opts ...Option即可
func NewCluster(opts ...Option) *Cluster {
	clusterOpts := options{}
	for _, opt := range opts {
		// 函数指针的赋值调用
		opt(&clusterOpts)
	}

	cluster := new(Cluster)
	cluster.opts = clusterOpts

	return cluster
}

func main() {

	// 前期储备，设定相关参数
	commonsOpts := []Option{
		ConnectionTimeout(1 * time.Second),
		ReadTimeout(2 * time.Second),
		WriteTimeout(3 * time.Second),
		LogError(func(ctx context.Context, err error) {
		}),
	}

	// 终极操作，构造函数
	cluster := NewCluster(commonsOpts...)

	// 测试验证
	fmt.Println(cluster.opts.connectionTimeout)
	fmt.Println(cluster.opts.writeTimeout)

}
```

### 1.通过闭包的方式为参数赋值
```go
// ConnectionTimeout 对关键数据变量的赋值采用一个方法来实现而不是直接设置
func ConnectionTimeout(d time.Duration) Option {
	return func(opts *options) {
		opts.connectionTimeout = d
	}
}
```
通过闭包的方式直接返回一个值的设置,使代码结构更优雅

### 2.构造函数初始化
`clusterOpts := options{}`仅仅起到一个重命名的作用。空结构体不占用内存。




reference: https://mp.weixin.qq.com/s/llmE9QpnrvA02AtvfHtqJQ 