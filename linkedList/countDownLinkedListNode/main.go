package main

type Node struct {
	value int
	next  *Node
}

// 获取链表中的倒数第 n 个节点
// 思路是设置快慢两个指针 slow 和 fast，首先，两个指针都指向 head，然后 fast 向前走 n 步，这样，两个指针 slow 和 fast 之间就间隔 n 个步长，最后 slow 和 fast 同时移动，当 fast 到达最末端的时候，slow 就是倒数 n 个节点的位置

// 时间复杂度 O(n)
// 空间复杂度 O(1)
func countDownTheNthNode(head *Node, n int) *Node {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.next
	}
	if fast == nil {
		return fast.next
	}
	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	return slow
}

func main() {
	//
}
