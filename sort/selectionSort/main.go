package main

import "fmt"

// 选择排序
// 思路和插入排序类似，也分为已排序区间和未排序区间，不过这次不是将未排序元素插入到已排序区间，
// 而是遍历未排序元素，找到最小的元素，将元素放入到已排序区间
func SelectionSort(array []int) []int {
	// 循环 n-1 次，因为在这个思路下当 n-1 个数排好序后，最后一个数肯定排好了
	// 如果这里不是 i < len(array)-1 的话，j 也不能写成 j = i+1, 只能写成 i，否则数组越界
	for i := 0; i < len(array)-1; i++ {
		min := i
		// 未排序数组可以从 i+1 开始找，反正会和 array[i] 比较，减少一次循环
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		// 如果这两个数不想等，说明选择到了更小的元素，应该交换位置
		if min != i {
			array[i], array[min] = array[min], array[i]
		}
	}
	return array
}

func main() {
	fmt.Println(SelectionSort([]int{3, 2, 1}))
}
