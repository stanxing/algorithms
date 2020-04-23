package main

import "fmt"

// 多线程交替打印字符
// 建立m个线程，每个线程只能打印一种字符，要求线程同时运行，交替打印n次字符。比如: n=3 m=2打印字符为A和B。要求线程1打印3次A，线程2打印3次B，在屏幕输出ABABAB

// 思路是利用无缓冲 channel 来做同步，main 函数通知 g0, g0 打印完之后通知 g1, 以此类推，最后一个协程 gn 再回来通知 g0
// 由于使用无缓冲 channel，整个过程都是同步按顺序进行的，直到 g0 打印的时候发现已经超过了打印次数，说明整个字符串已经打印完成，直接退出主程序即可。
func main() {
	m := 2 // 建立 m 个线程
	n := 3 // 打印三次

	g := make([]chan int, m)
	c := make([]string, m)
	stop := make(chan bool)

	for i := 0; i < m; i++ {
		g[i] = make(chan int)
		c[i] = string(65 + i)
	}
	// 开启 m 个线程
	for i := 0; i < m; i++ {
		i := i
		go func() {
			count := 0
			// 从 g[i] 接收到数据后才有资格打印
			for _ = range g[i] {
				if count < n {
					fmt.Printf(c[i])
					count++
					// 如果条件成立，说明是最后一个 g， 应该通知 g0, 否则，应该通知下一个 g 打印
					if i == m-1 {
						g[0] <- 1
					} else {
						g[i+1] <- 1
					}
				} else {
					// 这个时候说明所有 channel 都打印完 n 次了，可以直接关闭掉 stop 让程序退出
					close(stop)
				}
			}
		}()
	}

	g[0] <- 1
	<-stop
}
