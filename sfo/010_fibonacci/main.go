package main

// 求 fibonacci 数列的第 n 项
// fibonacci 数列的定义为
//     | 0 (n = 0)
// f = | 1 (n = 1)
//     | f(n-1) + f(n-2) (n > 1)
// 递归求法
// 缺点：用递归来求有严重的效率问题，会重复计算很多步的值，时间复杂度是以指数形式递增的。
func fibonacciRecusively(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacciRecusively(n-1) + fibonacciRecusively(n-2)
}

// 循环求法
// 可以避免重复计算，时间复杂度是 O(n)
func fibonacciIteratively(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	fiboNMinusOne := 1
	fiboNMinusTwo := 0
	fiboN := 0
	for i := 2; n <= n; i++ {
		fiboN = fiboNMinusOne + fiboNMinusTwo
		fiboNMinusOne = fiboN
		fiboNMinusTwo = fiboNMinusOne
	}
	return fiboN
}
