package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

// 用栈和循环实现
func printLinkedListInReversedOrderIteratively(head *Node) {
	var stack []int
	pNode := head
	for pNode != nil {
		stack = append(stack, pNode.value)
		pNode = pNode.next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Println(stack[i])
	}
}

// 用递归实现
func printLinkedListInReversedOrderRecursively(head *Node) {
	if head == nil {
		return
	}
	printLinkedListInReversedOrderRecursively(head.next)
	fmt.Println(head.value)
}
