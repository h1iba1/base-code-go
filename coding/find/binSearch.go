package main

import "fmt"

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
			}
			if !isNeedExchange {
				break
			}
		}
		return arr
	}
}

func binSearch(arr []int, data int) int {
	left := 0
	right := len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] > data {
			right = mid - 1
		} else if arr[mid] < data {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	var arr = []int{999, 2, 1, 2, 5, 3, 8, 0, 4, 7, 0, 1, 1, 1, 1, 1000, 3, 1, 3, 6, 8, 9, 99, 33, 0}
	BubbleSort(arr)

	fmt.Println(binSearch(arr, 9))
}
