package main

import "fmt"

// https://blog.csdn.net/bjweimengshu/article/details/102908292
// 基数排序
// 用来解决复杂元素的排序，比如计数排序中差值过大不容易创建计数数组的场景或者非整数排序（比如字符串）
// 主要思路是先对数据进行位数补齐，然后对每一位分别进行一次计数排序，一共循环 K 次，K 位数据的位数长度
// 对于每一位的比较，创建一个 ASCII 码大小的数组，即可涵盖到所有的字符

// 复杂度分析：
// 时间复杂度是 O(K * (m+n))，因为要进行 K 次计数排序
// 空间复杂度是 O(m+n)，需要创建计数排序专用的计数数组，该数组可以重用

var AsciiSize = 128

func RadixSort(array []string) []string {
	// 排序数组，用来存储每一次按位排序后的临时结果
	sortedArray := make([]string, len(array))
	maxLength := getMaxLength(array)
	for i := maxLength - 1; i >= 0; i-- {
		// 执行计数排序：
		// 1.创建辅助排序的计数数组
		// 2.对统计数组变形，后面的元素等于前面元素的和
		// 3.倒序遍历原始数组，从统计数组中读取到正确的位置，插入到 sortedArray 中
		countArray := make([]int, AsciiSize)
		for j := 0; j < len(array); j++ {
			// 将字符转换成 ascii 码
			asciiIndex := getAscii(array[j], i)
			countArray[asciiIndex]++
		}
		sum := 0
		for j, e := range countArray {
			sum += e
			countArray[j] = sum
		}

		for j := len(array) - 1; j >= 0; j-- {
			asciiIndex := getAscii(array[j], i)
			sortedArray[countArray[asciiIndex]-1] = array[j]
			countArray[asciiIndex]--
		}
		copy(array, sortedArray)
	}
	return array
}

func getMaxLength(array []string) int {
	maxLength := len(array[0])
	for _, e := range array {
		if len(e) > maxLength {
			maxLength = len(e)
		}
	}
	return maxLength
}

func getAscii(str string, position int) int {
	// 注意这里，如果 str 上对应的位置不存在字符，应该为其补上 0（ascii 码 48）
	if len(str)-1 < position {
		return 48 // 0 的 ascii 码
	}
	return int(str[position])
}

func main() {
	fmt.Println(RadixSort([]string{"qw", "abc", "qwe", "hhh", "a", "cws", "ope"}))
	fmt.Println(RadixSort([]string{"qwew", "abc", "qwe", "hhh", "a", "cws", "ope"}))
}
