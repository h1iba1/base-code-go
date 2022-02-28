package main

import "fmt"

// 从小到大排序
//var arr = [...]int{1, 2, 1, 2, 5, 3, 8, 0, 4, 7}

func BubbleFindMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	} else {
		// 发现最大值并移到最右边
		for i := 0; i < length-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		return arr[length-1]
	}
}

func BubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		// 冒泡排序的精髓 两两比较
		for i := 0; i < length-1; i++ {
			// 可以大大提高运算效率
			isNeedExchange := false
			for j := 0; j < length-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					// 如果不再有arr左边的数大于右边的数。则退出
					isNeedExchange = true
				}
				fmt.Println(arr)
			}
			if !isNeedExchange {
				break
			}
			fmt.Println("------------------")
			fmt.Println(arr)
		}
		return arr
	}
}

func main1() {
	//for i := 0; i < len(arr); i++ {
	//	// 关键点在这里，j=i
	//	// 顺序挑选出一个数，去和他后面的所有数比较，找到一个比他小的就交换位置，确保每一轮去比较的数字一定是最小的
	//	// 第二次for循环的数一定是该次循环中最小的
	//	for j := i + 1; j < len(arr); j++ {
	//		// 比他大
	//		if arr[i] > arr[j] {
	//			arr[i], arr[j] = arr[j], arr[i]
	//		}
	//		fmt.Println(arr)
	//	}
	//	fmt.Println("---------------")
	//}
	var arr = []int{999, 2, 1, 2, 5, 3, 8, 0, 4, 7, 0, 1, 1, 1, 1, 1000, 3, 1, 3, 6, 8, 9, 99, 33, 0}

	//fmt.Println(BubbleFindMax(arr))
	fmt.Println(BubbleSort(arr))
}
