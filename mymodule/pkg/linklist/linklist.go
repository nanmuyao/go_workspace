package linklist

import "fmt"

// 标准库

type Node struct {
	value int32
	pre   *Node
	next  *Node
	cnt   uint32
}

func NewNode(v int32) *Node {
	n := new(Node)
	n.value = v
	n.pre = nil
	n.next = nil
	return n
}

var Head *Node = &Node{value: 0, pre: nil, next: nil}
var Tail *Node = &Node{value: 0, pre: nil, next: nil}

type LinkList struct {
	head *Node
	tail *Node
}

func NewLinkList() *LinkList {
	l := new(LinkList)
	l.head = Head
	l.tail = Tail
	l.head.next = Tail
	l.tail.pre = Head
	return l
}

func (l *LinkList) PushFround(n *Node) {
	next := l.head.next
	l.head.next = n
	n.pre = l.head
	n.next = next
	next.pre = n
}

func (l *LinkList) PushTail(n *Node) {
	pre := l.tail.pre
	pre.next = n
	l.tail.pre = n
	n.next = l.tail
}

func (l *LinkList) DeleteNode(n *Node) {
	pre := n.pre
	next := n.next
	pre.next = next
	next.pre = pre
	println("delete node")
}

func (l *LinkList) PrintLinkList() {
	n := l.head.next
	for ; n != nil; n = n.next {
		fmt.Printf("value = %d ", n.value)
	}
}
