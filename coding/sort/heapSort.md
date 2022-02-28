堆排序的特点是利用了数据结构中的堆。每次将堆的最大值提取出来。



代码实现如下：

```go
package main

import "fmt"

func HeapSortMax(arr []int, length int) []int {
	//length := len(arr)
	if length <= 1 {
		return arr
	} else {
		// 二叉树深度
		depth := length/2 - 1
		// 循环所有的三节点
		for i := depth; i >= 0; i-- {
			topMax := i //假定最大的值在i的位置
			leftChild := 2*i + 1
			rightChild := 2*i + 2 //左右孩子的节点

			if leftChild <= length-1 && arr[leftChild] > arr[topMax] { //防止越界
				topMax = leftChild //如果左边比我大 记录最大
			}
			if rightChild <= length-1 && arr[rightChild] > arr[topMax] {
				topMax = rightChild //如果右边比我大，记录最大
			}
			if topMax != i { //确保i的值最大
				arr[i], arr[topMax] = arr[topMax], arr[i]
			}
		}
		return arr
	}
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastmesslen := length - i //每次截取一段
		HeapSortMax(arr, lastmesslen)
		if i < length {
			arr[0], arr[lastmesslen-1] = arr[lastmesslen-1], arr[0]
		}
	}
	return arr
}

func main4() {
	var arr = []int{1, 999, 2, 2, 5, 3, 8, 0, 4, 7, 0, 1, 1, 2222, 9999, 1, 1, 1000, 3, 1, 3, 6, 8, 9, 99, 33, 0, 1001}

	fmt.Println(HeapSort(arr))
}

// [1000 1 33 999 99 2 8 2 4 8 5 3 1 1 1 0 3 1 3 6 7 9 0 1 0]
/*
1000
1	3
999	99		2 8
2 4	 8 5 	3 1	1 1
*/
```



算法关键点在于堆构造。

一些堆的基本概念：
堆深度=数组长度/2-1

左子节点=2*深度+1

右子节点=2*深度+2

```go
// 深度
		depth := length/2 - 1
		// 循环所有的三节点
		for i := depth; i >= 0; i-- {
			topMax := i //假定最大的值在i的位置
			leftChild := 2*i + 1
			rightChild := 2*i + 2 //左右孩子的节点

			if leftChild <= length-1 && arr[leftChild] > arr[topMax] { //防止越界
				topMax = leftChild //如果左边比我大 记录最大
			}
			if rightChild <= length-1 && arr[rightChild] > arr[topMax] {
				topMax = rightChild //如果右边比我大，记录最大
			}
			if topMax != i { //确保i的值最大
				arr[i], arr[topMax] = arr[topMax], arr[i]
			}
		}
```



Tips:

堆排序一开始需要将n个数据存进堆里，所需时间为O（nlogn）。排序过程中，堆从空堆的状态开始，逐渐被数据填满。由于堆的高度小于log2n，所以插入1个数据所需要的时间为O（logn）。每轮取出最大的数据并重构堆所需要的时间为O（logn）。由于总共有n轮，所以重构后排序的时间也是O（nlogn）。因此，整体来看堆排序的时间复杂度为O（nlogn）。这样来看，堆排序的运行时间比之前讲到的冒泡排序、选择排序、插入排序的时间O（n2）都要短，但由于要使用堆这个相对复杂的数据结构，所以实现起来也较为困难。