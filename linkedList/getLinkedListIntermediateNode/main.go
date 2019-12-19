package main

func main() {
	//
}

type Node struct {
	value int
	next  *Node
}

// 获取链表中间的节点，要求只扫描一次链表,如果链表的长度是偶数，返回中间两个节点的任意一个，如果是奇数，返回中间节点

// 主要思路是设置快慢两个指针，快指针一次走两步，慢指针一次走一步，当快指针到达位置的时候，慢指针所指向的位置即为链表的中间节点
// 注意区分节点个数是奇数和偶数这两种情况：
// 奇数：返回中间节点
// 偶数：返回中间结点的下一位

func getIntermediateNode(head *Node) *Node {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast == nil || fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}
