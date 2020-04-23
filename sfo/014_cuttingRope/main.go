package main

import "fmt"

// 剪绳子
// 给一根长度为 n 的绳子，请把绳子剪成 m 段，（m，n 都是整数，m, n >1，剪出来的每段绳子长度也是整数）请问所有段绳子长度的乘积最大值是多少
// 思路：动态规划
func maxProductAfterCutting(length int) int {
	if length < 2 {
		return 0
	}
	if length == 2 {
		return 1
	}
	if length == 3 {
		return 2
	}
	products := make([]int, length+1)
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3
	for i := 4; i <= length; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			product := products[j] * products[i-j]
			if max < product {
				max = product
			}
		}
		products[i] = max
	}
	return products[length]
}

func main() {
	fmt.Println(maxProductAfterCutting(4))
}
