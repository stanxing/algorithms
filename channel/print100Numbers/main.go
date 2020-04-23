package main

import "fmt"

// 两个协程交替打印0-100的奇数偶数
// 要求是奇数协程打印奇数，偶数协程打印偶数

// 方法1， 使用 1 个 channel 来控制
// 半开模式
func Method1() {
	c := make(chan int)
	stop := make(chan bool)
	num := 100
	// 先会被阻塞的一定是打印奇数的协程，因为第一个数是 0
	go func() {
		for i := 0; i <= num; i++ {
			c <- 1
			if i%2 == 1 {
				fmt.Print(i)
			}
		}
	}()

	go func() {
		for i := 0; i <= num; i++ {
			<-c
			if i%2 == 0 {
				fmt.Print(i)
			}
			if i == 100 {
				close(stop)
			}
		}
	}()

	<-stop
}

func main() {

}
