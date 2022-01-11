## ArrayList 迭代器
### Iterator接口
```go
type Iterator interface{
	HasNext() bool				//是否有下一个
	Next()(interface{},error)	//下一个
	Remove()					//删除
	GetIndex()	int				//得到索引
}
```
实现了HasNext()，Next()(interface{},error)，Remove()，GetIndex()四个方法的struct我们就认为是迭代器。

### Iterable接口
```go
type Iterable interface {
	Iterator() Iterator         //构造初始化接口
}
```
实现了Iterator()迭代器初始化函数的struct就认为是Iterable

### ArrayListIterator struct
```go
type ArrayListIterator struct{
	list *ArrayList  //指针对象
	currentIndex int //当前索引
}
```
ArrayList迭代器主要有两个参数。ArrayList的指针对象，和当前的索引。

### Iterator()方法
```go
func (list *ArrayList) Iterator() Iterator {
	it :=new(ArrayListIterator)
	it.currentIndex=0
	it.list=list
	return it
}
```
构造器初始化方法。ArrayList的构造器，所以需要申明`(list *ArrayList)`标签。

想要list实现迭代器，
list接口也需要声明`Iterator() Iterator`

ArrayListIterator实现了Iterator接口。
通过在list接口中添加`Iterator() Iterator`的方式，来将Iterable和list接口串联起来。






