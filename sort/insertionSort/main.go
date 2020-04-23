package main

import "fmt"

// 插入排序
// 主要思路是将数组分成两部分，一部分是已排序数组，一部分是未排序数组，然后将未排序数组中的元素，想插入扑克牌一样一个个插入到以排序数组的正确位置，
// 默认可以将第一个元素作为已排序数组的元素
func InsertionSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	// 外层循环每次遍历未排序数组中的元素
	for i := 1; i < len(array); i++ {
		tmp := array[i]
		j := i - 1
		// 内层循环遍历已排序数组中的元素，倒序遍历更容易写
		// 如果元素小于，将已排序部分的元素向后移一个位置，腾出来一个插入的空
		for ; j >= 0; j-- {
			if tmp < array[j] {
				// 注意，这里 j+1 可能覆盖掉 i 的位置，所以要把 i 的值先保存下来
				array[j+1] = array[j]
			} else {
				break
			}
		}
		// 注意循环结束的时候会执行一遍  j--， 所以空出来的位置在 j+1 上
		array[j+1] = tmp
	}
	return array
}

func main() {
	fmt.Println(InsertionSort([]int{3, 2, 1}))
}
