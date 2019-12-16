package main

type Node struct {
	value int
	next  *Node
}

// 合并两个有序链表

// 主要思路是归并排序中合并两个数组的思路
func mergeSortedLinkedList(head1 *Node, head2 *Node) *Node {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	newHead := &Node{}
	tmp := newHead
	for head1 != nil && head2 != nil {
		if head1.value <= head2.value {
			tmp.next = head1
			head1 = head1.next
		} else {
			tmp.next = head2
			head2 = head2.next
		}
		tmp = tmp.next
	}

	if head1 != nil {
		tmp.next = head1
	}
	if head2 != nil {
		tmp.next = head2
	}
	return newHead.next
}
