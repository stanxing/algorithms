package main

import "fmt"

// 主要思路是分治法的思想，让左右两部分的元素先有序，然后将两个有序的部分合并成为一个有序的数组
func MergeSort(array []int) []int {
	return merge(array, 0, len(array)-1)
}

func merge(array []int, left, right int) []int {
	if left == right {
		return array
	}
	for left < right
	// return array
}

func main() {
	fmt.Println(MergeSort([]int{3, 2, 1}))
}
