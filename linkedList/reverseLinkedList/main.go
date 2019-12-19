package main

import "fmt"

func main() {
	fmt.Println("stan")
}

type Node struct {
	value int
	next  *Node
}

// 反转整个链表，返回其 head 节点
// 主要思路：
// 用三个指针 prev current tmp 分别保存下面的状态
// prev 用来指向反转后的新链表的第一个元素
// current 用来指向原链表的反转节点的当前位置
// tmp 用来保存原链表在反转时后续元素的位置，防止链表断开
// 注意：
// 对于头节点，不仅要反转其位置，还要设置其 next 指向为 nil
// 主要考虑下面两种情况：链表为空或者链表只有一个节点（不需要反转），链表有大于等于 2 个节点
func reverseByLoop(head *Node) *Node {
	// 链表为空或者链表只有一个节点
	if head == nil || head.next == nil{
		return head
	}
	// 
	var prev *Node = nil
	current := head
	for current != nil {
		tmp := current.next
		current.next = prev
		prev = current
		current = tmp
	}
	return prev
}

func reverseByRecursion(head *Node) *Node {
	// 链表为空或者链表只有一个元素，不需要反转
	if head == nil || head.next == nil {
		return head
	}
	// 递归调用时，每次返回的值是反转后的链表头节点
	newHead := reverseByRecursion(head.next)
	// 此时 反转后的链表的末尾节点是 head.next, 所以还需要将 head 节点再挂在反转后的链表的末尾
	// 挂载方法是下面两步，注意挂载完要将 head 节点的 next 置为空
	head.next.next = head
	head.next = nil
	return newHead
}
