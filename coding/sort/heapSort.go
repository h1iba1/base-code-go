package main

import "fmt"

func HeapSortMax(arr []int, length int) []int {
	//length := len(arr)
	if length <= 1 {
		return arr
	} else {
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

func main() {
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
