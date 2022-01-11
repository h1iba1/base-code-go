## go 实现arrayList
### interface
```go
type List interface {
	Size() int 									//数组大小
	Get(index int)(interface{},error) 			//获取第几个元素
	Set(index int, newval interface{})error 	//修改数据
	Insert(index int, newval interface{}) error //插入元素
	Append(newval interface{}) error  			//追加
	Clear()  									//清空
	Delete(index int) error 					//删除
	String() string 							//返回字符串
}
```
### ArrayList结构体
```go
type ArrayList struct{
	DataStore[]interface{}  //数组存储
	theSize int 			//数组大小
}
```
### newArrayList方法
```go
func NewArrayList() *ArrayList {
	list:=new(ArrayList)					//初始化结构体
	list.DataStore=make([]interface{},0,10)	//开辟空间10个
	list.theSize=0
	return list
}
```
为list开辟存储空间，设置大小。最后返回一个list对象。

### checkisFull方法
```go
func (list *ArrayList) checkisFull()  {
	if list.theSize==cap(list.DataStore){
		//第二个参数指定开辟空间长度，第三个参数指定预留空间
		//当为0时空间cap虽然=list.DataStore*2,但是len=0,copy无法存入数据
		newDataStore :=make([]interface{},list.theSize,list.theSize*2)

		copy(newDataStore,list.DataStore)
		//将添加数据的方式改为append追加也可以
		//for i:=0;i<list.theSize;i++{
		//	newDataStore=append(newDataStore,list.DataStore[i])
		//}
		list.DataStore=newDataStore
	}
}
```
在插入元素时，首先检查容量是否已满。满的话在原有大小的基础上增加两倍。
make分配内存时，第二个参数指定内存长度（len()获取），第三个参数指定内存空间容量（cap()获取）

### Insert()方法
```go
func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index<0 || index>list.theSize{
		return errors.New("数组越界")
	}
	list.checkisFull()
	//fmt.Println("test:",len(list.DataStore))
	list.DataStore=list.DataStore[:list.theSize+1]  //增加len可获取值
	//fmt.Println("test:",len(list.DataStore))

	//向后移位，移到index为止
	for i:=list.theSize;i>index;i--{
		list.DataStore[i]=list.DataStore[i-1]
	}
	//插入newval
	list.DataStore[index]=newval
	list.theSize++
	return nil
}
```
insert插入数据时，首先判断内存容量是否已满。
然后将index之后的数据依次向后移位，为插入的数据腾出空间。
最后再在index上插入新元素。

### Clear()方法
```go
func (list *ArrayList) Clear(){
	list.DataStore=make([]interface{},0,10)
	list.theSize=0
}
```
清除空间所有数据，重新分配内存空间。