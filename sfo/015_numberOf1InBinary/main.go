package main

import "fmt"

func main() {
	dig := -1
	// result := dig << 63
	// fmt.Println(numberOfOneByRightMove(dig))
	fmt.Println(numberOfOneByLeftMoveFlag(dig))
	// fmt.Printf("binary: %b, %d\n", result, result)
}

// 二进制中 1 的个数
// 输入一个整数，输出该二进制表示中 1 的个数
// 主要思路： 位运算

// 思路1：将整数与 1 做与运算，如果结果是 1,说明末位是1,否则末位为 0。
// 缺陷：输入一个负数的话会进入死循环,因为
func numberOfOneByRightMove(n int) int {
	count := 0
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}

// 解决方案：为了负数也能够位运算，不能使用右移。首先把 n 和 1 做与运算，判断 n 的最低位是不是 1，接着把 1 左移 1位，再和 n 做与运算，可以得到第2位是不是 1，这样一直将 1 左移下去，每次都能判断 n 中的一位。因为 n 是一个 int 类型的值，所以所以最多 32 次就比较了 n 中所有的位数，而且左移的话因为左边补 0,1 最后一直左移得到的值一定会是 0

func numberOfOneByLeftMoveFlag(n int) int {
	count := 0
	flag := 1
	for flag != 0 {
		if n&flag != 0 {
			count++
		}
		flag = flag << 1
	}
	return count
}
