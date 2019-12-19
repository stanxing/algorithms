package main

import (
	"errors"
	"fmt"
)

// 用两个栈实现一个队列
// 主要思路：stack1 用作保存插入的元素，stack2 用作辅助数组
type Queue struct {
	// https://yourbasic.org/golang/implement-stack/
	// 用 slice 来模拟栈
	stack1  []int
	stack2  []int
	size    int
	current int
}

func newQueue(size int) *Queue {
	return &Queue{size: size}
}

// 队尾插入
func (q *Queue) appendTail(value int) error {
	if q.current == q.size {
		return errors.New("full queue")
	}
	q.stack1 = append(q.stack1, value)
	q.current++
	return nil
}

// 队头删除
func (q *Queue) deleteHead() error {
	if q.current == 0 {
		return errors.New("empty queue")
	}
	// 如果 stack2 的长度为 0,将 stack1 中的数据搬过来，由于栈的特性，stack2中的数据会倒过来存储，刚好满足先进先去的原则。
	if len(q.stack2) == 0 {
		for len(q.stack1) > 0 {
			nStack1 := len(q.stack1) - 1
			q.stack2 = append(q.stack2, q.stack1[nStack1])
			q.stack1 = q.stack1[:nStack1]
		}
	}
	nStack2 := len(q.stack2) - 1
	q.stack2 = q.stack2[:nStack2]
	q.current--
	return nil
}

func main() {
	queue := newQueue(10)
	queue.appendTail(1)
	queue.appendTail(2)
	queue.appendTail(3)
	queue.appendTail(4)
	queue.appendTail(5)
	queue.appendTail(6)
	queue.appendTail(7)
	queue.appendTail(8)
	queue.appendTail(9)
	queue.appendTail(10)

	queue.deleteHead()
	queue.deleteHead()

	queue.appendTail(12)
	fmt.Println(queue)
}
