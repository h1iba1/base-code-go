package main

import "fmt"

// 当文件很多 内存装不下。可以从文件一个个选择再进行插入排序

func InsertTest(arr []int) []int {
	backup := arr[2]
	// 从上一个位置循环找到位置插入
	j := 2 - 1
	for j >= 0 && backup < arr[j] {
		// 从前往后移
		arr[j+1] = arr[j]
		j--
	}
	arr[j+1] = backup
	return arr
}

// InsertSort 那要是14个数，按你这个公式是6层，但是实际是4层，所以公式该是以2为底取对数，再取向上取整

func InsertSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 1; i < length; i++ {
			backup := arr[i]
			j := i - 1
			// 数据从前往后移 直接覆盖
			for j >= 0 && backup < arr[j] {
				arr[j+1] = arr[j]
				j--
			}
			// 将备份插入
			arr[j+1] = backup
		}
	}
	return arr
}

func main3() {
	var arr = []int{1, 999, 2, 2, 5, 3, 8, 0, 4, 7, 0, 1, 1, 1, 1, 1000, 3, 1, 3, 6, 8, 9, 99, 33, 0}

	fmt.Println(InsertTest(arr))

	fmt.Println(InsertSort(arr))

}
