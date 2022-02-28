冒泡排序就是重复“从序列右边开始比较相邻两个数字的大小，再根据结果交换两个数字的位置”这一操作的算法。在这个过程中，数字会像泡泡一样，慢慢从右往左“浮”到序列的顶端，所以这个算法才被称为“冒泡排序”。

go简单实现：

```go
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
				//fmt.Println(arr)
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

func main2() {
	var arr = []int{999, 2, 1, 2, 5, 3, 8, 0, 4, 7, 0, 1, 1, 1, 1, 1000, 3, 1, 3, 6, 8, 9, 99, 33, 0}

	fmt.Println(BubbleSort(arr))
}

```

冒泡排序的核心要点就是两两比较。通过不断的循环比较两边的值来生成一个稳定的序列。

```go
for i := 0; i < length-1; i++ {
			for j := 0; j < length-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
```

优化：在排序正确时，因为第一个for特性，还会继续循环。可以通过一些条件来进行判断,通过该方式可以大大的提高效率。

```go
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
```

当没有检查到左边比右边更大时，说明顺序正确，退出循环。



tips:

在冒泡排序中，第1轮需要比较n-1次，第2轮需要比较n-2次……第n-1轮需要比较1次。因此，总的比较次数为（n-1）+（n-2）+…+1≈n2/2。这个比较次数恒定为该数值，和输入数据的排列顺序无关。不过，交换数字的次数和输入数据的排列顺序有关。假设出现某种极端情况，如输入数据正好以从小到大的顺序排列，那么便不需要任何交换操作；反过来，输入数据要是以从大到小的顺序排列，那么每次比较数字后便都要进行交换。因此，冒泡排序的时间复杂度为O（n2）。

