package linkedlist

import "errors"

// Node linked list node
type Node struct {
	data interface{}
	next *Node
}

// LinkedList linked list
type LinkedList struct {
	head   *Node
	length int
}

// MakeEmptyList make an empty list
func MakeEmptyList() *LinkedList {
	return &LinkedList{
		head:   nil,
		length: 0,
	}
}

// Append insert an element at the end
func (l *LinkedList) Append(element interface{}) error {
	node := &Node{
		data: element,
		next: nil,
	}
	if l.head == nil {
		l.head = node
	} else {
		last := l.head
		for {
			if last.next == nil {
				last.next = node
				break
			} else {
				last = last.next
			}
		}
	}
	return nil
}

// Insert insert an element at specified position
func (l *LinkedList) Insert(element interface{}, k int) (bool, error) {
	if l.isIndexOver(k) {
		return false, errors.New("index out of range")
	}
	node := &Node{
		data: element,
	}
	i := 1
	last := l.head
	if k == 1 {
		l.head = &Node{
			data: element,
			next: last,
		}
	} else {
		for {
			if i == k-1 {
				last.next = &Node{
					data: element,
					next: last.next,
				}
				break
			}
			i++
			last = last.next
		}
	}
	l.length++
	return true, nil
}

// FindElement find the element according to the position
func (l *LinkedList) FindElement(k int) (interface{}, error) {

	return nil, nil
}

// FindPosition find the postion where the element first appearance
func (l *LinkedList) FindPosition(element interface{}) int {

}

// Delete delete an element
func (l *LinkedList) Delete(k int) (interface{}, error) {

	return nil, nil
}

// Length return list length
func (l *LinkedList) Length() int {
	return l.length
}

// Head return head node
func (l *LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) isIndexOver(k int) bool {
	if k <= 1 || k > l.length {
		return true
	}
	return false
}
