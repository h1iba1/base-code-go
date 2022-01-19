package main

import "fmt"

type Node struct{
	data interface{}
	pNext *Node
}

type List struct {
	headNode *Node
}

type Link interface {
	// IsEmpty 判断链表是否为空
	IsEmpty() bool

	// Length 获取列表长度
	Length() int

	// Add 从链表头部添加元素
	Add(data interface{}) *Node

	// Append 从链表尾部添加元素
	Append(data interface{})

	// Insert 链表指定位置添加元素
	Insert(index int, data interface{})

	// Remove 删除链表指定值的元素
	Remove(data interface{})

	// RemoveAtIndex 删除链表指定位置的元素
	RemoveAtIndex(index int)

	// Contain 查看链表是否包含某个元素
	Contain(data interface{}) bool

	// ShowList 遍历所有节点
	ShowList()
}

func CreateLink() *List {
	node :=&Node{data: nil}
	link:=&List{headNode: node}
	return link
}

// IsEmpty 判断链表是否为空
func (list *List) IsEmpty() bool{
	return list.headNode==nil
}

// Length 获取列表长度
func (list *List) Length() int{
	// 获取链表头部
	headNode := list.headNode
	// 计数器
	count:=0

	// 如果头节点不为空则count++
	for headNode!=nil{
		count++
		headNode=headNode.pNext
	}
	return count
}

// Add 从链表头部添加元素
func (list *List) Add(data interface{}) *Node{
	// 添加新节点，并初始化data
	node :=&Node{data: data}
	// 将新节点的下一个节点指向头节点
	node.pNext=list.headNode

	// 更新头部节点
	list.headNode=node
	return node
}

// Append 从链表尾部添加元素
func (list *List) Append(data interface{}){
	node :=&Node{data: data}
	// 链表为空，则添加到头节点
	if list.IsEmpty(){
		list.headNode=node
	}else{
		cur :=list.headNode
		// 如果节点为nil说明为最后一个节点
		// 不为nil，则循环。为空nil退出，此时cur等于最后一个节点
		for cur.pNext !=nil{
			cur=cur.pNext
		}
		// 最后一个节点指向新创建的节点
		cur.pNext=node
	}
}

// Insert 链表指定位置添加元素
func (list *List) Insert(index int, data interface{}){
	if index<0{
		list.Add(data)
	}
	if index>list.Length(){
		list.Append(data)
	}
	pre :=list.headNode
	count:=0
	// 寻址找到需要插入的index
	for count < (index-1){
		// 此时的pre=第index-1个node
		pre=pre.pNext
		count++
	}
	node :=&Node{data: data}
	//新创建节点指向第index个node的头部
	node.pNext=pre.pNext
	// 覆盖第index个值
	pre.pNext=node

}

// Remove 删除链表指定值的元素
func (list *List) Remove(data interface{}){
	pre :=list.headNode
	length :=list.Length()
	// data是头节点则直接删除
	if pre.data==data{
		list.headNode=pre.pNext
	}
	// 遍历链表，直到链表最后一个节点
	for pre.pNext!=nil{
		// 寻找到data，则将指针后移，达到删除的目的
		if pre.pNext.data==data{
			pre.pNext=pre.pNext.pNext
		}else{
			// 向下寻址
			pre=pre.pNext
		}
	}

	// 长度没变化，说明链表中没有data
	if length==list.Length(){
		fmt.Println("找不到data：",data)
	}
}

// RemoveAtIndex 删除链表指定位置的元素
func (list *List) RemoveAtIndex(index int){
	pre :=list.headNode
	//
	if index<0{
		list.headNode=pre.pNext
	}

	if index>list.Length(){
		fmt.Println("超出链表长度")
		return
	}
	// 计数器
	count :=0
	// 遍历链表，直到count=index-1 或 下一个节点为空
	for count !=(index-1) && pre.pNext !=nil{
		count++
		// 第index-1个节点
		pre=pre.pNext
	}
	//第index个节点用index+1个节点覆盖
	pre.pNext=pre.pNext.pNext
}


// Contain 查看链表是否包含某个元素
func (list *List) Contain(data interface{}) bool{
	cur:=list.headNode
	// 遍历链表
	for cur!=nil{
		if cur.data==data{
			return true
		}
		cur=cur.pNext
	}
	return false
}

// ShowList 遍历所有节点
func (list *List) ShowList(){

	if list.IsEmpty(){
		fmt.Println("链表为空")
	}
	cur :=list.headNode
	for cur!=nil{
		fmt.Printf("\t%v\n", cur.data)
		cur=cur.pNext
	}
}