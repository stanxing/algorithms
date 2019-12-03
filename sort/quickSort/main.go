package main

import "fmt"

func QuickSort(array []int) []int {
	quickSort(array, 0, len(array)-1)
	return array
}

func main() {
	fmt.Println(QuickSort([]int{3, 2, 1, 3, 4, 3}))
}

func quickSort(array []int, left, right int) {
	if left >= right {
		return
	}
	// 把第一个元素选为主元
	pivot := array[left]

	i, j := left+1, right

	for {
		for i <= j && array[i] <= pivot {
			i++
		}
		for i <= j && array[j] > pivot {
			j--
		}
		if i > j {
			break
		}
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}

	// 重点： 当循环结束的时候，j 指向的元素是最后一个小于等于中轴的元素，因为 i 会指向的位置，可能导致数组越界， 而 j 最小的值也会等于 left 即 pivot 的位置
	array[left], array[j] = array[j], array[left]
	quickSort(array, left, j-1)
	quickSort(array, j+1, right)
}
