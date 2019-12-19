package main

import "fmt"

// https://cloud.tencent.com/developer/article/1520494
// 计数排序
// 非基于比较的排序算法，主要思路是利用数组下标来确定元素的具体位置
// 适合的场景：
// 待排序的元素是非负整数，如果包含负数，需要先转换成非负整数来比较（先找到最小的负数，然后所有的数都加上该负数的绝对值，使得相对大小不变）
// 待排序数组中最大值和最小值的差距不能太大，否则就需要创建非常大的数组
// 当数列元素不是整数的时候，不适合用计数排序，因为元素下标无法表示为浮点数或者英文字符

// 时间空间复杂度：
// 遍历原始数组为 N，遍历计数数组为 M，总共用时 3N+M，即 O(N)
// 空间复杂度为 O(M+N) 计数数组和结果数组

// 该实现下的计数排序是不稳定算法，因为是统计元素下标的思路，无法去判断相同元素之间的顺序
func CountingSort(array []int) []int {
	// 数组中存放下标对应的值在原数组中出现的次数
	// max := getMax(array)
	// new := make([]int, max+1)

	// 上面使用最大值来创建计数数组的方式不够好，因为如果最大值很大，而最小值跟最大值差距不大，就会导致严重的空间浪费
	// 更好的思路是根据最大值和最小值的差距来创建统计数组
	max := getMax(array)
	min := getMin(array)
	interval := max - min + 1
	// 创建计数数组
	countArray := make([]int, interval)
	for _, e := range array {
		countArray[e-min]++
	}
	// 遍历统计数组，按顺序输出元素
	var sortedArray []int
	for i, ne := range countArray {
		for j := 0; j < ne; j++ {
			sortedArray = append(sortedArray, i+min)
		}
	}
	return sortedArray
}

// 改进计数排序，使其变成稳定的排序算法
func StableCountingSort(array []int) []int {
	min := getMin(array)
	max := getMax(array)
	interval := max - min
	// 创建计数数组
	countArray := make([]int, interval+1)
	for _, e := range array {
		countArray[e-min]++
	}
	// 对统计数组变形，后面的元素等于前面的元素之和\
	// 这一步做完以后，计数数组中存的值即为下标所对应的值的最后一个元素（如果有相同的元素）对应的位置
	sum := 0
	for i, e := range countArray {
		sum += e
		countArray[i] = sum
	}
	sortedArray := make([]int, len(array))
	// 输出排好序的数组的时候，倒序遍历原始数组，每次打印一个元素，原计数数组中元素的值 - 1
	for i := len(array) - 1; i >= 0; i-- {
		sortedArray[countArray[array[i]-min]-1] = array[i]
		countArray[array[i]-min]--
	}
	return sortedArray
}

func getMax(array []int) int {
	max := array[0]
	for _, e := range array {
		if e > max {
			max = e
		}
	}
	return max
}

func getMin(array []int) int {
	min := array[0]
	for _, e := range array {
		if e < min {
			min = e
		}
	}
	return min
}

func main() {
	fmt.Println("Counting sort: ", CountingSort([]int{4, 4, 6, 5, 3, 2, 8, 1, 7, 5, 6, 0, 10}))
	fmt.Println("Stable counting sort: ", StableCountingSort([]int{4, 4, 6, 5, 3, 2, 8, 1, 7, 5, 6, 0, 10}))
	// 90 99 95 94 95
}
