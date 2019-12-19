package main

import "fmt"

func main() {
	// 二进制表示有符号整数
	// https://stackoverflow.com/questions/37582550/golang-twos-complement-and-fmt-printf
	// golang 无法直接输入补码，不过可以将正整数原码通过 ^ 按位取反来计算补码。
	var binaryNumber int8 = ^0b0000000 // -1 的补码
	var binaryNumber2 int8 = 0b1111111 // 127
	var binaryNumber3 int8 = 0b1000101 // 68
	var and int8 = 0b0000101 & 0b0000100
	var or int8 = 0b0000001 | 0b0000101
	var xor int8 = 0b0000001 ^ 0b1000101 // 0b1000100
	fmt.Printf("binary: %#b, decimal: %d\n", binaryNumber, binaryNumber)
	fmt.Printf("binary: %#b, decimal: %d\n", binaryNumber2, binaryNumber2)
	fmt.Printf("binary: %#b, decimal: %d\n", binaryNumber3, binaryNumber3)
	fmt.Printf("binary: %#b, decimal: %d\n", and, and)
	fmt.Printf("binary: %#b, decimal: %d\n", or, or)
	fmt.Printf("binary: %#b, decimal: %d\n", xor, xor)
}
