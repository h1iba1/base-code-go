```go
package Queue

type MyQueue interface{
	Size() int				//大小
	Front() interface{}		// 第一个元素
	End() interface{}		// 最后一个元素
	IsEmpty() bool			// 是否为空
	EnQueue(data interface{})	// 入队
	Dequeue() interface{}	//出队
	Clear()					//清空
}

type Queue struct {
	dataStore []interface{}	// 队列的数据存储
	theSize int				//队列的大小
}

func NewQueue()*Queue{
	myqueue :=new(Queue)	//开辟结构体
	myqueue.dataStore=make([]interface{},0)
	myqueue.theSize=0
	return myqueue
}

func (myq *Queue)Size() int{
	return myq.theSize
}

func (myq *Queue)Front() interface{}{
	if myq.Size()==0{
		return nil
	}
	return myq.dataStore[0]
}

func (myq *Queue) End() interface{}{
	if myq.Size()==0{
		return nil
	}
	return myq.dataStore[myq.Size()-1]
}

func (myq *Queue) IsEmpty() bool{
	return myq.theSize==0
}

func (myq *Queue) EnQueue(data interface{}){
	myq.dataStore=append(myq.dataStore,data)
	myq.theSize++
}

func (myq *Queue) Dequeue() interface{}{
	if myq.Size()==0{
		return nil
	}
	data := myq.dataStore[0]
	if myq.Size()>1{
		myq.dataStore=myq.dataStore[1:myq.Size()]
	}else{
		myq.dataStore=make([]interface{},0)
	}
	myq.theSize--
	return data
}


func (myq *Queue) Clear(){
	myq.dataStore=make([]interface{},0)
	myq.theSize=0
}

```