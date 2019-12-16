package main

type Node struct {
	value int
	next  *Node
}

// 链表中环的检测

// 主要思路是设置快慢两个指针，快指针每次移动两步，慢指针每次移动一步，如果有环的存在，快慢指针一定会相遇
func hasCycle(head *Node) bool {
	// 如果链表为空，或者只有一个节点，肯定没有环
	if head == nil || head.next == nil {
		return false
	}
	slow, fast := head, head
	// 注意当双指针的时候，由于快指针每次移动两个步长，为了防止报错，应该检查 fast 和 fast.next 两个变量，防止空指针异常
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}

	}
	return false
}
