## StackArray
### StackArray interface
```go
type StackArray interface{
	Clear()					//清空
	Size() int				//大小
	Pop()interface{}	    //出栈
	Push(data interface{})	//入栈
	IsFull() bool			//是否满了
	IsEmpty() bool			//是否为空
}
```
#### Clear()         清空栈
```go
func (mystack *Stack) Clear(){
	mystack.dataSource=make([]interface{},0,10) //开辟内存空间
	mystack.capsize=10  	//容量
	mystack.currentsize=0
}
```

#### Size() int      栈当前大小
```go
func (mystack *Stack) Size() int{
	return mystack.currentsize
}
```

#### Pop()interface{} 出栈
前提：栈不为空
```go
func (mystack *Stack) Pop() interface{}{
	if !mystack.IsEmpty(){
		last :=mystack.dataSource[mystack.currentsize-1]
		// 删除出栈的数据
		mystack.dataSource=mystack.dataSource[:mystack.currentsize-1]
		mystack.currentsize--
		return last
	}
	return nil
}
```
出栈之后元素删除，size--


#### Push(data interface{}) 入栈
前提：栈还有大小
```go
func (mystack *Stack)  Push(data interface{}){
	if !mystack.IsFull(){
		mystack.dataSource=append(mystack.dataSource,data)
		mystack.currentsize++
	}
}
```
元素添加到最后一个，size++

#### IsFull() bool 栈是否已满
```go
func (mystack *Stack) IsFull() bool{
	return mystack.capsize==mystack.currentsize
}
```

#### IsEmpty() bool 栈是否已空
```go
func (mystack *Stack) IsEmpty() bool{
	return mystack.currentsize==0
}
```

### Stack struct
```go
type Stack struct {
	dataSource[] interface{}
	capsize int 	 	//最大范围
	currentsize int 	//实际使用大小
}
```

### NewStack() *Stack
```go
func NewStack() *Stack {
	mystack:=new(Stack)		//声明一个mystack
	mystack.dataSource=make([]interface{},0,10) //开辟内存空间
	mystack.capsize=10  	//容量
	mystack.currentsize=0
	return mystack
}
```