package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(power(1.5, -1))
}

// 数值的整数次方，实现一个函数，求 base 的 exponent 次方（exponent 不会是小数），不得使用库函数，不需要考虑大数问题

// 第一种思路：for 循环遍历 exponent 次
// 注意： 输入的是负数该怎么办？
// 当输入的指数是负数时，可以先对指数求绝对值，然后for 循环计算结果，最后对计算出的结果求倒数（注意对0求倒数会抛异常）
func power(base float64, exponent int) (float64, error) {
	var sum float64 = 1
	if exponent == 0 {
		// 0 的 0 次方也等于 1
		return 1, nil
	}
	var absExponent uint
	if exponent < 0 {
		absExponent = uint(-exponent)
	} else {
		absExponent = uint(exponent)
	}
	for i := 0; uint(i) < absExponent; i++ {
		sum = sum * base
	}
	if exponent < 0 {
		if sum == 0 {
			return 0, errors.New("invaild 0")
		}
		return 1 / sum, nil
	}
	return sum, nil
}

// 更优的算法
//       | a^(n/2) * a^(n/2)，当 n 为偶数
// a^n = |
//       | a^((n-1)/2) * a^((n-1)/2) * a，当 n 为奇数

func power2(base float64, exponent int) (float64, error) {
	if exponent == 0 {
		// 0 的 0 次方也等于 1
		return 1, nil
	}
	var absExponent uint
	if exponent < 0 {
		absExponent = uint(-exponent)
	} else {
		absExponent = uint(exponent)
	}
	sum := powerWithUnsignedExponent(base, absExponent)
	if exponent < 0 {
		if sum == 0 {
			return 0, errors.New("invaild 0")
		}
		return 1 / sum, nil
	}
	return sum, nil
}

func powerWithUnsignedExponent(base float64, exponent uint) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	result := powerWithUnsignedExponent(base, exponent>>1)
	result = result * result
	if exponent&1 == 1 {
		result = result * base
	}
	return result
}
