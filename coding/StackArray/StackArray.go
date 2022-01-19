package StackArray

type StackArray interface{
	Clear()					//清空
	Size() int				//大小
	Pop()interface{}	    //出栈
	Push(data interface{})	//入栈
	IsFull() bool			//是否满了
	IsEmpty() bool			//是否为空
}

type Stack struct {
	dataSource[] interface{}
	capsize int 	 	//最大范围
	currentsize int 	//实际使用大小
}

func NewStack() *Stack {
	mystack:=new(Stack)		//声明一个mystack
	mystack.dataSource=make([]interface{},0,1000) //开辟内存空间
	mystack.capsize=10  	//容量
	mystack.currentsize=0
	return mystack
}

func (mystack *Stack) Clear(){
	mystack.dataSource=make([]interface{},0,1000) //开辟内存空间
	mystack.capsize=10  	//容量
	mystack.currentsize=0
}

func (mystack *Stack) Size() int{
	return mystack.currentsize
}

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

func (mystack *Stack)  Push(data interface{}){
	if !mystack.IsFull(){
		mystack.dataSource=append(mystack.dataSource,data)
		mystack.currentsize++
	}
}

func (mystack *Stack) IsFull() bool{
	return mystack.capsize==mystack.currentsize
}

func (mystack *Stack) IsEmpty() bool{
	return mystack.currentsize==0
}
