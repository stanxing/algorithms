package main

import "fmt"

// 二维数组中的查找
// 在一个二维数组中，每一行都按照从左到右依次递增，每一列也按照从上到下依次递增。
// 请输入一个数，判断这个二维数组中是否存在这个数。
// 解题思路：
// 如果把这个二维数组想成一个矩阵，然后想着用二分法解题会陷入一个错误的思路，因为选中二维数组中的一个数和 value 相比，如果选的数比 value 大，说明 value 在这个数的左边或者上边，可是左边和上边有一部分重合的元素，会给比较带来复杂度。
// 正确的解题思路应该利用到从左到右递增和从上到下递增这个题目条件，说明这个二维数组的右上角是每一行的最大值，又是每一列的最小值（或者左下角也有这样的规律），如果用这个值去跟 value 比较的话，每次不管是比 value 大还是比 value 小，总能排除掉一行或者一列的元素。这样一直比较下去，就能找到 value 所对应的值。
// 一定不能用左上角或者右下角来比较，因为左上角既是行的最小值也是列的最小值，无法排除任何信息。
func findInTwoDimensionalArray(array [][]int, value int) bool {
	rowIndex := 0
	row := array[rowIndex]
	columnIndex := len(row) - 1
	for rowIndex < len(array) {
		row = array[rowIndex]
		for columnIndex >= 0 {
			// 拿右上角的元素跟其比较
			if row[columnIndex] > value {
				columnIndex--
			} else if row[columnIndex] < value {
				rowIndex++
				break
			} else {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(findInTwoDimensionalArray([][]int{
		[]int{1, 2, 7, 9},
		[]int{2, 4, 9, 12},
		[]int{4, 5, 10, 13},
		[]int{6, 8, 11, 15},
	}, 15))
}
