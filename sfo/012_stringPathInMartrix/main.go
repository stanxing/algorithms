package main

import "fmt"

// 矩阵中的路径，请设计一个函数，用来判断在一个矩阵中是否存在一条包含某个字符串所有字符的路径。路径可以从矩阵中的任意一格开始，
// 每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵中的某一格，那么该路径不能再次进入该格子。
// 解决思路： 回溯法
// 举例, 在下面 4×3 的矩阵中寻找是否包含 bfce 的字符串：
//     a  "b"   t   g
//     c  "f"  "c"  s
//     j   d   "e"  h
//
// 回溯法的本质就是穷举，它列出了解决问题的每一步的所有可能，然后寻找到一个满足所有条件的解决方案。
// 所以回溯法非常适合由多个步骤组成，而且每个步骤有多个选项的这类问题。

// matrix 字符矩阵，用二维数组表示
// str 表示要搜索的字符串
func hasPath(matrix [][]byte, str string) bool {
	// 先判断列，再判断行
	if matrix == nil || len(matrix) < 1 || len(matrix[0]) < 1 || str == "" {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	// visited 用来记录哪些节点被访问过，可以初始化为一个二维数组一一对应。
	visited := make([][]bool, rows)
	for index := range visited {
		visited[index] = make([]bool, cols)
	}
	currentPathPostion := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if hasPathCore(matrix, rows, cols, i, j, &currentPathPostion, str, visited) {
				fmt.Println(visited)
				return true
			}
		}
	}
	return false
}

func hasPathCore(matrix [][]byte, rows, cols, row, col int, currentPathPostion *int, str string, visited [][]bool) bool {
	if *currentPathPostion >= len(str) {
		return true
	}
	hasPath := false
	// 满足下面的条件，说明本次查找符合要求
	if row >= 0 && row < rows && col >= 0 && col < cols && matrix[row][col] == str[*currentPathPostion] && !visited[row][col] {
		*currentPathPostion++
		visited[row][col] = true
		hasPath = hasPathCore(matrix, rows, cols, row, col-1, currentPathPostion, str, visited) ||
			hasPathCore(matrix, rows, cols, row, col+1, currentPathPostion, str, visited) ||
			hasPathCore(matrix, rows, cols, row-1, col, currentPathPostion, str, visited) ||
			hasPathCore(matrix, rows, cols, row+1, col, currentPathPostion, str, visited)
		if !hasPath {
			*currentPathPostion--
			visited[row][col] = false
		}
	}

	return hasPath
}

func main() {
	fmt.Println(hasPath([][]byte{[]byte{
		byte('a'), byte('b'), byte('t'), byte('g'),
	}, []byte{
		byte('c'), byte('f'), byte('c'), byte('s'),
	}, []byte{
		byte('j'), byte('d'), byte('e'), byte('h'),
	}}, "bfce"))
}
