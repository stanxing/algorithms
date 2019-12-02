package main

import "fmt"

// 冒泡排序
// 时间复杂度 O(n^2), 最好 O(n)
// 空间复杂度 O(1)
// 每次循环中通过比较相邻两个元素的大小，一步步将较大的元素通过交换移动到数组的一端，
// 对于 n 个数组的元素，外层循环需要 n-1 次（每次排一个元素，最后一次只剩下一个元素了，不用派）
// 内层循环需要从 0 开始，但是结束条件跟外层循环有关，因此每过一次外层循环，内层需要比较的元素个数就少了一个
// 优化： 当一次循环的冒泡过程中没有发生过一次排序交换，说明数组已经是有序的了，这时候就可以跳出循环
func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		isOrdered := true
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				isOrdered = false
			}
		}
		if isOrdered {
			break
		}
	}
	return array
}

func main() {
	fmt.Println(BubbleSort([]int{3, 2, 1}))
}
