选择排序就是重复“从待排序的数据中寻找最小值，将其与序列最左边的数字进行交换”这一操作的算法。在序列中寻找最小值时使用的是线性查找。



代码实现：

```go
package main

import (
	"fmt"
	"strings"
)

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
		}
		return arr
	}
}

func main() {
	arr := []int{1, 2, 4, 7, 9, 0, 2, 3, 1, 4}
	fmt.Println(selectSort(arr))
}

```

算法关键点在于交换索引,先记录最小值索引往后进行比较，如果有更小的就交换。

```go
for i := 0; i < length-1; i++ {
			// 记录最小值的索引
			min := i
			// 从选择值后面值开始比较
			for j := i + 1; j < length; j++ {
				if arr[min] > arr[j] {
					min = j
				}
			}
			// 如果选中值不是最小值就交换索引
			if i != min {
				arr[i], arr[min] = arr[min], arr[i]
			}
		}
```

tips:

选择排序使用了线性查找来寻找最小值，因此在第1轮中需要比较n-1个数字，第2轮需要比较n-2个数字……到第n-1轮的时候就只需比较1个数字了。因此，总的比较次数与冒泡排序的相同，都是（n-1）+（n-2）+…+1≈n2/2次。每轮中交换数字的次数最多为1次。如果输入数据就是按从小到大的顺序排列的，便不需要进行任何交换。选择排序的时间复杂度也和冒泡排序的一样，都为O（n2）。



