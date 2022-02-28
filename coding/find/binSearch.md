二分查找也是一种在数组中查找数据的算法。它只能查找已经排好序的数据。二分查找通过比较数组中间的数据与目标数据的大小，可以得知目标数据是在数组的左边还是右边。因此，比较一次就可以把查找范围缩小一半。重复执行该操作就可以找到目标数据，或得出目标数据不存在的结论。

在例如sql盲注之类的场景中，二分查找法可以很快的定位出数据。



代码实现：

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
```



二分查找法的关键点：
要查找的数据比中位数大，left则移到mid右边一位。

查找的数据比中位数小，right移到mid左边一位。

不断重复直到arr[mid]=要查找的数据data。

```go
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
```



Tips:

二分查找利用已排好序的数组，每一次查找都可以将查找范围减半。查找范围内只剩一个数据时查找结束。数据量为n的数组，将其长度减半log2n次后，其中便只剩一个数据了。也就是说，在二分查找中重复执行“将目标数据和数组中间的数据进行比较后将查找范围减半”的操作log2n次后，就能找到目标数据（若没找到则可以得出数据不存在的结论），因此它的时间复杂度为O（logn）。

补充说明:

二分查找的时间复杂度为O（logn），与线性查找的O（n）相比速度上得到了指数倍提高（x=log2n，则n=2x）。但是，二分查找必须建立在数据已经排好序的基础上才能使用，因此添加数据时必须加到合适的位置，这就需要额外耗费维护数组的时间。而使用线性查找时，数组中的数据可以是无序的，因此添加数据时也无须顾虑位置，直接把它加在末尾即可，不需要耗费时间。综上，具体使用哪种查找方法，可以根据查找和添加两个操作哪个更为频繁来决定。