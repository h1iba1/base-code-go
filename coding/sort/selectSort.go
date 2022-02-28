package main

import (
	"fmt"
	"strings"
)

func selectSortMax(arr []int) int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i < length; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}

func selectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			// 记录最小值的索引
			min := i
			// 从选择值 后面值开始比较
			for j := i + 1; j < length; j++ {
				if arr[min] > arr[j] {
					min = j
				}
			}
			// 如果选中值不是最小值就交换索引
			if i != min {
				arr[i], arr[min] = arr[min], arr[i]
			}
			//fmt.Println(arr)
		}
		return arr
	}
}

func selectSortString(arr []string) []string {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			// 记录最小值的索引
			min := i
			// 从选择值 后面值开始比较
			for j := i + 1; j < length; j++ {
				if strings.Compare(arr[min], arr[j]) < 0 {
					min = j
				}
			}
			// 如果选中值不是最小值就交换索引
			if i != min {
				arr[i], arr[min] = arr[min], arr[i]
			}
			//fmt.Println(arr)
		}
		return arr
	}
}

func main4() {
	arr := []int{1, 2, 4, 7, 9, 0, 2, 3, 1, 4}
	arrStr := []string{"a", "v", "c", "b", "p", "s", "p", "z", "x"}
	fmt.Println(selectSortMax(arr))
	fmt.Println(selectSort(arr))

	fmt.Println(selectSortString(arrStr))
}
