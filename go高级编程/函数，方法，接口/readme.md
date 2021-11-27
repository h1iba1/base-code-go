## 函数

### 闭包
当匿名函数引用了外部作用域中的变量时就成了闭包函数，闭包函数是函数式编程的核心。

```go
var Add =func(a,b int )int{
	return a+b
}
```
闭包函数可以直接将函数作为一个参数传递给一个变量。

闭包对捕获的外部变量并不是以传值方式访问，而是以引用方式访问。


### 可变数量参数
在语法上，函数还支持可变数量的参数，可变数量的参数必须是最后出现的参数，可变数量的参数其实是一个切片类型的参数。

```go
// more对应int切片类型
func Sum(a int ,more ...int) int {
	for _,v :=range more{
		a+=v
}
	return a
}

```

### 函数返回值命名

如果返回值命名了，可以通过名字来修改返回值，也可以通过defer语句在return语句之后修改返回值：
```go
func Inc()(v int){
	defer func(){v++ }()
	return 42
}
```

## 方法

方法一般是面向对象编程oop的一个特性。
在C++语言中方法对应一个类对象的成员函数，是关联到具体对象上的虚表中的。但是Go语言的方法却是关联到类型的，这样可以在编译阶段完成方法的静态绑定。

### go oop
go实现oop只需要将类型作为一个指针变量写在方法的开头即可表明这个方法属于所写的类型。
```go
// 关闭文件
func (f *File) CloseFile() error{
	// ...
}
// 读取文件数据
func (f *File) ReadFile(int64 offset,data []byte)int  {
    // ...
}
```
通过这种编写方式函数CloseFile()和ReadFile()就成了File类型独有的方法了（而不是File对象方法

此时方法名字还可以再简化一下：
```go
// 关闭文件
func (f *File) Close() error{
	// ...
}
// 读取文件数据
func (f *File) Read(int64 offset,data []byte)int  {
    // ...
}
```

### 继承
Go语言不支持传统面向对象中的继承特性，而是以自己特有的组合方式支持了方法的继承。Go语言中，通过在结构体内置匿名的成员来实现继承
```go
import "image/color"
type Point struct{x,y float64}
type ColoredPoint struct{
	Point
	Color color.RGBA
}
```
通过嵌入匿名的成员，不仅可以继承匿名成员的内部成员，而且可以继承匿名成员类型所对应的方法。我们一般会将Point看作基类，把ColoredPoint看作Point的继承类或子类。

### 接口

Go的接口类型是对其他类型行为的抽象和概括，因为接口类型不会和特定的实现细节绑定在一起，通过这种抽象的方式我们可以让对象更加灵活和更具有适应能力


所谓鸭子类型说的是：只要走起路来像鸭子、叫起来也像鸭子，那么就可以把它当作鸭子

Go语言中的面向对象就是如此，如果一个对象只要看起来像是某种接口类型的实现，那么它就可以作为该接口类型使用。这种设计可以让你创建一个新的接口类型满足已经存在的具体类型却不用去破坏这些类型原有的定义。当使用的类型来自不受我们控制的包时这种设计尤其灵活有用。Go语言的接口类型是延迟绑定，可以实现类似虚函数的多态功能

```go
type Friend interface{
	sayHello()
}

type Dog struct {}

func (d *Dog) sayHello  {
    fmt.Println("hello")
}
```
存在一个friend接口，dog实现了sayHello说明dog也是friend。

满足类似虚函数的多态功能体现在
```go
type Friend interface{
	sayHello()
}

type Cat struct {}

func (c *Cat) sayHello  {
    fmt.Println("hello")
}
```
cat实现了sayHello，也说明cat是friend

这也被称为鸭子类型。“当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子。”

## 值传递与引用传递

### go语言的值类型；

int、float、bool、array、sturct等

值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数

声明一个值类型变量时，编译器会在栈中分配一个空间，空间里存储的就是该变量的值　　

### go语言引用传递：

slice，map，channel，interface，func，string等

声明一个引用类型的变量，编译器会把实例的内存分配在堆上

string和其他语言一样，是引用类型，string的底层实现struct String { byte* str; intgo len; }; 但是因为string不允许修改，每次操作string只能生成新的对象，所以在看起来使用时像值类型。

所谓引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

### go语言指针类型

一个指针变量指向了一个值的内存地址

当一个指针被定义后没有分配到任何变量时，它的值为 nil。nil 指针也称为空指针

一个指针变量通常缩写为 ptr

其实引用类型可以看作对指针的封装


### 参考
《go语言高级编程》1.4函数，方法，接口












