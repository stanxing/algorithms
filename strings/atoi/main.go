package main

import (
	"fmt"
)

// 将字符串转换成整数
func Atoi(str string) (int, error) {
	// a, b := strconv.Atoi("1(3")
	a := byte('/')
	b := '0'
	c := a - '0' // 减法运算的本质是先将被减数的二进制取反+1后再做加法，两个无符号数的减法永远不会是负数
	var d uint8 = 0b11111111
	var e int8 = -0b1111111
	fmt.Println(a, b, c, d, e)
	fmt.Printf("%b  %b %b", c, d, e)
	//
	return 0, nil
}

func main() {
	Atoi("")
}

// 00000000 - 00000001
