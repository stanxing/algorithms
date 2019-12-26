package main

import "fmt"

func main() {
	// 二进制表示有符号整数
	// https://stackoverflow.com/questions/37582550/golang-twos-complement-and-fmt-printf
	// golang 无法直接输入补码，不过可以将正整数原码通过 ^ 按位取反来计算补码。
	// var binaryNumber int8 = ^0b0000000 // -1 的补码
	// var binaryNumber2 int8 = 0b1111111 // 127
	// var binaryNumber3 int8 = 0b1000101 // 68
	// var and int8 = 0b0000101 & 0b0000100
	// var or int8 = 0b0000001 | 0b0000101
	// var xor int8 = 0b0000001 ^ 0b1000101 // 0b1000100
	// fmt.Printf("binary: %#b, decimal: %d\n", binaryNumber, binaryNumber)
	// fmt.Printf("binary: %#b, decimal: %d\n", binaryNumber2, binaryNumber2)
	// fmt.Printf("binary: %#b, decimal: %d\n", binaryNumber3, binaryNumber3)
	// fmt.Printf("binary: %#b, decimal: %d\n", and, and)
	// fmt.Printf("binary: %#b, decimal: %d\n", or, or)
	// fmt.Printf("binary: %#b, decimal: %d\n", xor, xor)

	// 整数溢出测试
	// 1 左移 8 位是 8个 0，再减 1 得到 8 个 1；-1 相当于 +(-1), -1 的算法是取反 +1,得到的是 8 个 1。
	// 通过这种方式可以求得一个无符号数的最大值
	var n1 uint8 = 1<<8 - 1
	// n2 = -1，即将一个无符号数强制转成一个有符号数，会将无符号数的第一位作为符号位
	var n2 = int8(n1)
	// 获取一个有符号数的最小值和最大值
	var n3 int8 = int8(^uint8(0) >> 1) // 最大值，根据补码，最大值的二进制是：首位为0,其余位为1
	// golang 的取反符号适用于任何有符号还是无符号数
	var n4 int8 = ^n3 // 最小值，根据补码，最小值的二进制是：首位为1,其余为0
	// var n3 int8 = int8(1 << 7)  这个为什么错了
	fmt.Printf("result %b, %d\n", n1, n1)
	fmt.Printf("result %b, %d\n", n2, n2)
	fmt.Printf("result %b, %d\n", n3, n3)
	fmt.Printf("result %b, %d\n", n4, n4)
}
