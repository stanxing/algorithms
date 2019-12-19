package main

import "fmt"

func BucketSort(array []int, interval int) []int {
	if len(array) == 0 || len(array) == 1 {
		return array
	}
	min, max := array[0], array[0]
	// 先找到 array 中的最小值和最大值，用来计算桶的个数
	for _, a := range array {
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	// 根据差值和传递的间隔来计算需要多少个桶
	nBuckets := (max-min)/interval + 1
	buckets := make([][]int, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}
	// 注意计算元素应该放在哪个桶的方式
	for _, e := range array {
		idx := (e - min) / interval
		buckets[idx] = append(buckets[idx], e)
	}
	sorted := make([]int, 0)
	for _, bucket := range buckets {
		insertionSort(bucket)
		fmt.Println(min, max, nBuckets, buckets)
		sorted = append(sorted, bucket...)
	}
	return sorted
}

func insertionSort(subArray []int) []int {
	if len(subArray) <= 1 {
		return subArray
	}
	for i := 1; i < len(subArray); i++ {
		tmp := subArray[i]
		j := i - 1
		for ; j >= 0; j-- {
			if tmp < subArray[j] {
				subArray[j+1] = subArray[j]
			} else {
				break
			}
		}
		subArray[j+1] = tmp
	}
	return subArray
}

func main() {
	fmt.Println("Bucket sort: ", BucketSort([]int{4, 1, 1, 3, 2, 2, 7, 0}, 2))
	fmt.Println("Insertion sort: ", insertionSort([]int{4, 1, 1, 3, 2, 2, 7, 0}))
}
