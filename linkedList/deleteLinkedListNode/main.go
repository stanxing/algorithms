package main

import "fmt"

func main() {
	fmt.Println("123")
}

type Node struct {
	value int
	next  *Node
}

// 在 O(1) 时间复杂度内删除链表节点
// 主要思路是根据给定的 node 节点拿到其 next 节点，将 next 节点的内容复制到给定节点内，
// 然后通过删除其 next 节点来达到删除该节点的目的，甚至不需要额外的空间来做这件事情
// 要考虑四个特殊情况：链表为空，链表只有一个元素，删除的元素是链表的末尾元素，删除的元素是链表的非末尾元素
func deleteNode(head *Node, toBeDeleted *Node) *Node {
	// 链表为空的情况
	if head == nil || toBeDeleted == nil {
		return head
	}
	// 链表只有一个元素的情况
	if head.next == nil && head == toBeDeleted {
		return nil
	}
	// 如果删除的是末尾元素，只能遍历整个链表来找到倒数第二个元素，改变他的 next 指向
	if toBeDeleted.next == nil {
		tmp := head
		for tmp.next != toBeDeleted {
			tmp = tmp.next
		}
		tmp.next = nil
		return head
	}
	toBeDeleted.value = toBeDeleted.next.value
	toBeDeleted.next = toBeDeleted.next.next
	return head
}
