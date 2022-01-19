## 介绍

列表又称线性表的顺序存储。指的是一组**地址连续的存储单元依次存储线性表的数据元素**。

数组，列表，切片都可以称为线性表的顺序存储。

## 接口

```go
type List interface {
	Size() int 															//数组大小
	Get(index int)(interface{},error) 			//获取第几个元素
	Set(index int, newval interface{}) error 	//修改数据
	Insert(index int, newval interface{}) error //插入元素
	Append(newval interface{}) error  			//追加
	Clear()  																//清空
	Delete(index int) error 								//删除
	String() string 												//返回字符串
}
```

## 实现

### List 结构

list结构体，主要含有两个元素。DataStore[] 存储数据，theSize存储list大小。

```go
type Array struct{
  DataStore[] interface{}				// 数组存储
  theSize int										// 数组大小
}
```

### NewList

NewArrayList直接创建一个list，初始空间大小为10（make开辟空间，第二个参数指定开辟空间长度(length)，第三个参数指定预留空间(cap)），初始大小为0。

```go
func NewArrayList() *interface{
  list :=new(ArrayList)										//结构体初始化
  list.DataStore=make([]interface{},0,10)	//开辟空间10个
  list.theSize=0
  return list
}
```

### Size()

获取list大小。

```go
func (list *ArrayList) Size() int{
  return list.theSize
}
```

### Get(index int)(interface{},error)

获取第index个元素。

前置条件：0<index<list.theSize 。否则无法获取到元素。

```go
func (list *ArrayList) Get(index int) (interface{},error){
  if index<0 || index>list.theSize{
    return nil,errors.New("list越界")
  }
  // 切片直接获取第index个元素
  return list.DataStore[index],nil
}
```

### Set(index int, newval interface{}) error 

将第index个元素修改为newel。

前置条件：0<index<list.theSize 。否则无法获取元素并修改。

```go
func (list *ArrayList) Set(index int,newval interface{}) (interface{},error){
  if index<0 || index>list.theSize{
    return nil,errors.New("list越界")
  }
  //切片形式直接修改第index个元素
 	list.DataStore[index]=newval
  return nil
}
```

### Insert(index int, newval interface{}) error 

在第index个位置插入元素newval。

前置条件：

0<index<list.theSize 。否则无法获取第index，并插入。

检查list内存是否满，满则 (x2) 扩展空间。

#### checkisFull()

```go
func (list *ArrayList) checkisFull(){
  // list大小等于list内存容量
  if list.theSize==cap(list.DataStore){
    // make第二个参数表示开辟空间的长度，第三个参数指定预留空间
    newDataStore :=make([]interface{},list.theSize,list.theSize*2)
    // 将原始list copy 到new list
    copy(newDataStore,list.DataStore)
    list.DataStore=newDataStore
  }
}
```

#### Insert实现

```go
func (list *ArrayList) Insert(index int,newval interface{}) error{
  if index<0 || index>list.theSize{
    return errors.New("list 越界")
  }
  
  list.checkisFull()
  // list 扩容
  list.DataStor=list.DataStore[:list.theSize+1]
  // 向后移位，移到index为止
  for i:=list.theSize;list>index;i--{
    // 下标从0开始，第i个就是扩容之后的最后一个。赋值为他的前一个，以此达到向后位移的目的
    list.DataStore[i]=list.DataStore[i-1]
  }
  // 向后移位之后，复制第index个位置为newval即可达到插入的目的
  list.DataStore[index]=newval
  list.theSize++
  return nil
}
```

### Append(newval interface{}) error

向list末尾追加元素。

```go
func (list *ArrayList) Append(newval interface{}) error{
  list.DataStore=append(list.DataStore,newVal)
  list.theSize++
  return nil
}
```

### 	Clear()

清空list.

```go
func (list *ArrayList) Clear(){
  list.DataStore=make([]interface{},0,10)
  list.theSize=0
}
```

### Delete(index int) error

删除第index个元素。

```go
func (list *ArrayList) Delete(index int) error{
  // 以index为分界线，切片DataStore然后组合，达到删除第index个元素的目的（说明[:index]切片，只到index-1个元素）
  list.DataStore=append(list.DataStore[:index],list.DataStore[index+1:]...)
  list.theSize--
  return nil
}
```

### String() string

返回所有元素。

```go
func (list *ArrayList) String() string{
	return fmt.Sprint(list.DataStore)
}
```

