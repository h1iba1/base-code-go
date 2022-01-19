package main

import (
	"codeing/ArrayList"
	"codeing/Queue"
	"codeing/StackArray"
	"fmt"
)

func main1() {
	list:=ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append("字符串")
	list.Append("test")

	fmt.Println(list.DataStore)
}

func mainArrayList(){
	list:=ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(2)
	list.Append(2)
	list.Append(2)
	list.Append(2)
	list.Append(2)
	list.Append(2)
	list.Append("字符串")
	list.Append("test")

	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")

	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")
	list.Insert(2,"insert_2")

	list.Delete(1)
	fmt.Println(list.String())

	list.Clear()
	fmt.Println(list.DataStore)
}

func mainArrayListIterator()  {
	//list :=ArrayList.NewArrayList()
	var list ArrayList.List=ArrayList.NewArrayList()
	list.Append(1)
	list.Append(2)
	list.Append(2)
	list.Append(2)
	list.Append(2)
	list.Append(2)
	list.Append(2)
	list.Append(2)
	list.Append("字符串")
	list.Append("test")

	for it:=list.Iterator();it.HasNext();{
		value,_:=it.Next()
		fmt.Println(value)
		it.GetIndex()
		it.Remove()
	}

	fmt.Println(list.Iterator().HasNext())


}

func mainStackArray(){
	mystack :=StackArray.NewStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
}

func mainQueue()  {
	myq :=Queue.NewQueue()
	myq.EnQueue(1)
	myq.EnQueue(2)
	myq.EnQueue(3)
	myq.EnQueue(4)
	myq.EnQueue(5)
	fmt.Println(myq.Dequeue())
	fmt.Println(myq.Dequeue())
	fmt.Println(myq.Dequeue())
	fmt.Println(myq.Dequeue())
	fmt.Println(myq.Dequeue())
	fmt.Println(myq.Dequeue())
	myq.EnQueue(5)
	fmt.Println(myq.Dequeue())


}