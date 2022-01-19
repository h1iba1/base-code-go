package main

import "fmt"

func main()  {
	//link :=CreateLink()
	node :=&Node{
		data: 1,
		pNext: nil,
	}

	link :=&List{headNode: node}

	link.Append(1)
	link.Append(2)
	link.Append(3)
	link.Append(4)

	link.Remove(1)
	link.Remove(1)

	link.Insert(3,999)

	link.ShowList()

	fmt.Println("链表长度：",link.Length())
	fmt.Println("链表是否包含1",link.Contain(1))
}