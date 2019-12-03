package main

import "fmt"

// 主要思路是分治法的思想，让左右两部分的元素先有序，然后将两个有序的部分合并成为一个有序的数组
func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	mergeSort(array, 0, len(array)-1)
	return array
}

func mergeSort(array []int, left, right int) {
	if left == right {
		return
	}
	mid := left + (right-left)/2

	mergeSort(array, left, mid)
	mergeSort(array, mid+1, right)
	merge(array, left, mid, right)
}

func merge(array []int, left, mid, right int) {
	tmpLength := right - left + 1
	tmp := make([]int, tmpLength)
	leftStart := left
	rightStart := mid + 1
	tmpStart := 0
	for ; leftStart <= mid && rightStart <= right; tmpStart++ {
		if array[leftStart] <= array[rightStart] {
			tmp[tmpStart] = array[leftStart]
			leftStart++
		} else {
			tmp[tmpStart] = array[rightStart]
			rightStart++
		}
	}

	for ; leftStart <= mid; leftStart++ {
		tmp[tmpStart] = array[leftStart]
		tmpStart++
	}
	for ; rightStart <= right; rightStart++ {
		tmp[tmpStart] = array[rightStart]
		tmpStart++
	}
	// 将 tmp 部分的数据合并到 array 中
	for i := 0; i < tmpLength; i++ {
		array[left+i] = tmp[i]
	}
}

func main() {
	fmt.Println(MergeSort([]int{3, 2, 1}))
}
