package main

// 二分查找
func BinarySearch(array []int, value int) int {
	// return 0
	return bSearchByLoop(array, value)
}

func bSearchByLoop(array []int, value int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := left + (right-left)/2
		if value < array[mid] {
			right = mid + 1
		} else if value > array[mid] {
			left = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func bSearchByRecursion(array []int, value int, left, right int) int {
	if left > right {
		return -1
	}
	return 0
	// mid := left + (right-left)/2
	// if array[mid] >
}

// 查找第一个值等于给定值的元素位置
func GetFirstEqualsValue(array []int, value int) int {
	return 0
}

// 查找最后一个值等于给定元值的元素位置
func GetLastEqualsValue(array []int, value int) int {
	return 0
}

// 查找第一个大于等于给定值的元素
func GetFirstGteValue(array []int, value int) int {
	return 0
}

// 查找最后一个小于等于给定值的元素
func GetLastLteValue(array []int, value int) int {
	return 0
}
