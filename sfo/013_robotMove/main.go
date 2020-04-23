package main

import "fmt"

func main() {
	fmt.Println(movingCount(10, 38, 38))
}

// 机器人的运动范围
// 思路： 回溯法
// 地上有一个 m 行 n 行的方格。一个机器人从坐标（0,0）开始移动，每次可以向左、右、上、下一个方向移动一个格子，但不能进入行坐标和列坐标的数位之和大于 k 的方格。例如，k = 18时，机器人不能进入方格（35,38），因为 3 + 5 + 3 + 8 = 19。请问机器人能够走多少个格子。
func movingCount(threshold int, rows, cols int) int {
	if threshold < 0 || rows <= 0 || cols <= 0 {
		return 0
	}
	visited := make([][]bool, rows)
	for index := range visited {
		visited[index] = make([]bool, cols)
	}
	count := movingCountCore(threshold, rows, cols, 0, 0, visited)
	return count
}

func movingCountCore(threshold, rows, cols, row, col int, visited [][]bool) int {
	count := 0
	if check(threshold, rows, cols, row, col, visited) {
		visited[row][col] = true
		count = 1 + movingCountCore(threshold, rows, cols, row, col-1, visited) +
			movingCountCore(threshold, rows, cols, row, col+1, visited) +
			movingCountCore(threshold, rows, cols, row-1, col, visited) +
			movingCountCore(threshold, rows, cols, row+1, col, visited)
	}
	return count
}

func check(threshold, rows, cols, row, col int, visited [][]bool) bool {
	if row >= 0 && row < rows && col >= 0 && col < cols &&
		getDigitSum(row)+getDigitSum(col) <= threshold && !visited[row][col] {
		return true
	}
	return false
}

func getDigitSum(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number = number / 10
	}
	return sum
}
