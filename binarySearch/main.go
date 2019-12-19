package main

import "fmt"

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
			right = mid - 1
		} else if value > array[mid] {
			left = mid + 1
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
	mid := left + (right-left)/2
	if value == array[mid] {
		return mid
	} else if value < array[mid] {
		return bSearchByRecursion(array, value, left, mid-1)
	}
	return bSearchByRecursion(array, value, mid+1, right)
}

// 查找第一个值等于给定值的元素位置
func GetFirstEqualsValue(array []int, value int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := left + (right-left)/2
		if value < array[mid] {
			right = mid - 1
		} else if value > array[mid] {
			left = mid + 1
		} else {
			// 如果 mid 已经在最左边了或者 mid 的前一个元素不等于 value，说明找到了第一个元素的位置
			// 否则就继续二分
			if mid == 0 || array[mid-1] != value {
				return mid
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 查找最后一个值等于给定元值的元素位置
func GetLastEqualsValue(array []int, value int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := left + (right-left)/2
		if value < array[mid] {
			right = mid - 1
		} else if value > array[mid] {
			left = mid + 1
		} else {
			if mid == len(array)-1 || array[mid+1] != value {
				return mid
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

// 查找第一个大于等于给定值的元素
func GetFirstGteValue(array []int, value int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := left + (right-left)/2
		if value <= array[mid] {
			if mid == 0 || array[mid-1] < value {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// 查找最后一个小于等于给定值的元素
func GetLastLteValue(array []int, value int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := left + (right-left)/2
		if value >= array[mid] {
			if mid == len(array)-1 || array[mid+1] > value {
				return mid
			} else {
				left = mid + 1
			}
		} else {
			left = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println("Get value postion", BinarySearch([]int{1, 3, 5, 7, 9}, 7))
	fmt.Println("Get first value postion", GetFirstEqualsValue([]int{1, 3, 3, 5, 7, 9}, 3))
	fmt.Println("Get last value postion", GetLastEqualsValue([]int{1, 3, 3, 5, 7, 9}, 3))
	fmt.Println("Get first gte value postion", GetFirstGteValue([]int{1, 3, 3, 5, 7, 9}, 3))
	fmt.Println("Get last lte value postion", GetLastLteValue([]int{1, 3, 3, 5, 7, 9}, 3))
}
