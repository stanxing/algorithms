// https://progolang.com/how-to-implement-singleton-pattern-in-go/
package singleton

import "sync"

type singleton struct{}

var instance *singleton
var mux sync.Mutex
var once sync.Once

// 简易版
// 存在的问题：
// 无法适应多线程编程，如多个 gorutine 同时访问到 nil 的位置
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// 改进版，使用 mutex 包, 每次在访问 if 之前加锁，保证只有在得到锁的线程里才能访问到
// 存在的问题：
// 在每次调用该访问的时候都要先去获取锁，即使在这个单例已经创建出来后，实际上完全不需要这么频繁的使用锁
func GetInstance1() *singleton {
	mux.Lock()
	defer mux.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// 改进版2，只有在检查到 instance 为 nil 的时候需要考虑原子性，当单例已经被创建出来，不再需要锁的存在，
// 因此在使用锁之前应该检查一下 instance 是否为 nil，即 check=>lock=>check
// 存在的问题：
// if 查询条件不是原子性的检查，仍然有概率存在失败。即一个 goroutine 执行完第一个 if 后进入抢锁环节，cpu 切到了另一个 goroutine 并创建好了 instance。这样还是会多检查一次锁。
// 其次，check=>lock=>check 使得代码复杂性提高。
func GetInstance2() *singleton {
	if instance != nil {
		return instance
	}
	mux.Lock()
	defer mux.Unlock()
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// 改进版3，使用 sync package Once.Do 方法，这个方法会确保指定的代码只执行一次
// 这个方法看起来代码既整洁又比较好理解
func GetInstance3() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

// https://draveness.me/golang/concurrency/golang-sync-primitives.html
// 对于 sync 包 https://golang.org/src/sync/once.go?s=1137:1164#L25 once.Do 方法的介绍，
// 从其代码可以看出，其实它用的思路也是 check => lock => check 的方式
// 只不过最外层的 if 判断使用的是原子操作 // if atomic.LoadUint32(&o.done) == 0，这样就将 if 判断这一步的原子性也完美的解决了。
// Once 的结构体中设置了 done 这么一个 uint32 的标志位，用原子操作去检查这个标志位是否为 0，如果是就执行函数部分。
